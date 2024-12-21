package post

import (
	"api_gateway/internal/models"
	"context"
	"time"

	"google.golang.org/grpc"

	pb "github.com/KiraTanaka/facelessbook_protos/gen/post"
)

type PostClient struct {
	Api pb.PostClient
}

func NewPostClient(conn *grpc.ClientConn) (*PostClient, error) {
	postClient := pb.NewPostClient(conn)

	return &PostClient{Api: postClient}, nil
}

func (c *PostClient) Create(post *models.Post) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	postMessage := &pb.PostInformation{
		AuthorId: post.AuthorId,
		Text:     post.Text}
	r, err := c.Api.Create(ctx, &pb.CreateRequest{Post: postMessage})
	if err != nil {
		return "", err
	}
	return r.Id, nil
}

func (c *PostClient) PostById(postId string) (*models.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Api.PostById(ctx, &pb.PostByIdRequest{Id: postId})
	if err != nil {
		return nil, err
	}

	return PostMessageToPostModel(r.Post)
}

func (c *PostClient) ListPosts() ([]*models.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Api.ListPosts(ctx, &pb.ListPostsRequest{})
	if err != nil {
		return nil, err
	}

	listPosts := []*models.Post{}

	for _, postMessage := range r.Posts {

		post, _ := PostMessageToPostModel(postMessage)
		listPosts = append(listPosts, post)
	}

	return listPosts, nil
}

func (c *PostClient) Update(postId string, newText string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := c.Api.Update(ctx, &pb.UpdateRequest{Id: postId, NewText: newText})
	if err != nil {
		return err
	}
	return nil
}

func (c *PostClient) Delete(postId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := c.Api.Delete(ctx, &pb.DeleteRequest{Id: postId})
	if err != nil {
		return err
	}
	return nil
}

func PostMessageToPostModel(postMessage *pb.PostInformation) (*models.Post, error) {
	createdTime, err := time.Parse(time.DateTime, postMessage.CreatedTime)
	if err != nil {
		return nil, err
	}

	return &models.Post{
		Id:          postMessage.Id,
		CreatedTime: createdTime,
		AuthorId:    postMessage.AuthorId,
		Text:        postMessage.Text}, nil

}
