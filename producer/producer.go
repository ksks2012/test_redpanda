package main

import (
	"log"

	"github.com/Shopify/sarama"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer([]string{"localhost:19092"}, config)
	if err != nil {
		log.Fatalln("Failed to start Sarama producer:", err)
	}
	defer producer.Close()

	// Send a message to the "example-topic"
	msg := &sarama.ProducerMessage{
		Topic: "example-topic",
		Value: sarama.StringEncoder("Hello, Redpanda!"),
	}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}
	log.Printf("Message sent to partition %d at offset %d\n", partition, offset)
}
