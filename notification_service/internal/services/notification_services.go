package services

import (
	grpc "notification_service/internal/grpc/clients/user_service"
	"notification_service/internal/models"

	log "github.com/sirupsen/logrus"
)

type notificationService struct {
	pushService          PushService
	userGrpcClient       *grpc.UserClient
	subscriberGrpcClient *grpc.SubscriberClient
}

type NotificationService interface {
	ProcessNewPost(message *models.NewPostMessage) error
}

func NewNotificationService(pushService PushService, userGrpcClient *grpc.UserClient, subscriberGrpcClient *grpc.SubscriberClient) (NotificationService, error) {
	return &notificationService{
		pushService:          pushService,
		userGrpcClient:       userGrpcClient,
		subscriberGrpcClient: subscriberGrpcClient}, nil
}
func (s *notificationService) ProcessNewPost(message *models.NewPostMessage) error {
	nick, err := s.userGrpcClient.Nickname(message.AuthorId)
	if err != nil {
		log.Error(err)
		return err
	}

	subscriberIds, err := s.subscriberGrpcClient.ListSubscribers(message.AuthorId)
	if err != nil {
		log.Error(err)
		return err
	}

	if err = s.pushService.Notify(subscriberIds, "new_post", nick); err != nil {
		log.Error(err)
		return err
	}

	return nil
}
