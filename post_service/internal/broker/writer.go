package broker

import (
	"context"
	"encoding/json"
	"fmt"
	"post_service/internal/config"
	"post_service/internal/models"

	"github.com/segmentio/kafka-go"
)

type PostWriter struct {
	broker *kafka.Writer
}

type Writer interface {
	SendMessage(mess *models.NewPostMessage) error
}

func NewConnect(config *config.KafkaConfig) Writer {
	return &PostWriter{&kafka.Writer{
		Addr:     kafka.TCP(fmt.Sprintf("%s:%d", config.Host, config.Port)),
		Topic:    "posts",
		Balancer: &kafka.LeastBytes{},
	}}
}

func (w *PostWriter) SendMessage(mess *models.NewPostMessage) error {
	value, err := NewPostMessageValue(mess)
	if err != nil {
		return fmt.Errorf("new post message value: %w", err)
	}
	err = w.broker.WriteMessages(context.Background(),
		kafka.Message{
			Value: value,
		},
	)
	if err != nil {
		return fmt.Errorf("broker write message about new post: %w", err)
	}
	return nil
}

func NewPostMessageValue(m *models.NewPostMessage) ([]byte, error) {
	value, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return value, nil
}
