package app

import (
	"post_service/internal/broker"
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
	log.Info("Finished reading the configuration.")

	repository, err := db.NewConnect(config.DbConfig)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	log.Info("Connected to the database.")

	writer := broker.NewConnect(config.KafkaConfig)
	log.Info("Created the writers for the broker.")

	services := services.New(repository, writer)
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
