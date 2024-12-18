package models

import "time"

type Post struct {
	Id          string    `json:"id" binding:"max=36"`
	CreatedTime time.Time `json:"created_time"`
	AuthorId    string    `json:"author_id" binding:"required,max=36"`
	Text        string    `json:"text" binding:"required"`
}
