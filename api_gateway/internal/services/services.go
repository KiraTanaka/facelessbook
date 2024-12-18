package services

import (
	grpc "api_gateway/internal/grpc/clients"
)

type Services struct {
	Auth AuthService
	Post PostService
}

func Init(grpcClients *grpc.Clients) (*Services, error) {
	auth, err := NewAuthService(grpcClients.Auth)
	if err != nil {
		return nil, err
	}
	post, err := NewPostService(grpcClients.Post)
	if err != nil {
		return nil, err
	}

	return &Services{
		Auth: auth,
		Post: post,
	}, nil
}
