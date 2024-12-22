package services

import (
	grpc "api_gateway/internal/grpc/clients/user_service"

	log "github.com/sirupsen/logrus"
)

type subscriberService struct {
	grpcClient *grpc.SubscriberClient
}

type SubscriberService interface {
	Subscribe(publisher_id, subscriber_id string) error
	Unsubscribe(publisher_id, subscriber_id string) error
	ListSubscribers(publisher_id string) ([]string, error)
}

func NewSubscriberService(grpcClient *grpc.SubscriberClient) SubscriberService {
	return &subscriberService{
		grpcClient: grpcClient}
}

func (s *subscriberService) Subscribe(publisher_id, subscriber_id string) error {
	err := s.grpcClient.Subscribe(publisher_id, subscriber_id)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (s *subscriberService) Unsubscribe(publisher_id, subscriber_id string) error {
	err := s.grpcClient.Unsubscribe(publisher_id, subscriber_id)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (s *subscriberService) ListSubscribers(publisher_id string) ([]string, error) {
	subscribers, err := s.grpcClient.ListSubscribers(publisher_id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return subscribers, nil
}
