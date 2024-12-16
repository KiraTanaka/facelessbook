package models

import "time"

type PushNotification struct {
	Id          string    `json:"id" db:"id" binding:"max=36"`
	CreatedTime time.Time `json:"created_time" db:"created_time" binding:"required"`
	UserId      string    `json:"user_id" db:"user_id" binding:"required,max=36"`
	Subject     string    `json:"subject" db:"subject" binding:"required"`
	PushMessege string    `json:"pushMessege" db:"pushMessege" binding:"required"`
	FullMessage string    `json:"fullMessage" db:"fullMessage" binding:"required"`
	IsReviewed  bool      `json:"is_reviewed" db:"is_reviewed" binding:"required"`
}
