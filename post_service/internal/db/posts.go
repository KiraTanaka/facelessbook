package db

import (
	_ "embed"
	"post_service/internal/models"
)

//go:embed queries/getListPosts.sql
var getListPostsQuery string

//go:embed queries/getPost.sql
var getPostQuery string

func (r *Repository) GetListPosts() ([]models.Post, error) {
	posts := []models.Post{}
	err := r.db.Select(&posts, getListPostsQuery)
	return posts, err
}
func (r *Repository) GetPost(postId string) (*models.Post, error) {
	post := &models.Post{}
	err := r.db.Get(post, getPostQuery, postId)
	return post, err
}
