package reader

import (
	"DataReader/adapter"
	"DataReader/data_supplier_reader/reader/dtos"
	"DataReader/environment"
	"DataReader/util"
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type KafkaReader struct{}

func (kafkaReader KafkaReader) Read(consumer util.Consumer[dtos.ReadCommand], wg *sync.WaitGroup) {
	// Get environment variables for Kafka configuration
	kafkaBroker := os.Getenv(string(environment.KafkaBroker))
	kafkaTopic := os.Getenv(string(environment.DataReaderWriteTopic))
	kafkaGroupID := os.Getenv(string(environment.DataReaderGroupID))

	// Setup Kafka reader configuration
	readerConfig := kafka.ReaderConfig{
		Brokers: []string{kafkaBroker},
		Topic:   kafkaTopic,
		GroupID: kafkaGroupID,
	}

	// Create a new Kafka reader
	reader := kafka.NewReader(readerConfig)
	defer reader.Close()

	// Create a signal channel to gracefully handle termination (Ctrl+C)
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("Starting Kafka consumer...")

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Error reading message: %v", err)
			continue
		}

		// Process the message (here we simply print it)
		fmt.Printf("Received message: %s\n", string(msg.Value))
		command := adapter.ConvertToCommand(msg.Value)
		consumer.Consume(command)

		// Gracefully shut down on receiving SIGINT or SIGTERM
		select {
		case <-sigchan:
			fmt.Println("\nReceived termination signal, shutting down...")
			return
		default:
		}
	}
}
