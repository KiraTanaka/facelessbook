package app

import (
	"notification_service/internal/broker"
	"notification_service/internal/config"
	"notification_service/internal/db"
	grpc "notification_service/internal/grpc/clients"
	"notification_service/internal/services"

	log "github.com/sirupsen/logrus"
)

type App struct {
	config        *config.Config
	brokerReaders *broker.Readers
	services      *services.Services
	//grpcServer *grpc.Server
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

	grpcClients, err := grpc.NewClients(config.GrpcConfig)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	log.Info("Created the new GRPC clients.")

	services := services.New(repository, grpcClients)
	log.Info("Services have been initialized.")

	brokerReaders := broker.NewConnect(config.KafkaConfig, services)
	log.Info("Created the readers for the broker.")

	//grpcServer := grpc.NewServer(config.GrpcConfig, services)

	return &App{
		config:   config,
		services: services,
		//grpcServer:    grpcServer,
		brokerReaders: brokerReaders,
	}, nil
}

func (app *App) Run() {
	//app.grpcServer.Run()
	app.brokerReaders.Run()
}
