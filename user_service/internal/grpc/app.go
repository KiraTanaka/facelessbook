package server_grpc

import (
	"fmt"
	"log/slog"
	"net"
	"user_service/internal/config"
	"user_service/internal/services"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	gRPCServer *grpc.Server
	port       int
}

func New(config *config.Config, authService services.AuthService) *GRPCServer {
	gRPCServer := grpc.NewServer()
	RegisterAuthServer(gRPCServer, authService)

	return &GRPCServer{
		gRPCServer: gRPCServer,
		port:       config.GrpcPort,
	}
}

func (s *GRPCServer) Run() error {
	const op = "grpcserver.Run"

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("grpc server started", slog.String("addr", listener.Addr().String()))

	// Запускаем обработчик gRPC-сообщений
	if err := s.gRPCServer.Serve(listener); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
