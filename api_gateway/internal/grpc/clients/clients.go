package grpc

import (
	"api_gateway/internal/config"

	log "github.com/sirupsen/logrus"
)

type Clients struct {
	Auth *AuthClient
	Post *PostClient
}

func NewClients(config *config.GrpcConfig) (*Clients, error) {
	authClient, err := NewAuthClient(config.GrpcHost, config.GrpcAuthPort)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	postClient, err := NewPostClient(config.GrpcHost, config.GrpcPostPort)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &Clients{
		Auth: authClient,
		Post: postClient,
	}, nil
}
