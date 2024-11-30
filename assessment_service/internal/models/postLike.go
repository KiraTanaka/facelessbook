package models

import "time"

type PostLike struct {
	Id          string    `json:"id" db:"id" binding:"max=36"`
	CreatedTime time.Time `json:"created_time" db:"created_time" binding:"required"`
	PostId      string    `json:"post_id" db:"post_id" binding:"required,max=36"`
	UserId      string    `json:"user_id" db:"user_id" binding:"required,max=36"`
	IsLike      bool      `json:"is_like" db:"is_like" binding:"required"`
}
