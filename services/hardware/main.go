package main

import (
	"log"

	"github.com/RoomFlow/backend/services/hardware/helpers"
)

func main() {
	// Initialize new firestore client
	err := helpers.NewFirestoreClient()
	if err != nil {
		log.Fatal(err)
	}

	// Initialize new kafka consumer
	err = helpers.NewArduinoConsumer()
	if err != nil {
		log.Fatal(err)
	}
}
