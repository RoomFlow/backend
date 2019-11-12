# backend

## Prerequisites

### Installing Go

Follow instructions outlined here: <https://golang.org/dl/>

## Running a Go service

`go run services/<service_name>/main.go`

## Notes about the `hardware` service

### Google Cloud Platform authentication

Ask Karlo for the authentication file which was not pushed to the repo and place it in the `secrets` directory. 

Next, you will need to run `export GOOGLE_APPLICATION_CREDENTIALS=<PATH_TO_BACKEND_REPO>/backend/secrets/roomflow-service-account.json`.

The authentication file will be automatically detected and used.

### Google Cloud PubSub

PubSub is Kafka as a Service as implemented by Google Cloud Platform (GCP).

Here is a sample message that can be execute in a GCP console `gcloud pubsub topics publish sensors --message '{"room":"bsb-102","content":{"sound":"123","motion":"on"}}'` (make sure you have switched to use the RoomFlow project on GCP).
