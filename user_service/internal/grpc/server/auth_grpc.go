package grpc

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"user_service/internal/services"

	pb "github.com/KiraTanaka/facelessbook_protos/gen/auth"
)

type authServer struct {
	pb.UnimplementedAuthServer
	authService services.AuthService
}

func RegisterAuthServer(gRPCServer *grpc.Server, authService services.AuthService) {
	pb.RegisterAuthServer(gRPCServer, &authServer{authService: authService})
}

func (s *authServer) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	if err := validateRegistrationData(request); err != nil {
		return nil, err
	}

	userId, err := s.authService.Register(request.Phone, request.Password)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.RegisterResponse{UserId: userId}, nil
}

func (s *authServer) Login(ctx context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error) {
	if err := validateLoginData(request); err != nil {
		return nil, err
	}

	token, err := s.authService.Login(request.Phone, request.Password)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())

	}

	return &pb.LoginResponse{Token: token}, nil
}

func validateRegistrationData(request *pb.RegisterRequest) error {
	if request.Phone == "" {
		return status.Error(codes.InvalidArgument, "phone is required")
	}
	if request.Password == "" {
		return status.Error(codes.InvalidArgument, "password is required")
	}
	return nil

}

func validateLoginData(request *pb.LoginRequest) error {
	if request.Phone == "" {
		return status.Error(codes.InvalidArgument, "phone is required")
	}
	if request.Password == "" {
		return status.Error(codes.InvalidArgument, "password is required")
	}
	return nil

}
