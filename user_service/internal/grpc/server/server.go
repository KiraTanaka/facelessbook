package grpc

import (
	"fmt"
	"net"
	"user_service/internal/config"
	"user_service/internal/services"

	"google.golang.org/grpc"
)

type Server struct {
	grpcServer *grpc.Server
	port       int
}

func NewServer(config *config.GrpcConfig, services *services.Services) *Server {
	grpcServer := grpc.NewServer()
	RegisterAuthServer(grpcServer, services.Auth)
	RegisterUserServer(grpcServer, services.User)
	RegisterSubscriberServer(grpcServer, services.Subscriber)

	return &Server{
		grpcServer: grpcServer,
		port:       config.GrpcPort,
	}
}

func (s *Server) Run() error {
	const op = "grpcserver.Run"

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if err := s.grpcServer.Serve(listener); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
