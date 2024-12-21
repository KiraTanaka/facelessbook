package broker

import (
	"context"
	"fmt"
	"notification_service/internal/config"
	"notification_service/internal/services"

	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
)

type PostReader struct {
	kafka               *kafka.Reader
	notificationService services.NotificationService
}

type Reader interface {
	Run()
}

func NewPostReader(config *config.KafkaConfig, notificationService services.NotificationService) *PostReader {
	return &PostReader{
		kafka: kafka.NewReader(kafka.ReaderConfig{
			Brokers:   []string{fmt.Sprintf("%s:%d", config.Host, config.Port)},
			Topic:     "posts",
			Partition: 0,
			MaxBytes:  10e6, // 10MB
		}),
		notificationService: notificationService}
}

func (r *PostReader) Run() {
	for {
		kafkaMessage, err := r.kafka.ReadMessage(context.Background())
		if err != nil {
			log.Error(err)
			return
		}
		message, err := ToNewPostMessage(kafkaMessage.Value)
		if err != nil {
			log.Error(err)
			continue
		}
		r.notificationService.ProcessNewPost(message)
	}
}
