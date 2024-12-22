package user

import (
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
