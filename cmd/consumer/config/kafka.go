package config

import (
	"fmt"
	"github.com/IBM/sarama"
	"os"
)

func InitializeConsumerGroup() (sarama.ConsumerGroup, error) {
	config := sarama.NewConfig()
	kafkaServerAddress := os.Getenv("KAFKA_SERVER")
	consumerGroup := os.Getenv("KAFKA_CONSUMER_GROUP")
	group, err := sarama.NewConsumerGroup(
		[]string{kafkaServerAddress}, consumerGroup, config)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize consumer group: %w", err)
	}
	return group, nil
}
