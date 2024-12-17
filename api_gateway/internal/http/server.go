package http

import (
	"api_gateway/internal/config"
	"api_gateway/internal/services"

	"github.com/gin-gonic/gin"
)

type Server struct {
	config *config.HttpConfig
	routes *gin.Engine
}

func NewServer(config *config.HttpConfig, services *services.Services) (*Server, error) {
	server := &Server{config: config}

	server.routes = InitRoutes(services)

	return server, nil
}

func (s *Server) Run() {
	s.routes.Run(s.config.ServerAddress)
}
