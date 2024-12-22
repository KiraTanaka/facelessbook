package post

import (
	"api_gateway/internal/models"
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	pb "github.com/KiraTanaka/facelessbook_protos/gen/post"
)

type PostClient struct {
	Api     pb.PostClient
	Timeout time.Duration
}

func NewPostClient(conn *grpc.ClientConn, timeout time.Duration) *PostClient {
	postClient := pb.NewPostClient(conn)

	return &PostClient{
		Api:     postClient,
		Timeout: timeout}
}

func (c *PostClient) Create(post *models.Post) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.Timeout)
	defer cancel()

	postMessage := &pb.PostInformation{
		AuthorId: post.AuthorId,
		Text:     post.Text}

	r, err := c.Api.Create(ctx, &pb.CreateRequest{Post: postMessage})
	if err != nil {
		return "", fmt.Errorf("post create: %w", err)
	}
	return r.Id, nil
}

func (c *PostClient) PostById(postId string) (*models.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.Timeout)
	defer cancel()

	r, err := c.Api.PostById(ctx, &pb.PostByIdRequest{Id: postId})
	if err != nil {
		return nil, fmt.Errorf("get post by id: %w", err)
	}

	return PostMessageToPostModel(r.Post)
}

func (c *PostClient) ListPosts() ([]*models.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.Timeout)
	defer cancel()

	r, err := c.Api.ListPosts(ctx, &pb.ListPostsRequest{})
	if err != nil {
		return nil, fmt.Errorf("get list posts: %w", err)
	}

	listPosts := make([]*models.Post, len(r.Posts))

	for i, postMessage := range r.Posts {
		post, err := PostMessageToPostModel(postMessage)
		if err != nil {
			log.Error(err)
		}
		listPosts[i] = post
	}

	return listPosts, nil
}

func (c *PostClient) Update(postId string, newText string) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.Timeout)
	defer cancel()

	_, err := c.Api.Update(ctx, &pb.UpdateRequest{Id: postId, NewText: newText})
	if err != nil {
		return fmt.Errorf("update post: %w", err)
	}
	return nil
}

func (c *PostClient) Delete(postId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.Timeout)
	defer cancel()

	_, err := c.Api.Delete(ctx, &pb.DeleteRequest{Id: postId})
	if err != nil {
		return fmt.Errorf("delete post: %w", err)
	}
	return nil
}

func PostMessageToPostModel(postMessage *pb.PostInformation) (*models.Post, error) {
	createdTime, err := time.Parse(time.DateTime, postMessage.CreatedTime)
	if err != nil {
		return nil, fmt.Errorf("parse createdTime from post message: %w", err)
	}

	return &models.Post{
		Id:          postMessage.Id,
		CreatedTime: createdTime,
		AuthorId:    postMessage.AuthorId,
		Text:        postMessage.Text}, nil

}
