# backend

## Prerequisites

### Installing Go

Follow instructions outlined here: <https://golang.org/dl/>

## Running a Go service

`go run services/<service_name>/main.go`

## Notes about the `hardware` service

### Kafka

Kafka needs to be running locally on `localhost:9092` in order for the service to work

Follow steps 1-4 in order to create a Zookeeper server, Kafka server, topic, and to send messages via producer: <https://kafka.apache.org/quickstart>

Be sure to create the topic `sensors` and produce to it instead of the one in the quickstart link.

### Firebase

Ask Karlo for the Firebase authentication file which was not pushed to the repo and place it in the `secrets` directory.

Here is a sample message that can be sent to the producer `{"room":"bsb-101","content":{"sound":"123","motion":"on"}}`.

This creates a document in the `bsb-101` collection with the fields `sound:"123"` and `motion:"on"`.
