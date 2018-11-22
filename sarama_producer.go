package kafkakit

// SaramaProducer implements kafkakit/Producer with Shopify/sarama
type SaramaProducer struct {
}

// Produce sends message to Kafka topic
func (smp *SaramaProducer) Produce(msg string) (ok bool) {
	return false
}

// NewSaramaProducer creates SaramaProducer pointer
func NewSaramaProducer() *SaramaProducer {
	return &SaramaProducer{}
}
