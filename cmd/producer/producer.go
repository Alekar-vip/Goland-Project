package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golearn/cmd/producer/config"
	"golearn/cmd/producer/services"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	ProducerPort := os.Getenv("PRODUCER_PORT")
	kfConfig := config.NewCreateKafkaConfig()

	defer kfConfig.Producer.Close()

	service := services.NewKafkaProducer(&kfConfig)
	router := gin.Default()
	router.GET("/produce", service.SendMessageHandler())

	server := &http.Server{
		Addr:    ProducerPort,
		Handler: router,
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		log.Println("receive interrupt signal")
		if err := server.Close(); err != nil {
			log.Fatal("Server Close:", err)
		}
	}()

	if err := server.ListenAndServe(); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			log.Println("Server closed under request")
		} else {

			log.Fatal("Server closed unexpect " + err.Error())
		}
	}

	log.Println("Server exiting")
}
