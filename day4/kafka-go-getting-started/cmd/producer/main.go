package main

import (
	"context"
	"crypto/tls"
	"kafka-go-getting-started/config"
	"kafka-go-getting-started/domain/message"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
)

func main() {
	cfg := config.LoadConfig(".")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mechanism := plain.Mechanism{
		Username: cfg.KafkaUser,
		Password: cfg.KafkaPassword,
	}

	dialer := &kafka.Dialer{
		Timeout:       10 * time.Second,
		DualStack:     true,
		SASLMechanism: mechanism,
		TLS:           &tls.Config{},
	}

	kafkaWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"pkc-ldvr1.asia-southeast1.gcp.confluent.cloud:9092"},
		Topic:   "notification",
		Dialer:  dialer,
	})

	defer kafkaWriter.Close()

	msgs := message.CreateMessage([]string{"This is 1st.", "This is 2nd"})

	err := kafkaWriter.WriteMessages(ctx, msgs...)
	if err != nil {
		log.Fatalln(err)
	}
}
