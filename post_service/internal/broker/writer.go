package broker

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type PostWriter struct {
	kafka *kafka.Writer
}

type Writer interface {
	SendMessage(mess Message) error
}

func NewConnect() Writer {
	return &PostWriter{&kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    "posts",
		Balancer: &kafka.LeastBytes{},
	}}
}

func (w *PostWriter) SendMessage(mess Message) error {
	value, err := mess.ToValue()
	if err != nil {
		return err
	}
	err = w.kafka.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("newPost"),
			Value: value,
		},
	)
	return err
}
