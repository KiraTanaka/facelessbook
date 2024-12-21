package app

import (
	"user_service/internal/config"
	"user_service/internal/db"
	grpc "user_service/internal/grpc/server"
	"user_service/internal/services"

	log "github.com/sirupsen/logrus"
)

type App struct {
	config     *config.Config
	services   *services.Services
	grpcServer *grpc.Server
}

func New() (*App, error) {
	config, err := config.GetAppConfig()
	if err != nil {
		log.Error(err)
		return nil, err
	}
	log.Info("Finished reading the configuration.")

	repository, err := db.NewConnect(config.DbConfig)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	log.Info("Connected to the database.")

	services := services.New(repository, config.TokenConfig)
	log.Info("Created services.")

	grpcServer := grpc.NewServer(config.GrpcConfig, services)
	log.Info("Registered the GRPC server.")

	return &App{
		config:     config,
		services:   services,
		grpcServer: grpcServer,
	}, nil
}

func (app *App) Run() {
	app.grpcServer.Run()
}
