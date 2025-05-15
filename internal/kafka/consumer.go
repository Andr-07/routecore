package kafka

import (
	"context"
	"log"
	"routecore/configs"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

type KafkaConsumer struct {
	Reader        *kafka.Reader
}

func NewKafkaConsumer(conf *configs.KafkaConfig) *KafkaConsumer {
	return &KafkaConsumer{
		Reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers: []string{conf.Broker},
			Topic:   conf.Topic,
			GroupID: uuid.NewString(),
		}),
	}
}

func (k *KafkaConsumer) ReadAll() error {
	for {
		log.Println("Waiting for message from Kafka...")

		msg, err := k.Reader.ReadMessage(context.Background())
		if err != nil {
			log.Println("❌ Error reading message:", err)
			continue
		}

		log.Printf("✅ Received message from Kafka: topic=%s partition=%d offset=%d value=%s",
			msg.Topic, msg.Partition, msg.Offset, string(msg.Value),
		)
	}
}
