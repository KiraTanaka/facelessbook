package app

import (
	"api_gateway/internal/config"
	grpc "api_gateway/internal/grpc/clients/user_service"
	"api_gateway/internal/http"
	"api_gateway/internal/services"

	log "github.com/sirupsen/logrus"
)

type App struct {
	config      *config.Config
	httpServer  *http.Server
	services    *services.Services
	grpcClients *grpc.Clients
}

func New() (*App, error) {
	app := &App{}

	var err error
	app.config, err = config.GetAppConfig()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	app.grpcClients, err = grpc.NewClient(app.config.GrpcConfig)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	app.services, err = services.Init(app.grpcClients)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	app.httpServer, err = http.NewServer(app.config.HttpConfig, app.services)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return app, nil
}

func (app *App) Run() {
	app.httpServer.Run()
}
