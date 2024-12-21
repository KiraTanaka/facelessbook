package user

import (
	"context"
	"time"

	"google.golang.org/grpc"

	auth "github.com/KiraTanaka/facelessbook_protos/gen/auth"
	log "github.com/sirupsen/logrus"
)

type AuthClient struct {
	Api auth.AuthClient
}

func NewAuthClient(conn *grpc.ClientConn) (*AuthClient, error) {

	client := auth.NewAuthClient(conn)

	return &AuthClient{Api: client}, nil
}

func (c *AuthClient) Register(phone string, password string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Api.Register(ctx, &auth.RegisterRequest{Phone: phone, Password: password})
	if err != nil {
		log.Error(err)
		return "", err
	}
	return r.UserId, nil
}
func (c *AuthClient) Login(phone string, password string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Api.Login(ctx, &auth.LoginRequest{Phone: phone, Password: password})
	if err != nil {
		log.Error(err)
		return "", err
	}
	return r.Token, nil
}
