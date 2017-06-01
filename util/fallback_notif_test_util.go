package util

import (
	"github.com/garyburd/redigo/redis"
	"github.com/usr-local-ts/kafka_hystrix_client/appcontext"
)

func WithClearFallback(key string, block func(redis.Conn)) {
	kafkaFallbackRedisClient := appcontext.KafkaFallbackRedisClient()
	conn := kafkaFallbackRedisClient.Get()
	clearFallbackKey(key, conn)
	block(conn)
	clearFallbackKey(key, conn)
	conn.Close()
}

func clearFallbackKey(key string, conn redis.Conn) {
	conn.Send("DEL", key)
	conn.Flush()
}
