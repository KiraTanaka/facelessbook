package services

import (
	grpc "api_gateway/internal/grpc/clients/user_service"
)

type Services struct {
	Auth AuthService
}

func Init(grpcClients *grpc.Clients) (*Services, error) {
	var err error
	services := &Services{}

	services.Auth, err = NewAuthService(grpcClients.Auth)
	if err != nil {
		return nil, err
	}

	return services, nil
}
