package services

import (
	"notification_service/internal/db"

	log "github.com/sirupsen/logrus"
)

type pushService struct {
	repository *db.Repository
}

type PushService interface {
	Notify(userIds []string, patternName string, params ...any) error
}

func NewPushService(repository *db.Repository) (PushService, error) {
	return &pushService{
		repository: repository}, nil
}
func (s *pushService) Notify(userIds []string, patternName string, params ...any) error {
	err := s.repository.SavePush(userIds, patternName, params)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
