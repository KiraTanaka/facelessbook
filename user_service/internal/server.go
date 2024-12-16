package server

import (
	"user_service/internal/config"
	"user_service/internal/db"
	server_grpc "user_service/internal/grpc"
	"user_service/internal/services"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	Config      *config.Config
	Routes      *gin.Engine
	AuthService services.AuthService
	GRPCServer  *server_grpc.GRPCServer
}

func NewServer() (*Server, error) {
	server := &Server{}

	var err error
	server.Config, err = config.GetConfig()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	repository, err := db.NewConnect(server.Config)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	authService, err := services.NewAuthService(repository, server.Config.Token_TTL)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	server.GRPCServer = server_grpc.New(server.Config, authService)

	//server.Routes = handlers.InitRoutes(repository)

	return server, nil
}

func (server *Server) Run() {
	//server.Routes.Run(server.Config.ServerAddress)
	server.GRPCServer.Run()
}
