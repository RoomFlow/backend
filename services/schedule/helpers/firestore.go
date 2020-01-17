package helpers

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
)

var client *firestore.Client

// NewFirestoreClient returns a connection to the firestore
func NewFirestoreClient() *firestore.Client {
	// New context
	ctx := context.Background()

	// Initialize firebase app
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		log.Fatalf("firebase.NewApp(): %v", err)
	}

	log.Println("Firebase app initialized")

	// Create new firestore client
	client, err = app.Firestore(ctx)
	if err != nil {
		log.Fatalf("app.Firestore(): %v", err)
	}

	log.Println("Firestore client initialized")

	return client
}
