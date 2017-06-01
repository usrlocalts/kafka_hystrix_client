package config

import "github.com/Shopify/sarama"

type Config struct {
	LogLevel                   string
	Topic                      string
	KafkaBootstrapServers      []string
	KafkaFallbackKey           string
	KafkaFallbackExpiry        int
	KafkaFallbackRedisURI      string
	KafkaFallbackRedisPassword string
	KafkaFallbackRedisIdleConn int
	KafkaFallbackRedisMaxConn  int
	SaramaConfig               *sarama.Config
}

var appConfig *Config

func Load(config Config) {

	appConfig = &Config{
		LogLevel:                   config.LogLevel,
		Topic:                      config.Topic,
		KafkaBootstrapServers:      config.KafkaBootstrapServers,
		KafkaFallbackKey:           config.KafkaFallbackKey,
		KafkaFallbackExpiry:        config.KafkaFallbackExpiry,
		KafkaFallbackRedisURI:      config.KafkaFallbackRedisURI,
		KafkaFallbackRedisPassword: config.KafkaFallbackRedisPassword,
		KafkaFallbackRedisIdleConn: config.KafkaFallbackRedisIdleConn,
		KafkaFallbackRedisMaxConn:  config.KafkaFallbackRedisMaxConn,
		SaramaConfig:               config.SaramaConfig,
	}
}

func LoadTest() {

	appConfig = &Config{
		LogLevel:                   "debug",
		Topic:                      "notification_requests",
		KafkaBootstrapServers:      []string{"localhost:9092"},
		KafkaFallbackKey:           "random_test_topic",
		KafkaFallbackExpiry:        180,
		KafkaFallbackRedisURI:      "localhost:6379",
		KafkaFallbackRedisPassword: "",
		KafkaFallbackRedisIdleConn: 5,
		KafkaFallbackRedisMaxConn:  5,
		SaramaConfig:               SaramaDefaultConfig(),
	}
}

func NewConfig() *Config {
	newConfig := &Config{
		LogLevel:                   "debug",
		KafkaBootstrapServers:      []string{"localhost:9092"},
		KafkaFallbackKey:           "default_topic",
		KafkaFallbackExpiry:        180,
		KafkaFallbackRedisURI:      "localhost:6379",
		KafkaFallbackRedisIdleConn: 5,
		KafkaFallbackRedisMaxConn:  5,
		SaramaConfig:               SaramaDefaultConfig(),
	}
	return newConfig
}

func LogLevel() string {
	return appConfig.LogLevel
}

func AppConfig() *Config {
	return appConfig
}

func NotificationRequestTopic() string {
	return appConfig.Topic
}

func KafkaBootstrapServers() []string {
	return appConfig.KafkaBootstrapServers
}

func KafkaFallbackKey() string {
	return appConfig.KafkaFallbackKey
}

func KafkaFallbackExpiry() int {
	return appConfig.KafkaFallbackExpiry
}

func KafkaFallbackRedisURI() string {
	return appConfig.KafkaFallbackRedisURI
}

func KafkaFallbackRedisPassword() string {
	return appConfig.KafkaFallbackRedisPassword
}

func KafkaFallbackRedisIdleConn() int {
	return appConfig.KafkaFallbackRedisIdleConn
}

func KafkaFallbackRedisMaxConn() int {
	return appConfig.KafkaFallbackRedisMaxConn
}

func SaramaConfig() *sarama.Config {
	return appConfig.SaramaConfig
}

func SaramaDefaultConfig() *sarama.Config {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Errors = true
	config.Producer.Return.Successes = true
	return config
}
