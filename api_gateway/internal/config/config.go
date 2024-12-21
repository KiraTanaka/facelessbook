package config

import (
	"fmt"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	HttpConfig *HttpConfig
	DbConfig   *DbConfig
	GrpcConfig *GrpcConfig
	/*HttpServerAddress string        `env:"SERVER_ADDRESS" env-required:"true"`
	DbHost            string        `env:"POSTGRES_HOST" env-required:"true"`
	DbPort            int           `env:"POSTGRES_PORT" env-required:"true"`
	Dbname            string        `env:"POSTGRES_DATABASE" env-required:"true"`
	DbUser            string        `env:"POSTGRES_USERNAME" env-required:"true"`
	DbPassword        string        `env:"POSTGRES_PASSWORD" env-required:"true"`
	GrpcHost          string        `env:"GRPC_HOST" env-required:"true"`
	GrpcPort          int           `env:"GRPC_PORT" env-required:"true"`
	GrpcTimeout       time.Duration `env:"GRPC_TIMEOUT" env-required:"true"`*/
}

type HttpConfig struct {
	ServerAddress string `env:"HTTP_SERVER_ADDRESS" env-required:"true"`
}

type DbConfig struct {
	Host     string `env:"DB_HOST" env-required:"true"`
	Port     int    `env:"DB_PORT" env-required:"true"`
	Dbname   string `env:"DB_DATABASE" env-required:"true"`
	User     string `env:"DB_USERNAME" env-required:"true"`
	Password string `env:"DB_PASSWORD" env-required:"true"`
}

type GrpcConfig struct {
	UserHost string        `env:"GRPC_USER_HOST" env-required:"true"`
	UserPort int           `env:"GRPC_USER_PORT" env-required:"true"`
	PostHost string        `env:"GRPC_POST_HOST" env-required:"true"`
	PostPort int           `env:"GRPC_POST_PORT" env-required:"true"`
	Timeout  time.Duration `env:"GRPC_TIMEOUT" env-required:"true"`
}

func GetAppConfig() (*Config, error) {

	log.Info("starting to read application config")

	httpConfig := &HttpConfig{}
	dbConfig := &DbConfig{}
	grpcConfig := &GrpcConfig{}
	if err := cleanenv.ReadConfig(".env", httpConfig); err != nil {
		return nil, fmt.Errorf("read http config error: %w", err)
	}
	if err := cleanenv.ReadConfig(".env", dbConfig); err != nil {
		return nil, fmt.Errorf("read db config error: %w", err)
	}
	if err := cleanenv.ReadConfig(".env", grpcConfig); err != nil {
		return nil, fmt.Errorf("read grpc config error: %w", err)
	}
	config := &Config{
		HttpConfig: httpConfig,
		DbConfig:   dbConfig,
		GrpcConfig: grpcConfig}

	return config, nil
}
