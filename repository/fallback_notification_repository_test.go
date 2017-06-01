package repository

import (
	"testing"
	"github.com/usr-local-ts/kafka_hystrix_client/util"
	"github.com/stretchr/testify/assert"
	"github.com/garyburd/redigo/redis"
	"github.com/usr-local-ts/kafka_hystrix_client/config"
	"github.com/usr-local-ts/kafka_hystrix_client/logger"
	"github.com/usr-local-ts/kafka_hystrix_client/appcontext"
	"encoding/base64"
)

func TestPublishToQueue(t *testing.T) {
	config.LoadTest()
	appcontext.Initiate()
	logger.SetupLogger()
	key := config.KafkaFallbackKey() + ":random_test_topic"
	message := []byte("hello")
	fallbackNotifRepo := FallbackNotificationRepository{}
	util.WithClearFallback(key, func(conn redis.Conn) {
		err := fallbackNotifRepo.PublishToQueue(key, message)
		assert.Equal(t, err.Code, "")
	})
}

func TestPublishToQueueAndRetrieve(t *testing.T) {
	config.LoadTest()
	appcontext.Initiate()
	logger.SetupLogger()
	key := config.KafkaFallbackKey() + ":" + config.NotificationRequestTopic()
	message := []byte("hello")
	fallbackNotifRepo := FallbackNotificationRepository{}
	util.WithClearFallback(key, func(conn redis.Conn) {

		base64Message := make([]byte, base64.StdEncoding.EncodedLen(len(message)))
		base64.StdEncoding.Encode(base64Message, []byte(message))

		err := fallbackNotifRepo.PublishToQueue(key, message)
		assert.Equal(t, err.Code, "")
		val, _ := getFallbackValue(key)
		assert.Equal(t, base64Message, []byte(val))
	})
}

func getFallbackValue(key string) (string, error) {
	kafkaFallbackRedisClient := appcontext.KafkaFallbackRedisClient()
	conn := kafkaFallbackRedisClient.Get()
	return redis.String(conn.Do("GET", key))
}
