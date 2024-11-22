package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golearn/cmd/consumer/config"
	"golearn/pkg/models"
	"log"
	"os"
)

func setupConsumerGroup(ctx context.Context, store *models.NotificationStore) {
	kafkaTopic := os.Getenv("KAFKA_TOPIC")
	consumerGroup, err := config.InitializeConsumerGroup()
	if err != nil {
		log.Printf("initialization error: %v", err)
	}
	defer consumerGroup.Close()
	consumer := &models.Consumer{Store: store}
	for {
		err = consumerGroup.Consume(ctx, []string{kafkaTopic}, consumer)
		if err != nil {
			log.Printf("error from consumer: %v", err)
		}
		if ctx.Err() != nil {
			return
		}
	}
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	consumerPort := os.Getenv("CONSUMER_PORT")
	store := &models.NotificationStore{
		Data: make(models.IndexRvNotifications),
	}
	ctx, cancel := context.WithCancel(context.Background())
	go setupConsumerGroup(ctx, store)
	defer cancel()
	log.Printf("Kafka CONSUMER (Group: %s) 👥📥 "+
		"started at http://localhost%s\n", "indexrv", consumerPort)
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	if err := router.Run(consumerPort); err != nil {
		log.Printf("failed to run the server: %v", err)
	}
}
