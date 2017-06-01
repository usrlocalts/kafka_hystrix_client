package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	configVars := map[string]string{
		"LOG_LEVEL":                      "debug",
		"TOPIC":                          "3000",
		"kafka_fallback_redis_uri":       "localhost:6379",
		"kafka_fallback_redis_idle_conn": "5",
		"kafka_fallback_redis_max_conn":  "5",
		"notification_request_topic":     "notification_requests",
		"kafka_bootstrap_servers":        "localhost:9092",
		"kafka_fallback_key":             "random_test_topic",
	}

	for k, v := range configVars {
		os.Setenv(k, v)
		defer os.Unsetenv(k)
	}

	LoadTest()
	assert.Equal(t, configVars["notification_request_topic"], NotificationRequestTopic())
	assert.Equal(t, configVars["notification_request_topic"], NotificationRequestTopic())
	assert.Equal(t, configVars["LOG_LEVEL"], LogLevel())
	assert.Equal(t, configVars["kafka_fallback_redis_uri"], KafkaFallbackRedisURI())
	assert.Equal(t, 5, KafkaFallbackRedisIdleConn())
	assert.Equal(t, 5, KafkaFallbackRedisMaxConn())
	assert.True(t, SaramaConfig().Producer.Return.Successes)
	assert.True(t, SaramaConfig().Producer.Return.Errors)
}
