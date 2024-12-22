package user

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"

	pb "github.com/KiraTanaka/facelessbook_protos/gen/subscriber"
)

type SubscriberClient struct {
	Api     pb.SubscriberClient
	Timeout time.Duration
}

func NewSubscriberClient(conn *grpc.ClientConn, timeout time.Duration) *SubscriberClient {
	client := pb.NewSubscriberClient(conn)

	return &SubscriberClient{
		Api:     client,
		Timeout: timeout}
}

func (c *SubscriberClient) ListSubscribers(publisherId string) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.Timeout)
	defer cancel()
	r, err := c.Api.ListSubscribers(ctx, &pb.ListSubscribersRequest{PublisherId: publisherId})
	if err != nil {
		return nil, fmt.Errorf("get list of subscribers: %v", err)
	}
	return r.Subscribers, nil
}
