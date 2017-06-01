package client

import (
	"github.com/usr-local-ts/kafka_hystrix_client/proto"
	"github.com/usr-local-ts/kafka_hystrix_client/logger"
	"github.com/usr-local-ts/kafka_hystrix_client/config"
	e "github.com/usr-local-ts/kafka_hystrix_client/errors"
	"github.com/usr-local-ts/kafka_hystrix_client/factories"
	"github.com/Shopify/sarama"
	"github.com/afex/hystrix-go/hystrix"
	"errors"
	"github.com/usr-local-ts/kafka_hystrix_client/repository"
	"github.com/usr-local-ts/kafka_hystrix_client/appcontext"
)

type KafkaHystrixClient struct {
	KafkaClientConfig config.Config
	SyncProducer      sarama.SyncProducer
}

type KafkaHystrixClientInterface interface {
	ProduceWithFallback(message []byte, topic string)
}

func NewKafkaHystrixClient(requestConfig config.Config) (*KafkaHystrixClient, proto.Errors) {
	config.Load(requestConfig)
	appcontext.Initiate()
	logger.SetupLogger()
	syncProducer, err := factories.CreateSaramaSyncProducer()
	if err != nil {
		logger.Log.Error(err)
		return &KafkaHystrixClient{}, proto.Errors{"900", "could not create client"}
	}
	kafkaHystrixClient := KafkaHystrixClient{requestConfig, syncProducer }
	return &kafkaHystrixClient, proto.Errors{}

}

func (self *KafkaHystrixClient) ProduceWithFallback(message []byte, topic string) proto.Errors {

	messageToSend := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(message),
	}

	hystrixErr := hystrix.Do(topic, func() (errHystrix error) {
		_, _, err := self.SyncProducer.SendMessage(messageToSend)
		if err != nil {
			return err
		}
		return nil
	}, func(err error) error {
		logger.Log.Warn("Failed to publish to kafka", err)
		fallbackRep := repository.FallbackNotificationRepository{}
		fallbackTopic := config.KafkaFallbackKey() + ":" + topic
		queueErr := fallbackRep.PublishToQueue(fallbackTopic, message)
		if queueErr.Code != "" {
			logger.Log.Error(queueErr)
			return errors.New(queueErr.Code)
		}
		return nil
	})

	if hystrixErr != nil {
		logger.Log.Error(hystrixErr)
		return proto.Errors{e.GenericServiceError, "could not save to fallback queue"}
	}

	return proto.Errors{}
}
