package services

import (
	grpc "api_gateway/internal/grpc/clients"
)

type Services struct {
	Auth       AuthService
	Post       PostService
	Subscriber SubscriberService
}

func New(grpcClients *grpc.Clients) *Services {
	auth := NewAuthService(grpcClients.Auth)
	post := NewPostService(grpcClients.Post)
	subscriber := NewSubscriberService(grpcClients.Subscriber)

	return &Services{
		Auth:       auth,
		Post:       post,
		Subscriber: subscriber,
	}
}
