package broker

import (
	"encoding/json"
	"notification_service/internal/models"
)

func ToNewPostMessage(value []byte) (*models.NewPostMessage, error) {
	message := &models.NewPostMessage{}
	err := json.Unmarshal(value, message)
	if err != nil {
		return nil, err
	}
	return message, nil
}
