package db

import (
	_ "embed"
	"time"
)

type Post struct {
	Id          string    `json:"id" db:"id" binding:"max=36"`
	CreatedTime time.Time `json:"created_time" db:"created_time" binding:"required"`
	AuthorId    string    `json:"author_id" db:"author_id" binding:"required,max=36"`
	Text        string    `json:"text" db:"text" binding:"required"`
}

//go:embed queries/getListPosts.sql
var getListPostsQuery string

//go:embed queries/getPost.sql
var getPostQuery string

func (r *Repository) GetListPosts() ([]Post, error) {
	posts := []Post{}
	err := r.db.Select(&posts, getListPostsQuery)
	return posts, err
}
func (r *Repository) GetPost(postId string) (*Post, error) {
	post := &Post{}
	err := r.db.Get(post, getPostQuery, postId)
	return post, err
}
