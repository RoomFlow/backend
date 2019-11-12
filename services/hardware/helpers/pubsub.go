package helpers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	config "github.com/RoomFlow/backend/config"

	"cloud.google.com/go/pubsub"
)

// Data struct which is expected from kafka
type hardwareData struct {
	Room    string                 `json:"room"`
	Content map[string]interface{} `json:"content"`
}

// NewPubSubReceive receives msgs
func NewPubSubReceive() error {
	// Initialize new PubSub client
	client, err := pubsub.NewClient(context.Background(), config.ProjectID)
	if err != nil {
		return fmt.Errorf("pubsub.NewClient(): %v", err)
	}

	log.Println("Pubsub client initialized")

	log.Printf("Receiving messages from sub '%s'\n", config.SubID)

	// Create PubSub subcription object
	sub := client.Subscription(config.SubID)

	// Start receiving messages from PubSub
	err = sub.Receive(context.Background(), func(ctx context.Context, msg *pubsub.Message) {
		log.Printf("Received message: %q\n", string(msg.Data))

		var data hardwareData

		// Convert bytes from kafka to map
		err := json.Unmarshal(msg.Data, &data)
		if err != nil {
			log.Printf("Unable to unmarshal kafka message: %v", err)
		} else {
			// Send data to firestore
			err = StoreSensorData(data.Room, data.Content)
			if err != nil {
				log.Printf("Error while storing data: %v", err)
			}
		}

		msg.Ack()
	})
	if err != nil {
		return fmt.Errorf("sub.Receive(): %v", err)
	}

	return nil
}
