package config

import (
	"fmt"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	DbConfig    *DbConfig
	GrpcConfig  *GrpcConfig
	TokenConfig *TokenConfig
}

type DbConfig struct {
	Host     string `env:"DB_HOST" env-required:"true"`
	Port     int    `env:"DB_PORT" env-required:"true"`
	Dbname   string `env:"DB_DATABASE" env-required:"true"`
	User     string `env:"DB_USERNAME" env-required:"true"`
	Password string `env:"DB_PASSWORD" env-required:"true"`
}

type GrpcConfig struct {
	GrpcPort    int           `env:"GRPC_PORT" env-required:"true"`
	GrpcTimeout time.Duration `env:"GRPC_TIMEOUT" env-required:"true"`
}

type TokenConfig struct {
	Token_TTL time.Duration `env:"TOKEN_TTL" env-required:"true"`
}

func GetAppConfig() (*Config, error) {
	dbConfig := &DbConfig{}
	grpcConfig := &GrpcConfig{}
	tokenConfig := &TokenConfig{}
	if err := cleanenv.ReadConfig(".env", dbConfig); err != nil {
		return nil, fmt.Errorf("read db config error: %w", err)
	}
	if err := cleanenv.ReadConfig(".env", grpcConfig); err != nil {
		return nil, fmt.Errorf("read grpc config error: %w", err)
	}
	if err := cleanenv.ReadConfig(".env", tokenConfig); err != nil {
		return nil, fmt.Errorf("read token config error: %w", err)
	}

	return &Config{
		DbConfig:    dbConfig,
		GrpcConfig:  grpcConfig,
		TokenConfig: tokenConfig}, nil
}
