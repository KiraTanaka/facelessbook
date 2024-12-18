package services

import "post_service/internal/db"

type Services struct {
	Post PostService
}

func Init(repository *db.Repository) (*Services, error) {
	postService, err := NewPostService(repository)
	if err != nil {
		return nil, err
	}

	return &Services{
		Post: postService,
	}, nil
}
