package services

import (
	"post_service/internal/broker"
	"post_service/internal/db"
)

type Services struct {
	Post PostService
}

func New(repository *db.Repository, writer broker.Writer) *Services {
	postService := NewPostService(repository, writer)

	return &Services{
		Post: postService,
	}
}
