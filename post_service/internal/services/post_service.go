package services

import (
	"post_service/internal/db"
	"post_service/internal/models"

	log "github.com/sirupsen/logrus"
)

type postService struct {
	repository *db.Repository
	//grpcClient *grpc.Client
}

type PostService interface {
	Create(post *models.Post) (postId string, err error)
	PostById(postId string) (post *models.Post, err error)
	ListPosts() (posts []*models.Post, err error)
	Update(postId string, newText string) (text string, err error)
	Delete(postId string) error
}

func PostToDTO(post *models.Post) models.PostDTO {
	return models.PostDTO{
		Id:          post.Id,
		CreatedTime: post.CreatedTime,
		AuthorId:    post.AuthorId,
		Text:        post.Text,
	}
}

func NewPostService(repository *db.Repository) (PostService, error) {
	return &postService{
		repository: repository,
	}, nil

}

func (s *postService) Create(post *models.Post) (string, error) {
	postId, err := s.repository.CreatePost(post)
	if err != nil {
		log.Error(err)
		return "", err
	}
	return postId, nil
}

func (s *postService) PostById(postId string) (*models.Post, error) {
	post, err := s.repository.PostById(postId)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return post, nil
}

func (s *postService) ListPosts() ([]*models.Post, error) {
	posts, err := s.repository.ListPosts()
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return posts, nil
}

func (s *postService) Update(postId string, newText string) (string, error) {
	text, err := s.repository.UpdatePost(postId, newText)
	if err != nil {
		log.Error(err)
		return "", err
	}
	return text, nil
}

func (s *postService) Delete(postId string) error {
	err := s.repository.DeletePost(postId)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
