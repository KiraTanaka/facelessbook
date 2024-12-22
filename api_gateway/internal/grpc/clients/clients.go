package grpc

import (
	"api_gateway/internal/config"
	post "api_gateway/internal/grpc/clients/post_service"
	user "api_gateway/internal/grpc/clients/user_service"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Clients struct {
	Auth       *user.AuthClient
	User       *user.UserClient
	Subscriber *user.SubscriberClient
	Post       *post.PostClient
}

func NewClients(config *config.GrpcConfig) (*Clients, error) {
	clients := &Clients{}
	err := clients.NewUserServiceClients(config)
	if err != nil {
		return nil, err
	}
	err = clients.NewPostServiceClients(config)
	if err != nil {
		return nil, err
	}

	return clients, nil
}

func (c *Clients) NewUserServiceClients(config *config.GrpcConfig) error {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", config.UserHost, config.UserPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("new user grpc clients: %v", err)
	}

	c.Auth = user.NewAuthClient(conn, config.Timeout)
	c.User = user.NewUserClient(conn, config.Timeout)
	c.Subscriber = user.NewSubscriberClient(conn, config.Timeout)

	return nil
}

func (c *Clients) NewPostServiceClients(config *config.GrpcConfig) error {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", config.PostHost, config.PostPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("new post grpc clients: %v", err)
	}

	c.Post = post.NewPostClient(conn, config.Timeout)

	return nil
}
