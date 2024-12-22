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

func (c *SubscriberClient) Subscribe(publisher_id, subscriber_id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.Timeout)
	defer cancel()
	r, err := c.Api.Subscribe(ctx, &pb.SubscribeRequest{PublisherId: publisher_id, SubscriberId: subscriber_id})
	if err != nil {
		return fmt.Errorf("subdcribe: %w", err)
	} else if !r.Success {
		return fmt.Errorf("subdcribe ended unsuccessfully")
	}
	return nil
}
func (c *SubscriberClient) Unsubscribe(publisher_id, subscriber_id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.Timeout)
	defer cancel()
	r, err := c.Api.Unsubscribe(ctx, &pb.UnsubscribeRequest{PublisherId: publisher_id, SubscriberId: subscriber_id})
	if err != nil {
		return fmt.Errorf("unsubscribe: %w", err)
	} else if !r.Success {
		return fmt.Errorf("unsubscribe ended unsuccessfully")
	}

	return nil
}
func (c *SubscriberClient) ListSubscribers(publisher_id string) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.Timeout)
	defer cancel()
	r, err := c.Api.ListSubscribers(ctx, &pb.ListSubscribersRequest{PublisherId: publisher_id})
	if err != nil {
		return nil, fmt.Errorf("get list of subdcribers: %w", err)
	}

	return r.Subscribers, nil
}
