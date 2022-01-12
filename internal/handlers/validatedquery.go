package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/hietkamp/norma-out/internal/eventstream"
	"github.com/rs/zerolog/log"
	uuid "github.com/satori/go.uuid"
)

func PostValidatedQueryAnswer(v ValidatedQueriesProcessed) error {
	httpClient := &http.Client{}

	reqResponse := RequestResponse{
		Id:        v.Header.MessageId,
		Resultset: v.Payload,
	}
	bodyBytes, _ := json.Marshal(reqResponse)
	r, err := http.NewRequest("POST", v.Header.ReplyTo, bytes.NewReader(bodyBytes))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	u1 := uuid.NewV4()
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Accept", "application/json")
	r.Header.Add("Accept-Charset", "UTF-8")
	r.Header.Add("Message-ID", u1.String())
	r.Header.Add("Subject", "validated-query-response")
	// Switch the from and to because its a response message of the request
	r.Header.Add("From", v.Header.To)
	r.Header.Add("To", v.Header.From)
	// Keep the conversation in the reference
	r.Header.Add("References", v.Header.References)

	resp, err := httpClient.Do(r)
	if err != nil {
		return fmt.Errorf("failed to post answer: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New("failed to post answer: service not available")
	}
	return nil
}

func ProduceMessage(v ValidatedQueriesAnswered) error {
	messageBytes, _ := json.Marshal(v)
	brokers := strings.Split(os.Getenv("kafka_url"), ",")
	es := eventstream.New("tcp", brokers[1])
	err := es.Produce(os.Getenv("topic_producer"), messageBytes)
	if err != nil {
		return fmt.Errorf("failed to produce message: %w", err)
	}
	return nil
}

func HandleValidatedQueryProcessed(message []byte) error {
	var vqProcessed ValidatedQueriesProcessed
	json.Unmarshal(message, &vqProcessed)
	log.Debug().Msgf("Outgoing answer reveived: %s", vqProcessed.Header.MessageId)
	err := PostValidatedQueryAnswer(vqProcessed)
	if err != nil {
		return err
	}
	ts := time.Now()
	log.Debug().Msg("Produce event stream")
	vqAnswered := ValidatedQueriesAnswered{
		Timestamp: ts.Format(time.RFC3339),
		Header:    vqProcessed.Header,
		Payload:   vqProcessed.Payload,
	}
	err = ProduceMessage(vqAnswered)
	if err != nil {
		return err
	}
	return nil
}
