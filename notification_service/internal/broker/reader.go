package broker

import (
	"notification_service/internal/config"
	"notification_service/internal/services"
)

type Readers struct {
	Post *PostReader
}

func NewConnect(config *config.KafkaConfig, services *services.Services) *Readers {
	postReader := NewPostReader(config, services.NotificationService)
	return &Readers{
		Post: postReader}
}

func (r *Readers) Run() {
	//go
	r.Post.Run()
}
