package services

import (
	"assessment_service/internal/db"

	log "github.com/sirupsen/logrus"
)

type postLikeService struct {
	repository *db.Repository
}

func NewPostLikeService(repository *db.Repository) (PostLikeService, error) {
	return &postLikeService{
		repository: repository,
	}, nil
}

type PostLikeService interface {
	GetLikeCount(
		post_id string,
	) (cnt int, err error)
}

func (s *postLikeService) GetLikeCount(postId string) (int, error) {
	postLikeCount, err := s.repository.GetPostLikeCount(postId)

	if err != nil {

		log.Error(err)
		return 0, err
	}

	return postLikeCount, err
}
