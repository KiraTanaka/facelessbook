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
	if err := validateListSubscribersRequest(request); err != nil {
		return nil, err
	}

	subscriberIds, err := s.subscriberService.ListSubscribers(request.PublisherId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.ListSubscribersResponse{Subscribers: subscriberIds}, nil
}

func validateListSubscribersRequest(request *pb.ListSubscribersRequest) error {
	if request.PublisherId == "" {
		return status.Error(codes.InvalidArgument, "publisher_id is required")
	}
	return nil

}
