package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/Shopify/sarama/mocks"
	"github.com/Shopify/sarama"
	"github.com/usr-local-ts/kafka_hystrix_client/config"
	"github.com/usr-local-ts/kafka_hystrix_client/logger"
)

func TestProduceMessage(t *testing.T) {
	config.LoadTest()
	logger.SetupLogger()
	producer := mocks.NewSyncProducer(t, sarama.NewConfig())
	kafkaHystrixClient := KafkaHystrixClient{*config.AppConfig(), producer}
	notificationReq := []byte("Test Message to Kafka")
	producer.ExpectSendMessageAndSucceed()
	err := kafkaHystrixClient.ProduceWithFallback(notificationReq, "random_topic")
	assert.Equal(t, err.Code, "")
}
