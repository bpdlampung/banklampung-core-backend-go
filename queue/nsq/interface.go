package nsq

import (
	goNsq "github.com/nsqio/go-nsq"
	"time"
)

// Collections is nsq's collection of function
type Collection interface {
	Publish(topic string, body []byte) error
	DeferredPublish(topic string, delay time.Duration, body []byte) error
	RegisterHandlerConsumer(topic string, workerCount int, handler goNsq.HandlerFunc) error
}
