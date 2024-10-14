//go:build consumer
// +build consumer

package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	// Create Kafka consumer
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "iot_group",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		log.Fatalf("Failed to create consumer: %s", err)
	}
	defer consumer.Close() // Ensure the consumer is closed when main exits

	// Subscribe to the topic
	err = consumer.Subscribe("iot_topic", nil)
	if err != nil {
		log.Fatalf("Failed to subscribe to topic: %s", err)
	}

	// channel to listen for interrupt signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

    // SIGINT: Interrupt signal, SIGTERM: Terminate signal

	go func() {
		<-sigChan
		log.Println("Received shutdown signal, closing consumer...")
		consumer.Close() // graceful exit
		os.Exit(0)     
	}()

	for {
		// new messages
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			log.Printf("Received message: %s", msg.Value)
		} else {
			log.Printf("Consumer error: %v (%v)", err, msg)
		}
	}
}