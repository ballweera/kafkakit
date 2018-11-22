package kafkakit

import (
	"log"

	"github.com/Shopify/sarama"
)

// SaramaProducer implements kafkakit/Producer with Shopify/sarama
type SaramaProducer struct {
}

// Produce sends message to Kafka topic
func (smp *SaramaProducer) Produce(msg string) (ok bool) {
	producer, err := sarama.NewAsyncProducer([]string{"localhost:9092"}, nil)

	if err != nil {
		log.Fatal(err)
		return false
	}

	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	producer.Input() <- &sarama.ProducerMessage{
		Topic: "example",
		Key:   nil,
		Value: sarama.StringEncoder(msg),
	}

	return true
}

// NewSaramaProducer creates SaramaProducer pointer
func NewSaramaProducer() *SaramaProducer {
	return &SaramaProducer{}
}
