package server

import (
	"assessment_service/internal/config"
	"assessment_service/internal/db"
	server_grpc "assessment_service/internal/grpc"
	"assessment_service/internal/handlers"
	"assessment_service/internal/services"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	Config          *config.Config
	Routes          *gin.Engine
	PostLikeService *services.PostLikeService
	GRPCServer      *server_grpc.GRPCServer
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
	postLikeService, err := services.NewPostLikeService(repository)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	server.GRPCServer = server_grpc.New(server.Config, postLikeService)

	server.Routes = handlers.InitRoutes(repository)

	return server, nil
}

func (server *Server) Run() {
	server.Routes.Run(server.Config.ServerAddress)
}
