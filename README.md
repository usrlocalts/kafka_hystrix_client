# Kafka Hystrix Client

This Go Client library publishes a message to kafka given a topic, and pushes to a redis fallback queue in case of a failure.
The library uses a circuit breaker hystrix logic for fallbacks

## Getting Started

`glide get github.com/usr-local-ts/kafka_hystrix_client`

You can include this library in your glide.yaml/project by using:

`import github.com/usr-local-ts/kafka_hystrix_client`

If you want to avoid entering username and password, in your glide.yaml, specifiy as follows:

```
- package: github.com/usr-local-ts/kafka_hystrix_client
  repo: git@github.com:usr-local-ts/kafka_hystrix_client.git
  vcs: git

```

## Contributing to the library

- `mkdir -p github.com/usr-local-ts`
- Clone the repository `git clone git@github.com:usr-local-ts/kafka_hystrix_client.git`

- Run `make setup`

This does the following things:

- glide install

- Run `make test` to run tests


## Setting up go

This service runs on go.

- Install go
    - On OSX run `brew install go`.
    - Follow instructions on https://golang.org/doc/install for other OSes.
- Setup go
      - Make sure that the executable `go` is in your shell's path.
      - Add the following in your .zshrc or .bashrc: (where `<workspace_dir>` is the directory in
        which you'll checkout your code)
- Run Test
    make test

```
GOPATH=<workspace_dir>
export GOPATH
PATH="${PATH}:${GOPATH}/bin"
export PATH
```

## Usage

After you include the library in your application,

##### Create a client by doing the following:

```
kafkaHystrixNotificationClient, hystrixClientErr := client.NewKafkaHystrixClient(*KafkaHystrixClientConfig())

func KafkaHystrixClientConfig() *kafkaHystrixConfig.Config {
	kafkaHystrixClientConfig := kafkaHystrixConfig.NewConfig()
	kafkaHystrixClientConfig.Topic = NotificationRequestTopic()
	kafkaHystrixClientConfig.LogLevel = LogLevel()
	kafkaHystrixClientConfig.KafkaBootstrapServers = KafkaBootstrapServers()
	kafkaHystrixClientConfig.KafkaFallbackKey = KafkaFallbackKey()
	kafkaHystrixClientConfig.KafkaFallbackExpiry = KafkaFallbackExpiry()
	kafkaHystrixClientConfig.KafkaFallbackRedisURI = KafkaFallbackRedisURI()
	kafkaHystrixClientConfig.KafkaFallbackRedisPassword = KafkaFallbackRedisPassword()
	kafkaHystrixClientConfig.KafkaFallbackRedisIdleConn = KafkaFallbackRedisIdleConn()
	kafkaHystrixClientConfig.KafkaFallbackRedisMaxConn = KafkaFallbackRedisMaxConn()
	return kafkaHystrixClientConfig
}

Assuming all these confgurations are in your config file

```

`Please refer to application.yml.sample` for setting up config values. Use these values in your source repository.

##### The default configurations when you do a `kafkaHystrixConfig.NewConfig()` are:

```
LogLevel:                   "debug",
KafkaBootstrapServers:      []string{"localhost:9092"},
KafkaFallbackKey:           "default_topic",
KafkaFallbackExpiry:        180,
KafkaFallbackRedisURI:      "localhost:6379",
KafkaFallbackRedisIdleConn: 5,
KafkaFallbackRedisMaxConn:  5,

```

##### The default configurations for sarama producer are:

```
config.Producer.Partitioner = sarama.NewRandomPartitioner
config.Producer.RequiredAcks = sarama.WaitForAll
config.Producer.Return.Errors = true
config.Producer.Return.Successes = true
```

##### You can edit these config by editing the sarama configs:
 
```
kafkaHystrixClientConfig.SaramaConfig.Producer.Partitioner = sarama.NewRandomPartitioner
kafkaHystrixClientConfig.SaramaConfig.Producer.RequiredAcks = sarama.WaitForLocal
kafkaHystrixClientConfig.SaramaConfig.Producer.Return.Errors = true
kafkaHystrixClientConfig.SaramaConfig.Producer.Return.Successes = true

```

##### You can call the `ProduceWithFallback method` as follows

```
KafkaHystrixClient.ProduceWithFallback(message, config.NotificationRequestTopic())
```

`The message is expected to be a byte array`

# Tasks Glossary
`make setup` - Sets up the project
`make test` - Runs Tests
`make proto-gen` - Generate proto for update proto contracts

Please refer the Makefile for the entire tasks list
