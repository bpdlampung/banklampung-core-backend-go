package kafka

import (
	"context"
	"fmt"
	goKafka "github.com/Shopify/sarama"
	"time"
)

func (c KafkaConfig) RegisterHandlerConsumerGroup(groupId, topic string, handler HandlerFunc) error {
	for {
		isAvailableTopic, err := c.isAvailableTopic(topic)

		if err != nil {
			if err.Error() == "EOF" {
				c.logger.Error(fmt.Sprintf("Cannot Connect Kafka at Consumer -EOF -%s", c.address))
				time.Sleep(10 * time.Second)

				c = c.generateSyncProducer()
				c = c.generateClusterAdmin()
				return c.RegisterHandlerConsumerGroup(groupId, topic, handler)
			}

			fmt.Println(err.Error(), "<<<< Lagi??")
			return err
		}

		if !isAvailableTopic {
			c.logger.Info(fmt.Sprintf("Topic not found -%s", topic))
			time.Sleep(15 * time.Second)
			continue
		}

		break
	}

	c.addConcurrentHandlers(handler, groupId, topic)
	return nil
}

func (c KafkaConfig) addConcurrentHandlers(handler HandlerFunc, groupId, topic string) {
	consumerGroupConn, err := goKafka.NewConsumerGroup([]string{c.address}, groupId, c.getKafkaConfig())

	if err != nil {
		panic(err)
	}

	defer consumerGroupConn.Close()

	/**
	* Setup a new Sarama consumer group
	 */
	consumer := ConsumerGroup{
		handler: handler,
	}

	fmt.Println(fmt.Sprintf("Worker %v, Topic: '%s' - GroupId: '%s' is Running", 1, topic, groupId))

	for {
		err := consumerGroupConn.Consume(context.Background(), []string{topic}, &consumer)

		if err != nil {
			fmt.Println("kafkaConsumer", err.Error(), topic, "")
		}
	}
}

// -- config consumer group

type Message struct {
	Key, Value []byte
	Topic      string
	Partition  int32
	Offset     int64
	Context    context.Context
}

type HandlerConsumer interface {
	HandleMessage(message *Message) error
}

/**
 * HandlerFunc is a convenience type to avoid having to declare a structs
 * to implement the Handler interface, it can be used like this:
 *
 * handlers := qKafka.HandlerFunc(func(m *qKafka.Message) error {
 * 		   var age int
 *  	   if err := json.Unmarshal(m.Value, &age); err != nil {
 * 				  return err
 *		   }
 *		   return nil
 * })
 **/
type HandlerFunc func(message *Message) error

/**
 * HandleMessage implements the Handler interface
 **/
func (h HandlerFunc) HandleMessage(m *Message) error {
	return h(m)
}

// Consumer represents a Sarama consumer group consumer
type ConsumerGroup struct {
	handler HandlerFunc
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (consumer *ConsumerGroup) Setup(goKafka.ConsumerGroupSession) error {
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (consumer *ConsumerGroup) Cleanup(goKafka.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (consumer *ConsumerGroup) ConsumeClaim(session goKafka.ConsumerGroupSession, claim goKafka.ConsumerGroupClaim) error {

	// NOTE:
	// Do not move the code below to a goroutine.
	// The `ConsumeClaim` itself is called within a goroutine
	//defer func() {
	//	if v := recover(); v != nil {
	//		e := apmPkg.GetTracer().Recovered(v)
	//		e.Send()
	//	}
	//}()

	for message := range claim.Messages() {
		fmt.Println("consumer", string(message.Value), message.Topic, fmt.Sprintf("Partition: %v - Offset: %v", message.Partition, message.Offset))

		err := consumer.handler.HandleMessage(&Message{
			Value:     message.Value,
			Topic:     message.Topic,
			Key:       message.Key,
			Offset:    message.Offset,
			Partition: message.Partition,
			//Context:   ctx,
		})

		if err != nil {
			fmt.Println("kafkaConsumer", err.Error(), message.Topic, "")
			return err
		}

		session.MarkMessage(message, "")
	}

	return nil
}
