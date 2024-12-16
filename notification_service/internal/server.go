package server

import (
	"notification_service/internal/config"
	"notification_service/internal/db"
	server_grpc "notification_service/internal/grpc"
	"notification_service/internal/handlers"
	"notification_service/internal/services"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	Config              *config.Config
	Routes              *gin.Engine
	NotificationService services.NotificationService
	GRPCServer          *server_grpc.GRPCServer
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
	notificationService, err := services.NewNotificationService(repository)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	server.GRPCServer = server_grpc.New(server.Config, notificationService)

	server.Routes = handlers.InitRoutes(repository)

	return server, nil
}

func (server *Server) Run() {
	server.Routes.Run(server.Config.ServerAddress)
}
