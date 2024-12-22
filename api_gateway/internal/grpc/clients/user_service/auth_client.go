package user

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"

	auth "github.com/KiraTanaka/facelessbook_protos/gen/auth"
)

type AuthClient struct {
	Api     auth.AuthClient
	Timeout time.Duration
}

func NewAuthClient(conn *grpc.ClientConn, timeout time.Duration) *AuthClient {
	client := auth.NewAuthClient(conn)

	return &AuthClient{
		Api:     client,
		Timeout: timeout}
}

func (c *AuthClient) Register(phone string, password string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.Timeout)
	defer cancel()
	r, err := c.Api.Register(ctx, &auth.RegisterRequest{Phone: phone, Password: password})
	if err != nil {
		return "", fmt.Errorf("user register: %w", err)
	}
	return r.UserId, nil
}
func (c *AuthClient) Login(phone string, password string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.Timeout)
	defer cancel()
	r, err := c.Api.Login(ctx, &auth.LoginRequest{Phone: phone, Password: password})
	if err != nil {
		return "", fmt.Errorf("user login: %w", err)
	}

	return r.Token, nil
}
