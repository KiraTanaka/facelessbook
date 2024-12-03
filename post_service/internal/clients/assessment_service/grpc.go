package grpc

import (
	"context"
	"fmt"
	"post_service/internal/config"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	assessment "github.com/KiraTanaka/facelessbook_protos/gen/assessment"
	log "github.com/sirupsen/logrus"
)

type Client struct {
	AssessmentApi assessment.PostLikesClient
}

func NewClient(config *config.Config) (*Client, error) {
	// Set up a connection to the server.
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", config.GrpcHost, config.GrpcPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect to grpc server: %v", err)
	}

	grpcClient := assessment.NewPostLikesClient(conn)

	return &Client{AssessmentApi: grpcClient}, nil
}

func (c *Client) GetLikeCount(postId string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.AssessmentApi.GetLikeCount(ctx, &assessment.GetLikeCountRequest{PostId: postId})
	if err != nil {
		log.Error(err)
		return 0, err
	}
	return int(r.Cnt), nil
}
