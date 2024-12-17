package services

import (
	grpc "api_gateway/internal/clients/user_service"

	log "github.com/sirupsen/logrus"
)

type authService struct {
	grpcClient *grpc.AuthClient
}

type AuthService interface {
	Register(phone string, password string) (user_id string, err error)
	Login(phone string, password string) (toker string, err error)
}

func NewAuthService(grpcClient *grpc.AuthClient) (AuthService, error) {
	return &authService{
		grpcClient: grpcClient,
	}, nil
}

func (s *authService) Register(phone string, password string) (string, error) {

	userId, err := s.grpcClient.Register(phone, password)
	if err != nil {
		log.Error(err)
		return "", err
	}

	return userId, err
}

func (s *authService) Login(phone string, password string) (string, error) {
	token, err := s.grpcClient.Login(phone, password)
	if err != nil {
		log.Error(err)
		return "", err
	}

	return token, err
}
