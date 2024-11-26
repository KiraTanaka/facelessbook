package server

import (
	"facelessbook/post_service/app/internal/config"

	log "github.com/sirupsen/logrus"
)

type Server struct {
	Config *config.Config
}

func NewServer() (*Server, error) {
	server := &Server{}

	var err error
	server.Config, err = config.GetConfig()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return server, nil
}

func (server *Server) Run() {
	return
}
