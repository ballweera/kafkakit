package kafkakit

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSaramaProducer(t *testing.T) {
	t.Run("produces message to 'example' topic successfully", func(t *testing.T) {
		producer := NewSaramaProducer()
		msg := "Hello Kafka!!"

		got := producer.Produce(msg)
		want := true

		assert.Equal(t, want, got, fmt.Sprintf("want %v, but got %v", want, got))
	})
}
