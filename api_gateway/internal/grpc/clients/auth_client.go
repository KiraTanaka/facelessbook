package grpc

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	auth "github.com/KiraTanaka/facelessbook_protos/gen/auth"
	log "github.com/sirupsen/logrus"
)

type AuthClient struct {
	Api auth.AuthClient
}

func NewAuthClient(host string, port int) (*AuthClient, error) {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", host, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect to grpc server: %v", err)
	}

	authClient := auth.NewAuthClient(conn)

	return &AuthClient{Api: authClient}, nil
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
