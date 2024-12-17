package supplier

import (
	"DataReader/environment"
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"os"
)

var writer *kafka.Writer

func GetWriter() *kafka.Writer {
	if writer == nil {
		writer = getInitializedWriter()
	}

	return writer
}

func getInitializedWriter() *kafka.Writer {
	kafkaBroker := os.Getenv(string(environment.KafkaBroker))
	kafkaTopic := os.Getenv(string(environment.DataReaderReadTopic))

	return &kafka.Writer{
		Addr:  kafka.TCP(kafkaBroker),
		Topic: kafkaTopic,
	}
}

func Send(objectToSend interface{}, key string) error {
	messageToSend, jsonMarshallingError := json.Marshal(objectToSend)
	if jsonMarshallingError != nil {
		log.Printf("Error serializing user: %v", jsonMarshallingError)
		return jsonMarshallingError
	}

	msg := kafka.Message{
		Key:   []byte(key),
		Value: messageToSend,
	}

	writeMessageError := GetWriter().WriteMessages(context.Background(), msg)
	if writeMessageError != nil {
		log.Printf("Failed to write message: %v", jsonMarshallingError)
	} else {
		fmt.Printf("Produced message: %s\n", messageToSend)
	}

	//defer writer.Close()
	return nil
}
