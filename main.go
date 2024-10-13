package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
    temperatureGauge = prometheus.NewGaugeVec(
        prometheus.GaugeOpts{
            Name: "temperature_celsius",
            Help: "Temperature in Celsius",
        },
        []string{"device"},
    )
)

func init() {
    // Register the temperature gauge
    prometheus.MustRegister(temperatureGauge)
}

func generateTemperature(deviceID string, producer *kafka.Producer) {
    for {
        // Generate random temperature
        temperature := rand.Float64()*100
        log.Printf("Device %s: Temperature: %.2f°C\n", deviceID, temperature)

        // Produce to Kafka
        value := fmt.Sprintf(`{"device": "%s", "temperature": %.2f}`, deviceID, temperature)
        producer.Produce(&kafka.Message{
            TopicPartition: kafka.TopicPartition{Topic: &"iot_topic", Partition: kafka.PartitionAny},
            Value:          []byte(value),
        }, nil)

        // Update Prometheus gauge
        temperatureGauge.WithLabelValues(deviceID).Set(temperature)

        time.Sleep(2 * time.Second) // Simulate data generation every 2 seconds
    }
}

func main() {
    // Create Kafka producer
    producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
    if err != nil {
        log.Fatalf("Failed to create producer: %s", err)
    }
    defer producer.Close()

    // Start HTTP server for Prometheus metrics
    http.Handle("/metrics", promhttp.Handler())
    go http.ListenAndServe(":8080", nil)

    // Simulate multiple IoT devices
    for i := 1; i <= 5; i++ {
        deviceID := fmt.Sprintf("device_%d", i)
        go generateTemperature(deviceID, producer)
    }

    // Block forever
    select {}
}
