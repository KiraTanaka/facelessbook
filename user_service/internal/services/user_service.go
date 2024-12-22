package services

import (
	"user_service/internal/db"

	log "github.com/sirupsen/logrus"
)

type userService struct {
	repository *db.Repository
}

type UserService interface {
	Nickname(userId string) (string, error)
}

func NewUserService(repository *db.Repository) UserService {
	return &userService{
		repository: repository}
}

func (s *userService) Nickname(userId string) (string, error) {
	nick, err := s.repository.Nickname(userId)
	if err != nil {
		log.Error(err)
		return "", err
	}

	return nick, err
}
