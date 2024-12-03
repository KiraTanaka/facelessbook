package services

import (
	grpc "post_service/internal/clients/assessment_service"
	"post_service/internal/db"
	"post_service/internal/models"
)

type postService struct {
	repository *db.Repository
	grpcClient *grpc.Client
}

type PostService interface {
	GetListPosts() ([]models.Post, error)
	GetPost(string) (models.PostDTO, error)
}

func PostToDTO(post *models.Post) models.PostDTO {
	return models.PostDTO{
		Id:          post.Id,
		CreatedTime: post.CreatedTime,
		AuthorId:    post.AuthorId,
		Text:        post.Text,
	}
}

func Init(repository *db.Repository, grpcClient *grpc.Client) (PostService, error) {
	return &postService{repository: repository, grpcClient: grpcClient}, nil

}

func (s *postService) GetListPosts() ([]models.Post, error) {
	posts, err := s.repository.GetListPosts()
	if err != nil {
		return posts, err
	}
	return posts, nil
}

func (s *postService) GetPost(postId string) (models.PostDTO, error) {
	var dto models.PostDTO
	post, err := s.repository.GetPost(postId)
	if err != nil {
		return dto, err
	}
	dto = PostToDTO(post)

	likeCount, err := s.grpcClient.GetLikeCount(postId)
	if err != nil {
		return dto, err
	}
	dto.LikeCount = likeCount

	return dto, err
}
