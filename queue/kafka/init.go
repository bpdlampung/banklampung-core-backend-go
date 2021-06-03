package kafka

import (
	"errors"
	"fmt"
	goKafka "github.com/Shopify/sarama"
	"github.com/bpdlampung/banklampung-core-backend-go/logs"
	"time"
)

type KafkaConfig struct {
	username     string
	password     string
	address      string
	logger       logs.Collections
	syncProducer goKafka.SyncProducer
	clusterAdmin goKafka.ClusterAdmin
}

func (kafkaConfig KafkaConfig) generateSyncProducer() KafkaConfig {
	kafkaConfig.logger.Info("Trying Generate Sync Producer -Kafka")
	syncProducer, err := goKafka.NewSyncProducer([]string{kafkaConfig.address}, kafkaConfig.getKafkaConfig())

	if err != nil {
		if err == goKafka.ErrOutOfBrokers {
			kafkaConfig.logger.Error(fmt.Sprintf("Cannot Connect Kafka -%s", kafkaConfig.address))
			time.Sleep(10 * time.Second)

			return kafkaConfig.generateSyncProducer()
		}
		kafkaConfig.logger.Error(err.Error())
		panic(err)
	}

	kafkaConfig.syncProducer = syncProducer
	kafkaConfig.logger.Info("Success Generate Sync Producer -Kafka")
	return kafkaConfig
}

func (kafkaConfig KafkaConfig) generateClusterAdmin() KafkaConfig {
	kafkaConfig.logger.Info("Trying Generate Cluster Admin -Kafka")
	clusterAdmin, err := goKafka.NewClusterAdmin([]string{kafkaConfig.address}, kafkaConfig.getKafkaConfig())

	if err != nil {
		if err == goKafka.ErrOutOfBrokers {
			kafkaConfig.logger.Error(fmt.Sprintf("Cannot Connect Kafka -%s", kafkaConfig.address))
			time.Sleep(10 * time.Second)

			return kafkaConfig.generateSyncProducer()
		}

		kafkaConfig.logger.Error(err.Error())
		panic(err)
	}

	kafkaConfig.clusterAdmin = clusterAdmin
	kafkaConfig.logger.Info("Success Generate Cluster Admin -Kafka")
	return kafkaConfig
}

func InitConfig(kafkaUrl, username, password string, logger logs.Collections) KafkaConfig {

	kafkaCfg := KafkaConfig{
		address:  kafkaUrl,
		username: username,
		password: password,
		logger:   logger,
	}

	kafkaCfg = kafkaCfg.generateClusterAdmin()
	kafkaCfg = kafkaCfg.generateSyncProducer()

	logger.Info("Kafka Generate Config --Success")

	return kafkaCfg
}

func (k KafkaConfig) getKafkaConfig() *goKafka.Config {
	kafkaConfig := goKafka.NewConfig()

	kafkaConfig.Producer.Return.Successes = true
	kafkaConfig.Net.WriteTimeout = 1 * time.Second

	//if k.username != "" {
	//	kafkaConfig.Net.SASL.Enable = true
	//	kafkaConfig.Net.SASL.User = k.username
	//	kafkaConfig.Net.SASL.Password = k.password
	//}

	version, err := goKafka.ParseKafkaVersion("2.1.1")
	if err != nil {
		panic(fmt.Sprintf("Error parsing Kafka version: %v", err))
	}

	kafkaConfig.Version = version
	kafkaConfig.Consumer.Offsets.Initial = goKafka.OffsetOldest
	kafkaConfig.Consumer.Group.Rebalance.Strategy = goKafka.BalanceStrategyRange
	kafkaConfig.Consumer.Fetch.Min = 1
	kafkaConfig.Consumer.Fetch.Max = 1000

	return kafkaConfig
}

func (k KafkaConfig) isAvailableTopic(topic string) (bool, error) {
	if k.clusterAdmin == nil {
		return false, errors.New("EOF")
	}

	topicList, err := k.clusterAdmin.ListTopics()

	if err != nil {
		return false, err
	}

	_, ok := topicList[topic]

	if ok {
		return true, nil
	}

	return false, nil
}
