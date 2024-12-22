package services

import (
	"fmt"
	"time"
	"user_service/internal/config"
	"user_service/internal/db"
	"user_service/internal/jwt"

	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	repository *db.Repository
	tokenTTL   time.Duration
}

type AuthService interface {
	Register(phone string, password string) (user_id string, err error)
	Login(phone string, password string) (toker string, err error)
}

func NewAuthService(repository *db.Repository, tokenConfig *config.TokenConfig) AuthService {
	return &authService{
		repository: repository,
		tokenTTL:   tokenConfig.Token_TTL,
	}
}

func (s *authService) Register(phone string, password string) (string, error) {
	passHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Error(fmt.Errorf("generate the bcrypt hash of the password: %w", err))
		return "", err
	}

	userId, err := s.repository.CreateUser(phone, passHash)
	if err != nil {
		log.Error(err)
		return "", err
	}

	return userId, err
}

func (s *authService) Login(phone string, password string) (string, error) {
	user, err := s.repository.User(phone)
	if err != nil {
		log.Error(err)
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword(user.PassHash, []byte(password)); err != nil {
		log.Error(fmt.Errorf("compares a bcrypt hashed password: %w", err))
		return "", err
	}

	token, err := jwt.NewToken(user, s.tokenTTL)
	if err != nil {
		log.Error(err)
		return "", err
	}

	return token, err
}
