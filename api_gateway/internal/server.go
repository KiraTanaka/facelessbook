package server

import (
	grpc "api_gateway/internal/clients/user_service"
	"api_gateway/internal/config"
	"api_gateway/internal/handlers"
	"api_gateway/internal/services"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	Config      *config.Config
	Routes      *gin.Engine
	AuthService services.AuthService
	grpcClients *grpc.Clients
}

func NewServer() (*Server, error) {
	server := &Server{}

	var err error
	server.Config, err = config.GetConfig()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	server.grpcClients, err = grpc.NewClient(server.Config)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	authService, err := services.NewAuthService(server.grpcClients.Auth)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	server.Routes = handlers.InitRoutes(authService)

	return server, nil
}

func (server *Server) Run() {
	server.Routes.Run(server.Config.ServerAddress)
}
