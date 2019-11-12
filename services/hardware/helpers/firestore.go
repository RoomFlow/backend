package helpers

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
)

var client *firestore.Client

// NewFirestoreClient returns a connection to the firestore
func NewFirestoreClient() error {
	// New context
	ctx := context.Background()

	// Initialize firebase app
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		return fmt.Errorf("firebase.NewApp(): %v", err)
	}

	log.Println("Firebase app initialized")

	// Create new firestore client
	client, err = app.Firestore(ctx)
	if err != nil {
		return fmt.Errorf("app.Firestore(): %v", err)
	}

	log.Println("Firestore client initialized")

	return nil
}

// StoreSensorData stores data into firestore
func StoreSensorData(collection string, data map[string]interface{}) error {
	// Store data into firestore
	documentRef, writeResult, err := client.Collection(collection).Add(context.Background(), data)
	if err != nil {
		return fmt.Errorf("client.Collection.Add(): %v", err)
	}

	log.Printf("Data stored successfully into document %v at %v\n", documentRef.ID, writeResult.UpdateTime)

	return nil
}
