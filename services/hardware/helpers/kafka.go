package helpers

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/Shopify/sarama"
)

// NewArduinoConsumer starts a new kafka consumer
func NewArduinoConsumer() error {
	// Kafka config
	config := sarama.NewConfig()
	config.ClientID = "hardware-consumer-client"
	config.Version = sarama.MaxVersion
	config.Consumer.Return.Errors = true
	// config.Net.SASL.Enable = true
	// config.Net.TLS.Enable = true
	// config.Net.SASL.User = "admin"
	// config.Net.SASL.Password = "password"

	// Hook up kafka logger to Stdout
	sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)

	// Kafka brokers to connect to
	brokers := []string{"localhost:9092"}

	// Create new consumer group
	group, err := sarama.NewConsumerGroup(brokers, "hardware-consumer-group", config)
	if err != nil {
		return err
	}

	// Topics to listen on
	topics := []string{"sensors"}

	// Create new consumer
	err = group.Consume(context.Background(), topics, consumerHandler{})
	if err != nil {
		group.Close()
		return err
	}

	// Go routine for logging errors
	go func() {
		for err := range group.Errors() {
			log.Printf("Consumer group error: %v", err)
		}
	}()

	return nil
}

// Data struct which is expected from kafka
type hardwareData struct {
	Room    string                 `json:"room"`
	Content map[string]interface{} `json:"content"`
}

type consumerHandler struct{}

func (consumerHandler) Setup(_ sarama.ConsumerGroupSession) error { return nil }

func (consumerHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }

func (h consumerHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		log.Printf("topic: %v, partition: %v, offset: %v", msg.Topic, msg.Partition, msg.Offset)

		log.Printf("msg: %v", string(msg.Value))

		var data hardwareData

		// Convert bytes from kafka to map
		err := json.Unmarshal(msg.Value, &data)
		if err != nil {
			log.Printf("Unable to unmarshal kafka message: %v", err)
		} else {
			// Send data to firestore
			err = StoreSensorData(data.Room, data.Content)
		}

		// Mark message as failed processing if error
		if err == nil {
			session.MarkMessage(msg, "Successfully stored sensor data")
		} else {
			session.MarkMessage(msg, "Error while processing sensor data")
		}
	}

	return nil
}
