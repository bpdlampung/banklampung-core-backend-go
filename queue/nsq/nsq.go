package nsq

import (
	"banklampung-core/errors"
	"fmt"
	goNsq "github.com/nsqio/go-nsq"
	"time"
)

func (nsqConfig NsqConfig) Publish(topic string, body []byte) error {
	nsqConfig.logger.Info(fmt.Sprintf("Publish with topic: %s - %s", topic, string(body)))

	if producer == nil {
		getProducer, err := goNsq.NewProducer(nsqConfig.nsqAddress, goNsq.NewConfig())

		if err != nil {
			return err
		}

		producer = getProducer
	}

	return producer.Publish(topic, body)
}

func (nsqConfig NsqConfig) DeferredPublish(topic string, delay time.Duration, body []byte) error {
	nsqConfig.logger.Info(fmt.Sprintf("Deferred Publish with topic: %s, delay: %v - %s", topic, delay, string(body)))

	if producer == nil {
		getProducer, err := goNsq.NewProducer(nsqConfig.nsqAddress, goNsq.NewConfig())

		if err != nil {
			return err
		}

		producer = getProducer
	}

	return producer.DeferredPublish(topic, delay, body)
}

func (nsqConfig NsqConfig) RegisterHandlerConsumer(topic string, workerCount int, handler goNsq.HandlerFunc) error {
	nsqConfig.logger.Info(fmt.Sprintf("Registering NSQ consumer, topic: %s", topic))

	getConsumer, err := goNsq.NewConsumer(topic, nsqConfig.channel, goNsq.NewConfig())

	if err != nil {
		msg := fmt.Sprintf("Registering NSQ consumer, topic: %s - Failed, err: %v", topic, err.Error())

		nsqConfig.logger.Error(msg)
		return errors.InternalServerError(msg)
	}

	getConsumer.AddConcurrentHandlers(handler, workerCount)
	getConsumer.ChangeMaxInFlight(workerCount)

	err = getConsumer.ConnectToNSQLookupd(nsqConfig.nsqLookupdAddress)

	if err != nil {
		msg := fmt.Sprintf("Registering NSQ consumer, topic: %s - Failed, err: %v", topic, err.Error())

		nsqConfig.logger.Error(msg)
		return errors.InternalServerError(msg)
	}

	return nil
}
