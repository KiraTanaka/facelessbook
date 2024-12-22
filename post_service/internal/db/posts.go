package db

import (
	_ "embed"
	"fmt"
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
	if err := r.db.Select(&posts, getListPostsQuery); err != nil {
		return nil, fmt.Errorf("select list posts: %w", err)
	}
	return posts, nil
}
func (r *Repository) PostById(postId string) (*models.Post, error) {
	post := &models.Post{}
	if err := r.db.Get(post, getPostQuery, postId); err != nil {
		return nil, fmt.Errorf("get post by id: %w", err)
	}
	return post, nil
}

func (r *Repository) CreatePost(post *models.Post) (string, error) {
	var postId string
	tx, err := r.db.Beginx()
	if err != nil {
		return "", fmt.Errorf("begin a transaction for create post: %w", err)
	}
	defer tx.Rollback()
	if err = tx.QueryRow(createPostQuery, post.AuthorId, post.Text).Scan(&postId); err != nil {
		return "", fmt.Errorf("create post: %w", err)
	}
	tx.Commit()
	return postId, nil
}

func (r *Repository) UpdatePost(postId string, newText string) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return fmt.Errorf("begin a transaction for update post: %w", err)
	}
	defer tx.Rollback()
	if _, err = tx.Exec(updatePostQuery, postId, newText); err != nil {
		return fmt.Errorf("update post: %w", err)
	}
	tx.Commit()
	return nil
}

func (r *Repository) DeletePost(postId string) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return fmt.Errorf("begin a transaction for delete post: %w", err)
	}
	defer tx.Rollback()
	if _, err = tx.Exec(deletePostQuery, postId); err != nil {
		return fmt.Errorf("delete post: %w", err)
	}
	tx.Commit()
	return nil
}
