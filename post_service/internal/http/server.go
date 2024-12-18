package http

import (
	grpc "post_service/internal/clients/assessment_service"
	"post_service/internal/config"
	"post_service/internal/db"
	"post_service/internal/handlers"
	"post_service/internal/services"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	Config *config.Config
	Routes *gin.Engine
}

func NewServer() (*Server, error) {
	server := &Server{}

	var err error
	server.Config, err = config.GetAppConfig()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	repository, err := db.NewConnect(server.Config)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	grpcClient, err := grpc.NewClient(server.Config)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	postService, err := services.Init(repository, grpcClient)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	server.Routes = handlers.InitRoutes(postService)

	return server, nil
}

func (server *Server) Run() {
	server.Routes.Run(server.Config.ServerAddress)
}
