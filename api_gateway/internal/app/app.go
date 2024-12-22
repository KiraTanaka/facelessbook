package app

import (
	"api_gateway/internal/config"
	grpc "api_gateway/internal/grpc/clients"
	"api_gateway/internal/http"
	"api_gateway/internal/services"

	log "github.com/sirupsen/logrus"
)

type App struct {
	config     *config.Config
	httpServer *http.Server
	services   *services.Services
}

func New() (*App, error) {
	config, err := config.GetAppConfig()
	if err != nil {
		log.Error(err)
		return nil, err
	}
	log.Info("Finished reading the configuration.")

	grpcClients, err := grpc.NewClients(config.GrpcConfig)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	log.Info("Created the new GRPC clients.")

	services := services.New(grpcClients)
	log.Info("Services have been initialized.")

	httpServer := http.NewServer(config.HttpConfig, services)
	log.Info("Created the new http server.")

	return &App{
		config:     config,
		services:   services,
		httpServer: httpServer,
	}, nil
}

func (app *App) Run() {
	app.httpServer.Run()
}
