package helpers

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var client *firestore.Client

// NewFirestoreClient returns a connection to the firestore
func NewFirestoreClient() error {
	// New context
	ctx := context.Background()

	// Specify credentials file
	opt := option.WithCredentialsFile("secrets/sensorDataStoreKey.json")

	// Firebase config
	config := &firebase.Config{ProjectID: "roomflow-e2004"}

	// Initialize firebase app
	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		return err
	}

	// Create new firestore client
	client, err = app.Firestore(ctx)
	if err != nil {
		return err
	}

	return nil
}

// StoreSensorData stores data into firestore
func StoreSensorData(collection string, data map[string]interface{}) error {
	_, _, err := client.Collection(collection).Add(context.Background(), data)
	if err != nil {
		log.Printf("Failed storing sensor data: %v", err)
		return err
	}

	return nil
}
