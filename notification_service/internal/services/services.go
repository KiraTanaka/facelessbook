package services

import (
	"notification_service/internal/db"
	grpc "notification_service/internal/grpc/clients"
)

type Services struct {
	NotificationService NotificationService
	Push                PushService
}

func New(repository *db.Repository, grpcClients *grpc.Clients) *Services {
	pushService := NewPushService(repository)
	notificationService := NewNotificationService(pushService, grpcClients.User, grpcClients.Subscriber)

	return &Services{
		NotificationService: notificationService,
		Push:                pushService}
}
