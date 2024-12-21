package services

import (
	"user_service/internal/config"
	"user_service/internal/db"
)

type Services struct {
	Auth       AuthService
	User       UserService
	Subscriber SubscriberService
}

func New(repository *db.Repository, tokenconfig *config.TokenConfig) *Services {
	authService := NewAuthService(repository, tokenconfig)
	userService := NewUserService(repository)
	subscriberService := NewSubscriberService(repository)

	return &Services{
		Auth:       authService,
		User:       userService,
		Subscriber: subscriberService}
}
