package server

import (
	"facelessbook/post_service/app/internal/config"
	"facelessbook/post_service/app/internal/db"
	"facelessbook/post_service/app/internal/handlers"

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

	server.Routes = handlers.InitRoutes(repository)

	return server, nil
}

func (server *Server) Run() {
	server.Routes.Run(server.Config.ServerAddress)
}
