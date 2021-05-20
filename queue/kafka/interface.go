package kafka

// Collections is kafka's collection of function
type Collections interface {
	RegisterHandlerConsumerGroup(groupId, topic string, handler HandlerFunc) error
	Publish(topic string, partition int32, payload interface{}) (*int32, error)
}
