package main

import (
	"log"

	"github.com/Shopify/sarama"
)

func main() {
	consumer, err := sarama.NewConsumer([]string{"localhost:19092"}, nil)
	if err != nil {
		log.Fatalln("Failed to start Sarama consumer:", err)
	}
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition("example-topic", 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalln("Failed to start Sarama partition consumer:", err)
	}
	defer partitionConsumer.Close()

	for msg := range partitionConsumer.Messages() {
		log.Printf("Received message: %s\n", string(msg.Value))
	}
}
