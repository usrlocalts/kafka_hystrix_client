package appcontext

import (
	"fmt"
	"time"
	"math/rand"

	raven "github.com/getsentry/raven-go"
	"github.com/garyburd/redigo/redis"
	_ "github.com/lib/pq"

	"github.com/usr-local-ts/kafka_hystrix_client/config"
	"github.com/usr-local-ts/kafka_hystrix_client/logger"
)

type appContext struct {
	kafkaFallbackRedisClient *redis.Pool
}

var context *appContext

type appContextError struct {
	Error error
}

func panicIfError(err error, werr error) {
	if err != nil {
		panic(appContextError{werr})
	}
}

func Initiate() {
	kafkaFallbackedisClient := initRedis(config.KafkaFallbackRedisURI(), config.KafkaFallbackRedisMaxConn(), config.KafkaFallbackRedisIdleConn())

	context = &appContext{
		kafkaFallbackRedisClient: kafkaFallbackedisClient,
	}
}

func initRedis(redisURI string, redisMaxConn, redisIdleConn int) *redis.Pool {
	redis := &redis.Pool{
		MaxIdle:   redisIdleConn,
		MaxActive: redisMaxConn,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redisURI)
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			randSource := rand.NewSource(time.Now().UnixNano())
			randGenerator := rand.New(randSource)
			if randGenerator.Intn(100) < 10 {
				_, err := c.Do("PING")
				return err
			}
			return nil

		},
		IdleTimeout: 5 * time.Second,
		Wait:        true,
	}
	conn := redis.Get()
	defer conn.Close()

	err := conn.Send("PING")

	if err != nil {
		raven.CaptureError(err, map[string]string{"Unable to connect to redis server": err.Error()})
		logger.Log.Errorf("Unable to connect to redis server: %s %s", redisURI, err)
		panic(fmt.Errorf("Unable to connect to redis server %s: %s", redisURI, err.Error()))
	}
	return redis
}

func KafkaFallbackRedisClient() *redis.Pool {
	return context.kafkaFallbackRedisClient
}
