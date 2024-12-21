package grpc

import (
	"fmt"
	"log/slog"
	"net"
	"post_service/internal/config"
	"post_service/internal/services"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type Server struct {
	grpcServer *grpc.Server
	port       int
}

func NewServer(config *config.GrpcConfig, services *services.Services) *Server {
	grpcServer := grpc.NewServer()
	RegisterPostServer(grpcServer, services.Post)

	return &Server{
		grpcServer: grpcServer,
		port:       config.Port,
	}
}

func (s *Server) Run() error {
	const op = "grpcserver.Run"

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("grpc server started", slog.String("addr", listener.Addr().String()))

	if err := s.grpcServer.Serve(listener); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
