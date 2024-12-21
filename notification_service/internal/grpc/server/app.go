package server_grpc

import (
	"fmt"
	"log/slog"
	"net"
	"notification_service/internal/config"
	"notification_service/internal/services"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type Server struct {
	grpcServer *grpc.Server
	port       int
}

func New(config *config.Config, notificationService services.NotificationService) *GRPCServer {
	gRPCServer := grpc.NewServer()
	Register(gRPCServer, notificationService)

	return &GRPCServer{
		gRPCServer: gRPCServer,
		port:       config.GrpcPort,
	}
}

func NewServer(config *config.GrpcConfig, services *services.Services) *Server {
	grpcServer := grpc.NewServer()
	RegisterPostServer(grpcServer, services.Noti)

	return &Server{
		grpcServer: grpcServer,
		port:       config.GrpcPort,
	}
}

func (s *GRPCServer) MustRun() {
	if err := s.Run(); err != nil {
		panic(err)
	}
}

func (s *GRPCServer) Run() error {
	const op = "grpcapp.Run"

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

type Server struct {
	grpcServer *grpc.Server
	port       int
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
