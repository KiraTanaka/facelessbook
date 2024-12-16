package server_grpc

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"user_service/internal/services"

	auth "github.com/KiraTanaka/facelessbook_protos/gen/auth"
)

type serverAuth struct {
	auth.UnimplementedAuthServer
	authService services.AuthService
}

func RegisterAuthServer(gRPCServer *grpc.Server, authService services.AuthService) {
	auth.RegisterAuthServer(gRPCServer, &serverAuth{authService: authService})
}

func (s *serverAuth) Register(ctx context.Context, request *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	if err := validateRegistrationData(request); err != nil {
		return nil, err
	}

	userId, err := s.authService.Register(request.GetPhone(), request.GetPassword())
	if err != nil {
		return nil, status.Error(codes.Internal, "failed register")
	}

	return &auth.RegisterResponse{UserId: userId}, nil
}

func (s *serverAuth) Login(ctx context.Context, request *auth.LoginRequest) (*auth.LoginResponse, error) {
	if err := validateLoginData(request); err != nil {
		return nil, err
	}

	token, err := s.authService.Login(request.GetPhone(), request.GetPassword())
	if err != nil {
		return nil, status.Error(codes.Internal, "failed login")

	}

	return &auth.LoginResponse{Token: token}, nil
}

func validateRegistrationData(request *auth.RegisterRequest) error {
	if request.GetPhone() == "" {
		return status.Error(codes.InvalidArgument, "phone is required")
	}
	if request.GetPassword() == "" {
		return status.Error(codes.InvalidArgument, "password is required")
	}
	return nil

}

func validateLoginData(request *auth.LoginRequest) error {
	if request.GetPhone() == "" {
		return status.Error(codes.InvalidArgument, "phone is required")
	}
	if request.GetPassword() == "" {
		return status.Error(codes.InvalidArgument, "password is required")
	}
	return nil

}
