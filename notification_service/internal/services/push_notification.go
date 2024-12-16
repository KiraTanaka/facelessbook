package services

import (
	"notification_service/internal/db"

	log "github.com/sirupsen/logrus"
)

type notificationService struct {
	repository *db.Repository
}

func NewNotificationService(repository *db.Repository) (NotificationService, error) {
	return &notificationService{
		repository: repository,
	}, nil
}

type NotificationService interface {
	GetLikeCount(
		post_id string,
	) (cnt int, err error)
}

func (s *notificationService) GetLikeCount(postId string) (int, error) {
	postLikeCount, err := s.repository.GetPostLikeCount(postId)

	if err != nil {

		log.Error(err)
		return 0, err
	}

	return postLikeCount, err
}
