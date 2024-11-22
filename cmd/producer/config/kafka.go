package config

import (
	"encoding/json"
	"errors"
	"github.com/IBM/sarama"
	"golearn/pkg/models"
	"log"
	"os"
)

type KafkaConfig struct {
	Producer sarama.SyncProducer
}

func NewCreateKafkaConfig() KafkaConfig {
	KafkaServerAddress := os.Getenv("KAFKA_SERVER")
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer([]string{KafkaServerAddress}, config)
	if err != nil {
		log.Fatalln("Failed to start Sarama producer:", err)
	}
	return KafkaConfig{producer}
}
func (ks *KafkaConfig) SendKafkaMessage(msg models.IndexRvModel) (bool, error) {
	kafkaTopic := os.Getenv("KAFKA_TOPIC")

	jsonMessage, err := json.Marshal(msg)
	if err != nil {
		return false, errors.New("error marshalling message")
	}
	res := &sarama.ProducerMessage{
		Topic: kafkaTopic,
		Key:   sarama.StringEncoder("INDEX"),
		Value: sarama.StringEncoder(jsonMessage),
	}
	_, _, err = ks.Producer.SendMessage(res)
	if err != nil {
		return false, err
	}
	return true, nil
}
