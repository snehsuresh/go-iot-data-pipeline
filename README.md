# Real-time Data Processing Pipeline

## Project Overview
This project implements a real-time data processing pipeline using Go (Golang) for simulating IoT devices, Kafka for message brokering, and Prometheus for monitoring metrics 
The producer generates random temperature data for multiple simulated devices and sends it to a Kafka topic 
The consumer listens to that topic and processes the incoming messages, logging the received data 
Additionally, Prometheus collects and exposes the temperature metrics via an HTTP endpoint for monitoring

## Technologies Used
- **Programming Language**: Go (Golang)
- **Message Broker**: Apache Kafka
- **Monitoring**: Prometheus
- **Metrics Library**: Prometheus Go Client
- **Concurrency**: Goroutines

## Project Structure
```plaintext
iot-data-pipeline/
├── producer/
│   └── main.go
└── consumer/
    └── main.go
```

## Getting Started

### Prerequisites
- Go (1.16 or later)
- Apache Kafka (with a running broker)
- Prometheus (for monitoring)

### Installation
1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd iot-data-pipeline
   ```

2. Install the required Go packages:
   ```bash
   go get github.com/confluentinc/confluent-kafka-go/kafka
   go get github.com/prometheus/client_golang/prometheus
   go get github.com/prometheus/client_golang/prometheus/promhttp
   ```

3. Start your Kafka broker (if not already running).

### Running the Producer
.1. Navigate to the producer directory:
   ```bash
   cd producer
   ```

2. Run the producer with the build tag:
   ```bash
   go run -tags producer main.go
   ```

### Running the Consumer
1. Open a new terminal window and navigate to the consumer directory:
   ```bash
   cd consumer
   ```

2. Run the consumer with the build tag:
   ```bash
   go run -tags consumer main.go
   ```

## Monitoring with Prometheus
1. Start a Prometheus server with the following configuration:
   ```yaml
   global:
     scrape_interval: 15s
   scrape_configs:
     - job_name: 'go_app'
       static_configs:
         - targets: ['localhost:8080']
   ```

2. Access Prometheus UI at `http://localhost:9090` and configure it to scrape from the producer's metrics endpoint at `http://localhost:8080/metrics`.

## How It Works
1. The producer simulates multiple IoT devices that generate random temperature readings every 2 seconds.
2. Each temperature reading is sent as a JSON message to the `iot_topic` in Kafka.
3. The consumer listens for messages on the `iot_topic`, logging each received temperature data.
4. Prometheus collects metrics from the producer's HTTP endpoint and makes them available for monitoring.

## Future Improvements
- Implement error handling for Kafka message production and consumption.
- Add more metrics to monitor the health and performance of the pipeline.
- Store temperature readings in a database for historical analysis.
- Create a web interface to visualize the temperature data over time.

