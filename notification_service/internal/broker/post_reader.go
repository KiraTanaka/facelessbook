package broker

import (
	"context"
	"encoding/json"
	"fmt"
	"notification_service/internal/config"
	"notification_service/internal/models"
	"notification_service/internal/services"

	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
)

type PostReader struct {
	broker              *kafka.Reader
	notificationService services.NotificationService
}

type Reader interface {
	Run()
}

func NewPostReader(config *config.KafkaConfig, notificationService services.NotificationService) *PostReader {
	return &PostReader{
		broker: kafka.NewReader(kafka.ReaderConfig{
			Brokers:   []string{fmt.Sprintf("%s:%d", config.Host, config.Port)},
			Topic:     "posts",
			Partition: 0,
			MaxBytes:  10e6, // 10MB
		}),
		notificationService: notificationService}
}

func (r *PostReader) Run() {
	log.Info("Runs post reader.")
	for {
		brokerMessage, err := r.broker.ReadMessage(context.Background())
		if err != nil {
			log.Error(fmt.Errorf("broker read message about new post: %w", err))
			return
		}
		message, err := NewPostMessage(brokerMessage.Value)
		if err != nil {
			log.Error(fmt.Errorf("new post message: %w", err))
			continue
		}
		if err = r.notificationService.ProcessNewPost(message); err != nil {
			log.Error(fmt.Errorf("process new post: %w", err))
			continue
		}
	}
}

func NewPostMessage(value []byte) (*models.NewPostMessage, error) {
	message := &models.NewPostMessage{}
	err := json.Unmarshal(value, message)
	if err != nil {
		return nil, err
	}
	return message, nil
}
