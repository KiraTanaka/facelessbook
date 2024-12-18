package grpc

import (
	"api_gateway/internal/models"
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/KiraTanaka/facelessbook_protos/gen/post"
	log "github.com/sirupsen/logrus"
)

type PostClient struct {
	Api pb.PostClient
}

func NewPostClient(host string, port int) (*PostClient, error) {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", host, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect to grpc server: %v", err)
	}

	postClient := pb.NewPostClient(conn)

	return &PostClient{Api: postClient}, nil
}

func (c *PostClient) Create(post *models.Post) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	postMessage := &pb.PostMessage{
		AuthorId: post.AuthorId,
		Text:     post.Text}

	r, err := c.Api.Create(ctx, &pb.CreateRequest{Post: postMessage})
	if err != nil {
		log.Error(err)
		return "", err
	}
	return r.Id, nil
}

func (c *PostClient) PostById(postId string) (*models.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Api.PostById(ctx, &pb.PostByIdRequest{Id: postId})
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return PostMessageToPostModel(r.Post)
}

func (c *PostClient) ListPosts() ([]*models.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Api.ListPosts(ctx, &pb.ListPostsRequest{})
	if err != nil {
		log.Error(err)
		return nil, err
	}

	listPosts := []*models.Post{}

	for _, postMessage := range r.Posts {

		post, _ := PostMessageToPostModel(postMessage)
		listPosts = append(listPosts, post)
	}

	return listPosts, nil
}

func PostMessageToPostModel(postMessage *pb.PostMessage) (*models.Post, error) {
	createdTime, err := time.Parse(time.DateTime, postMessage.CreatedTime)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &models.Post{
		Id:          postMessage.Id,
		CreatedTime: createdTime,
		AuthorId:    postMessage.AuthorId,
		Text:        postMessage.Text}, nil

}
