package handlers

type HeaderEnvelop struct {
	MessageId  string `header:"Message-ID"`
	Subject    string `header:"Subject"`
	From       string `header:"From"`
	To         string `header:"To"`
	References string `header:"References"`
	ReplyTo    string `header:"Reply-To"`
}

type QueryResult struct {
	Result interface{} `json:"result,omitempty"`
}

type ValidatedQueriesProcessed struct {
	Timestamp string        `json:"ts,omitempty"`
	Header    HeaderEnvelop `json:"header,omitempty"`
	Payload   []QueryResult `json:"payload,omitempty"`
}

type ValidatedQueriesAnswered struct {
	Timestamp string        `json:"ts,omitempty"`
	Header    HeaderEnvelop `json:"header,omitempty"`
	Payload   []QueryResult `json:"payload,omitempty"`
}

type RequestResponse struct {
	Id        string        `json:"id,omitempty"`
	Resultset []QueryResult `json:"resultset,omitempty"`
}
