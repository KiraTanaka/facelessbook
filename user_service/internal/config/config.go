package config

import (
	"fmt"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	ServerAddress string        `env:"SERVER_ADDRESS" env-required:"true"`
	Host          string        `env:"POSTGRES_HOST" env-required:"true"`
	Port          int           `env:"POSTGRES_PORT" env-required:"true"`
	Dbname        string        `env:"POSTGRES_DATABASE" env-required:"true"`
	User          string        `env:"POSTGRES_USERNAME" env-required:"true"`
	Password      string        `env:"POSTGRES_PASSWORD" env-required:"true"`
	GrpcPort      int           `env:"GRPC_PORT" env-required:"true"`
	GrpcTimeout   time.Duration `env:"GRPC_TIMEOUT" env-required:"true"`
	Token_TTL     time.Duration `env:"TOKEN_TTL" env-required:"true"`
}

func GetConfig() (*Config, error) {
	config := &Config{}

	log.Info("Read application config")

	err := cleanenv.ReadConfig(".env", config)

	if err != nil {
		return nil, fmt.Errorf("server config error: %w", err)
	}

	log.Info("Reading of the server configuration parameters is completed.")
	return config, nil
}
