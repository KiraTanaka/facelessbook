package services

import (
	"notification_service/internal/db"
	grpc "notification_service/internal/grpc/clients"
)

type Services struct {
	NotificationService NotificationService
	Push                PushService
}

func New(repository *db.Repository, grpcClients *grpc.Clients) (*Services, error) {
	pushService, err := NewPushService(repository)
	if err != nil {
		return nil, err
	}

	notificationService, err := NewNotificationService(pushService, grpcClients.User, grpcClients.Subscriber)
	if err != nil {
		return nil, err
	}

	return &Services{
		NotificationService: notificationService,
		Push:                pushService,
	}, nil
}
