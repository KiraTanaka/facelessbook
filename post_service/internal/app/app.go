package app

import (
	"post_service/internal/config"
	"post_service/internal/db"
	grpc "post_service/internal/grpc/server"
	"post_service/internal/services"

	log "github.com/sirupsen/logrus"
)

type App struct {
	config *config.Config
	//httpServer *http.Server
	services   *services.Services
	grpcServer *grpc.Server
}

func New() (*App, error) {
	config, err := config.GetAppConfig()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	repository, err := db.NewConnect(config.DbConfig)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	services, err := services.Init(repository)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	/*app.httpServer, err = http.NewServer(app.config.HttpConfig, app.services)
	if err != nil {
		log.Error(err)
		return nil, err
	}*/

	grpcServer := grpc.NewServer(config.GrpcConfig, services)

	return &App{
		config:     config,
		services:   services,
		grpcServer: grpcServer,
	}, nil
}

func (app *App) Run() {
	app.grpcServer.Run()
}
