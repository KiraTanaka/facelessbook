package config

import (
	"fmt"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	DbConfig    *DbConfig
	GrpcConfig  *GrpcConfig
	KafkaConfig *KafkaConfig
}

type DbConfig struct {
	Host     string `env:"DB_HOST" env-required:"true"`
	Port     int    `env:"DB_PORT" env-required:"true"`
	Dbname   string `env:"DB_DATABASE" env-required:"true"`
	User     string `env:"DB_USERNAME" env-required:"true"`
	Password string `env:"DB_PASSWORD" env-required:"true"`
}

type GrpcConfig struct {
	GrpcUserHost string        `env:"GRPC_USER_HOST" env-required:"true"`
	GrpcUserPort int           `env:"GRPC_USER_PORT" env-required:"true"`
	Timeout      time.Duration `env:"GRPC_TIMEOUT" env-required:"true"`
}
type KafkaConfig struct {
	Host string `env:"KAFKA_HOST" env-required:"true"`
	Port int    `env:"KAFKA_POST" env-required:"true"`
}

func GetAppConfig() (*Config, error) {
	dbConfig := &DbConfig{}
	grpcConfig := &GrpcConfig{}
	kafkaConfig := &KafkaConfig{}
	if err := cleanenv.ReadConfig(".env", dbConfig); err != nil {
		return nil, fmt.Errorf("read db config error: %w", err)
	}
	if err := cleanenv.ReadConfig(".env", grpcConfig); err != nil {
		return nil, fmt.Errorf("read grpc config error: %w", err)
	}
	if err := cleanenv.ReadConfig(".env", kafkaConfig); err != nil {
		return nil, fmt.Errorf("read kafka config error: %w", err)
	}

	return &Config{
		DbConfig:    dbConfig,
		GrpcConfig:  grpcConfig,
		KafkaConfig: kafkaConfig}, nil
}
