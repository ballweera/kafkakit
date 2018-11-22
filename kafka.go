package kafkakit

// Producer defines producer's behaviors
type Producer interface {
	Produce() (ok bool)
}
