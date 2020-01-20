package main

import (
	"context"
	"log"

	pkg "github.com/RoomFlow/backend/pkg/helpers"
	"github.com/RoomFlow/backend/services/schedule/helpers"
)

func main() {
	// Fetch room data from McMaster room directory
	rooms := helpers.FetchRoomData()

	// Initialize new firestore client
	client := pkg.NewFirestoreClient()
	defer client.Close()

	log.Println("Creating new batch write")

	// Create new batch write so we can avoid creating one write per room
	batch := client.Batch()

	// Collection which will hold the room data
	collRef := client.Collection("rooms")

	// Iterate through room data and add new document to batch
	for _, data := range rooms {
		docRef := collRef.Doc(data.ID)
		batch.Set(docRef, data)
	}

	log.Println("Committing batch write")

	// Commit the batch.
	_, err := batch.Commit(context.TODO())
	if err != nil {
		log.Fatalf("An error has occurred: %s", err)
	}

	log.Println("Done updating room data")
}
