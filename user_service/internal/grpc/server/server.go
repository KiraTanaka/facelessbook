package grpc

import (
	"fmt"
	"net"
	"user_service/internal/config"
	"user_service/internal/services"

	log "github.com/sirupsen/logrus"

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
		port:       config.ServerPort,
	}
}

func (s *Server) Run() error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return fmt.Errorf("listen announces on the local network address: %w", err)
	}

	log.Info("Grpc server started, ", fmt.Sprintf("addr = %s", listener.Addr().String()))

	if err = s.grpcServer.Serve(listener); err != nil {
		return fmt.Errorf("receiving incoming connections on the listener: %w", err)
	}

	return nil
}
