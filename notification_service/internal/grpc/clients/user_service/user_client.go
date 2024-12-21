package user

import (
	"context"
	"time"

	"google.golang.org/grpc"

	pb "github.com/KiraTanaka/facelessbook_protos/gen/user"
)

type UserClient struct {
	Api     pb.UserClient
	Timeout time.Duration
}

func NewUserClient(conn *grpc.ClientConn, timeout time.Duration) *UserClient {
	return &UserClient{
		Api:     pb.NewUserClient(conn),
		Timeout: timeout}
}

func (c *UserClient) Nickname(userId string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.Timeout)
	defer cancel()
	r, err := c.Api.NickName(ctx, &pb.NickNameRequest{Id: userId})
	if err != nil {
		return "", err
	}
	return r.Nickname, nil
}
