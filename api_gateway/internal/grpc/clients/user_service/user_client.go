package user

import (
	"google.golang.org/grpc"

	pb "github.com/KiraTanaka/facelessbook_protos/gen/user"
)

type UserClient struct {
	Api pb.UserClient
}

func NewUserClient(conn *grpc.ClientConn) (*UserClient, error) {
	client := pb.NewUserClient(conn)

	return &UserClient{Api: client}, nil
}
