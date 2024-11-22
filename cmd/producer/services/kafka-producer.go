package services

import (
	"github.com/gin-gonic/gin"
	"golearn/cmd/producer/config"
	"golearn/pkg/models"
	"net/http"
)

type KafkaProducer struct {
	kfConfig *config.KafkaConfig
}

func NewKafkaProducer(kfConfig *config.KafkaConfig) KafkaProducer {
	return KafkaProducer{kfConfig}
}

func (kf *KafkaProducer) SendMessageHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		message := models.GenerateRandomMessage()
		res, err := kf.kfConfig.SendKafkaMessage(message)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "System error!"})
			return
		}
		if res {
			ctx.JSON(http.StatusOK, gin.H{"message": "Success"})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Error no se emitió el evento"})
		}
	}
}
