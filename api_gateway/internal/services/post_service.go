package services

import (
	grpc "api_gateway/internal/grpc/clients"
	"api_gateway/internal/models"

	log "github.com/sirupsen/logrus"
)

type postService struct {
	grpcClient *grpc.PostClient
}

type PostService interface {
	Create(post *models.Post) (id string, err error)
	PostById(postId string) (post *models.Post, err error)
	ListPosts() (posts []*models.Post, err error)
	Update(postId string, newText string) (err error)
	Delete(postId string) (err error)
}

func NewPostService(grpcClient *grpc.PostClient) (PostService, error) {
	return &postService{
		grpcClient: grpcClient,
	}, nil
}

func (s *postService) Create(post *models.Post) (string, error) {
	postId, err := s.grpcClient.Create(post)
	if err != nil {
		log.Error(err)
		return "", err
	}
	return postId, nil
}

func (s *postService) Update(postId string, newText string) error {
	err := s.grpcClient.Update(postId, newText)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (s *postService) Delete(postId string) error {
	err := s.grpcClient.Delete(postId)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (s *postService) PostById(postId string) (*models.Post, error) {
	post, err := s.grpcClient.PostById(postId)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return post, nil
}

func (s *postService) ListPosts() ([]*models.Post, error) {
	posts, err := s.grpcClient.ListPosts()
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return posts, nil
}
