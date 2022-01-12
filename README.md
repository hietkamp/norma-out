# NORMA Outbound

Norma is de Engelse benaming van het sterrenbeeld winkelhaak. Norma is geimplementeerd als prototype om specificaties te verifiëren op volledigheid en correctheid. Het betreffen de specificaties voor gevalideerde vragen van het programma KIK-V. 

Norma outbound is geimplementeerd voor het datastation van een bronhouder van de data. Het datastation is een concept waarin leverancierneutraal data beschikbaar wordt gesteld door een zorginstelling, voor haarzelf en haar omgeving. Met de Norma outbound kan een bronhouder de antwoorden op gevalideerde vragen terugsturen naar de afnemer. De outboud ontvangt de antwoorden via Kafka Streams.

De worker maakt gebruik van Kafka Streams. Zie hiervoor morma-gw om kafka te kunnen starten.

**NOTE:**
:warning: De software is "as-is".
Er wordt geen ondersteuning op verleend en is niet voor bedoeld voor productie!
---

## Installeren

```bash
git clone git@github.com:hietkamp/norma-out.git
```

## Opstarten

```bash
./manage go
```
