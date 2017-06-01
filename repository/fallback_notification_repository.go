//Fallback Notification Client that will publish to hystrix redis
package repository

import (
	"github.com/usr-local-ts/kafka_hystrix_client/config"
	"github.com/usr-local-ts/kafka_hystrix_client/logger"
	"github.com/usr-local-ts/kafka_hystrix_client/proto"
	"github.com/usr-local-ts/kafka_hystrix_client/appcontext"
	e "github.com/usr-local-ts/kafka_hystrix_client/errors"
	"encoding/base64"

)

type FallbackNotificationInterface interface {
	PublishToQueue(key string, message []byte) proto.Errors
}

type FallbackNotificationRepository struct {
}

func (self *FallbackNotificationRepository) PublishToQueue(key string, message []byte) proto.Errors {
	kafkaFallbackRedis := appcontext.KafkaFallbackRedisClient()
	conn := kafkaFallbackRedis.Get()
	defer conn.Close()

	base64Message := make([]byte, base64.StdEncoding.EncodedLen(len(message)))
	base64.StdEncoding.Encode(base64Message, []byte(message))

	logger.Log.Info("Saving to fallback Queue")
	_, err := conn.Do("SET", key, base64Message)
	if err == nil {
		conn.Do("EXPIRE", key, config.KafkaFallbackExpiry())

	} else {
		logger.Log.Error("Failed to save to fallback queue", err)
		return proto.Errors{e.GenericServiceError, "could not publish to kafka"}
	}
	return proto.Errors{}
}
