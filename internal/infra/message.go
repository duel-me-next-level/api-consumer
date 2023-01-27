package message

import (
	"context"

	"github.com/segmentio/kafka-go"
	"google.golang.org/grpc"
)

// SendGRPCMessage sends a message over GRPC
func SendGRPCMessage(conn *grpc.ClientConn, topic string, message *kafka.Message) error {
	// Create a new client for the GRPC service
	client := NewGRPCClient(conn)

	// Send the message to the specified topic
	_, err := client.SendMessage(context.Background(), &GRPCMessage{
		Topic:   topic,
		Message: message,
	})
	if err != nil {
		return err
	}

	return nil
}

// SendKafkaMessage sends a message over Kafka
func SendKafkaMessage(writer *kafka.Writer, topic string, message []byte) error {
	// Write the message to the specified topic
	err := writer.WriteMessages(context.Background(),
		kafka.Message{
			Value: message,
			Topic: topic,
		},
	)
	if err != nil {
		return err
	}
	return nil
}
