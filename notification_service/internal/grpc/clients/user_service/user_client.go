package user

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"

	pb "github.com/KiraTanaka/facelessbook_protos/gen/user"
)

type UserClient struct {
	Api     pb.UserClient
	Timeout time.Duration
}

func NewUserClient(conn *grpc.ClientConn, timeout time.Duration) *UserClient {
	client := pb.NewUserClient(conn)
	return &UserClient{
		Api:     client,
		Timeout: timeout}
}

func (c *UserClient) Nickname(userId string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.Timeout)
	defer cancel()
	r, err := c.Api.NickName(ctx, &pb.NickNameRequest{Id: userId})
	if err != nil {
		return "", fmt.Errorf("get nickname: %v", err)
	}
	return r.Nickname, nil
}
