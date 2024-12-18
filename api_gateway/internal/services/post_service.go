package services

import (
	grpc "api_gateway/internal/grpc/clients"
	"api_gateway/internal/models"
)

type postService struct {
	grpcClient *grpc.PostClient
}

type PostService interface {
	Create(post *models.Post) (id string, err error)
	PostById(postId string) (post *models.Post, err error)
	ListPosts() (posts []*models.Post, err error)
}

func NewPostService(grpcClient *grpc.PostClient) (PostService, error) {
	return &postService{
		grpcClient: grpcClient,
	}, nil
}

func (s *postService) Create(post *models.Post) (string, error) {
	return s.grpcClient.Create(post)
}

func (s *postService) PostById(postId string) (*models.Post, error) {
	return s.grpcClient.PostById(postId)
}

func (s *postService) ListPosts() ([]*models.Post, error) {
	return s.grpcClient.ListPosts()
}
