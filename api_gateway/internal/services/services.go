package services

import (
	grpc "api_gateway/internal/grpc/clients"
)

type Services struct {
	Auth AuthService
	Post PostService
}

func New(grpcClients *grpc.Clients) *Services {
	auth := NewAuthService(grpcClients.Auth)
	post := NewPostService(grpcClients.Post)

	return &Services{
		Auth: auth,
		Post: post,
	}
}
