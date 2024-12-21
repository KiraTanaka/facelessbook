package grpc

import (
	"api_gateway/internal/config"
	post "api_gateway/internal/grpc/clients/post_service"
	user "api_gateway/internal/grpc/clients/user_service"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	log "github.com/sirupsen/logrus"
)

type Clients struct {
	Auth *user.AuthClient
	User *user.UserClient
	Post *post.PostClient
}

func NewClients(config *config.GrpcConfig) (*Clients, error) {
	clients := &Clients{}
	err := clients.NewUserServiceClients(config)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	err = clients.NewPostServiceClients(config)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return clients, nil
}

func (c *Clients) NewUserServiceClients(config *config.GrpcConfig) error {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", config.UserHost, config.UserPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect to grpc server: %v", err)
	}

	c.Auth, err = user.NewAuthClient(conn)
	c.User, err = user.NewUserClient(conn)

	return nil
}

func (c *Clients) NewPostServiceClients(config *config.GrpcConfig) error {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", config.PostHost, config.PostPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect to grpc server: %v", err)
	}

	c.Post, err = post.NewPostClient(conn)

	return nil
}
