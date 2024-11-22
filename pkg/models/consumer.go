package models

import (
	"encoding/json"
	"github.com/IBM/sarama"
	"log"
)

type Consumer struct {
	Store *NotificationStore
}

func (*Consumer) Setup(sarama.ConsumerGroupSession) error   { return nil }
func (*Consumer) Cleanup(sarama.ConsumerGroupSession) error { return nil }

func (consumer *Consumer) ConsumeClaim(
	sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		clave := string(msg.Key)
		var notification IndexRvModel
		err := json.Unmarshal(msg.Value, &notification)
		if err != nil {
			log.Printf("failed to unmarshal notification: %v", err)
			continue
		}
		consumer.Store.Add(clave, notification)
		sess.MarkMessage(msg, "")
	}
	return nil
}
