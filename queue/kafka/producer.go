package kafka

import (
	"encoding/json"
	goKafka "github.com/Shopify/sarama"
)

func (p KafkaConfig) Publish(topic string, partition int32, payload interface{}) (*int32, error) {
	err := p.createTopic(topic, partition)

	if err != nil {
		return nil, err
	}

	byteData, err := json.Marshal(payload)

	if err != nil {
		return nil, err
	}

	partititon, _, err := p.syncProducer.SendMessage(&goKafka.ProducerMessage{
		Topic: topic,
		Value: goKafka.ByteEncoder(byteData),
	})

	if err != nil {
		return nil, err
	}

	return &partititon, err
}

func (k KafkaConfig) createTopic(topic string, partition int32) error {
	exist, err := k.isAvailableTopic(topic)

	if err != nil {
		return err
	}

	if exist {
		return nil
	}

	err = k.clusterAdmin.CreateTopic(topic, &goKafka.TopicDetail{
		NumPartitions:     partition,
		ReplicationFactor: 1,
	}, false)

	if err != nil {
		return err
	}

	return nil
}
