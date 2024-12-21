package grpc

import (
	"fmt"
	"notification_service/internal/config"
	user "notification_service/internal/grpc/clients/user_service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	log "github.com/sirupsen/logrus"
)

type Clients struct {
	User       *user.UserClient
	Subscriber *user.SubscriberClient
}

func NewClients(config *config.GrpcConfig) (*Clients, error) {
	clients := &Clients{}
	err := clients.NewUserServiceClients(config)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return clients, nil
}

func (c *Clients) NewUserServiceClients(config *config.GrpcConfig) error {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", config.GrpcUserHost, config.GrpcUserPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect to grpc server: %v", err)
	}

	c.User = user.NewUserClient(conn, config.Timeout)
	c.Subscriber = user.NewSubscriberClient(conn, config.Timeout)

	return nil
}
