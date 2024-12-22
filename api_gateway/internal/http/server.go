package http

import (
	"api_gateway/internal/config"
	"api_gateway/internal/services"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	config *config.HttpConfig
	routes *gin.Engine
}

func NewServer(config *config.HttpConfig, services *services.Services) *Server {
	server := &Server{config: config}

	server.routes = NewRoutes(services)

	return server
}

func (s *Server) Run() {
	log.Info("Runs http server.")
	s.routes.Run(s.config.ServerAddress)
}
