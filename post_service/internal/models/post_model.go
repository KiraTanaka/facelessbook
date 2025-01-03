package models

import "time"

type Post struct {
	Id          string    `json:"id" db:"id" binding:"max=36"`
	CreatedTime time.Time `json:"created_time" db:"created_time"`
	AuthorId    string    `json:"author_id" db:"author_id" binding:"required,max=36"`
	Text        string    `json:"text" db:"text" binding:"required"`
}

type PostDTO struct {
	Id          string    `json:"id" binding:"max=36"`
	CreatedTime time.Time `json:"created_time"  binding:"required"`
	AuthorId    string    `json:"author_id"  binding:"required,max=36"`
	Text        string    `json:"text"  binding:"required"`
	LikeCount   int       `json:"like_count"  binding:"required"`
}
