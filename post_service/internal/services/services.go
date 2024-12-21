package services

import (
	"post_service/internal/broker"
	"post_service/internal/db"
)

type Services struct {
	Post PostService
}

func Init(repository *db.Repository, writer broker.Writer) (*Services, error) {
	postService, err := NewPostService(repository, writer)
	if err != nil {
		return nil, err
	}

	return &Services{
		Post: postService,
	}, nil
}
