package factories

import (
	"github.com/Shopify/sarama"
	"github.com/usr-local-ts/kafka_hystrix_client/config"
	"github.com/usr-local-ts/kafka_hystrix_client/logger"
)

type SaramaSyncProducerFactory interface {
	CreateSaramaSyncProducer() (sarama.SyncProducer, error)
}

func CreateSaramaSyncProducer() (sarama.SyncProducer, error) {
	logger.Log.Info("Trying with Kafka Bootstrap servers")
	logger.Log.Info(config.KafkaBootstrapServers())
	producer, err := sarama.NewSyncProducer(config.KafkaBootstrapServers(), config.SaramaConfig())
	if err != nil {
		logger.Log.Warn(err)
		return nil, err
	}
	return producer, nil
}
