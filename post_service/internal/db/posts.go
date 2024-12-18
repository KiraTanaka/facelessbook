package db

import (
	_ "embed"
	"post_service/internal/models"
)

//go:embed queries/getListPosts.sql
var getListPostsQuery string

//go:embed queries/getPost.sql
var getPostQuery string

//go:embed queries/createPost.sql
var createPostQuery string

//go:embed queries/updatePost.sql
var updatePostQuery string

//go:embed queries/deletePost.sql
var deletePostQuery string

func (r *Repository) ListPosts() ([]*models.Post, error) {
	posts := []*models.Post{}
	err := r.db.Select(&posts, getListPostsQuery)
	return posts, err
}
func (r *Repository) PostById(postId string) (*models.Post, error) {
	post := &models.Post{}
	err := r.db.Get(post, getPostQuery, postId)
	return post, err
}

func (r *Repository) CreatePost(post *models.Post) (string, error) {
	var postId string
	tx, err := r.db.Beginx()
	if err != nil {
		return "", err
	}
	defer tx.Rollback()
	err = tx.QueryRow(createPostQuery, post.AuthorId, post.Text).Scan(&postId)
	if err != nil {
		return "", err
	}
	tx.Commit()
	return postId, nil
}

func (r *Repository) UpdatePost(postId string, newText string) (string, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return "", err
	}
	defer tx.Rollback()
	_, err = tx.Exec(updatePostQuery, postId, newText)
	if err != nil {
		return "", err
	}
	tx.Commit()
	return newText, nil
}

func (r *Repository) DeletePost(postId string) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	_, err = tx.Exec(deletePostQuery, postId)
	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}
