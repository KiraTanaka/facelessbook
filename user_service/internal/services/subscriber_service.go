package services

import (
	"user_service/internal/db"

	log "github.com/sirupsen/logrus"
)

type subscriberService struct {
	repository *db.Repository
}

type SubscriberService interface {
	ListSubscribers(publisherId string) ([]string, error)
}

func NewSubscriberService(repository *db.Repository) SubscriberService {
	return &subscriberService{
		repository: repository}
}

func (s *subscriberService) ListSubscribers(publisherId string) ([]string, error) {
	subscriberIds, err := s.repository.ListSubscribers(publisherId)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return subscriberIds, err
}
