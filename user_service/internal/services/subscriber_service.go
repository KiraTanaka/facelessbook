package services

import (
	"user_service/internal/db"

	log "github.com/sirupsen/logrus"
)

type subscriberService struct {
	repository *db.Repository
}

type SubscriberService interface {
	Subscribe(publisher_id, subscriber_id string) error
	Unsubscribe(publisher_id, subscriber_id string) error
	ListSubscribers(publisherId string) ([]string, error)
}

func NewSubscriberService(repository *db.Repository) SubscriberService {
	return &subscriberService{
		repository: repository}
}

func (s *subscriberService) Subscribe(publisher_id, subscriber_id string) error {
	err := s.repository.Subscribe(publisher_id, subscriber_id)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (s *subscriberService) Unsubscribe(publisher_id, subscriber_id string) error {
	err := s.repository.Unsubscribe(publisher_id, subscriber_id)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (s *subscriberService) ListSubscribers(publisherId string) ([]string, error) {
	subscriberIds, err := s.repository.ListSubscribers(publisherId)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return subscriberIds, err
}
