package grpc

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"user_service/internal/services"

	pb "github.com/KiraTanaka/facelessbook_protos/gen/subscriber"
)

type subscriberServer struct {
	pb.UnimplementedSubscriberServer
	subscriberService services.SubscriberService
}

func RegisterSubscriberServer(gRPCServer *grpc.Server, subscriberService services.SubscriberService) {
	pb.RegisterSubscriberServer(gRPCServer, &subscriberServer{subscriberService: subscriberService})
}

func (s *subscriberServer) ListSubscribers(ctx context.Context, request *pb.ListSubscribersRequest) (*pb.ListSubscribersResponse, error) {
	subscriberIds, err := s.subscriberService.ListSubscribers(request.PublisherId)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed get list of subscribers")
	}

	return &pb.ListSubscribersResponse{Subscribers: subscriberIds}, nil
}
