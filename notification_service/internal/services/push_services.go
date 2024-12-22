package services

import (
	"notification_service/internal/db"
)

type pushService struct {
	repository *db.Repository
}

type PushService interface {
	Notify(userIds []string, patternName string, params ...any) error
}

func NewPushService(repository *db.Repository) PushService {
	return &pushService{
		repository: repository}
}
func (s *pushService) Notify(userIds []string, patternName string, params ...any) error {
	err := s.repository.SavePush(userIds, patternName, params)
	if err != nil {
		return err
	}

	return nil
}
