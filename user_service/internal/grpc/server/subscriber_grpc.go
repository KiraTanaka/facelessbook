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

func (s *subscriberServer) Subscribe(ctx context.Context, request *pb.SubscribeRequest) (*pb.SubscribeResponse, error) {
	if err := validateSubscribeRequest(request); err != nil {
		return nil, err
	}

	err := s.subscriberService.Subscribe(request.PublisherId, request.SubscriberId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.SubscribeResponse{Success: true}, nil
}

func (s *subscriberServer) Unsubscribe(ctx context.Context, request *pb.UnsubscribeRequest) (*pb.UnsubscribeResponse, error) {
	if err := validateUnsubscribeRequest(request); err != nil {
		return nil, err
	}

	err := s.subscriberService.Unsubscribe(request.PublisherId, request.SubscriberId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.UnsubscribeResponse{Success: true}, nil
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

func validateSubscribeRequest(request *pb.SubscribeRequest) error {
	if request.PublisherId == "" {
		return status.Error(codes.InvalidArgument, "publisher_id is required")
	}
	if request.SubscriberId == "" {
		return status.Error(codes.InvalidArgument, "subscriber_id is required")
	}
	return nil

}

func validateUnsubscribeRequest(request *pb.UnsubscribeRequest) error {
	if request.PublisherId == "" {
		return status.Error(codes.InvalidArgument, "publisher_id is required")
	}
	if request.SubscriberId == "" {
		return status.Error(codes.InvalidArgument, "subscriber_id is required")
	}
	return nil

}

func validateListSubscribersRequest(request *pb.ListSubscribersRequest) error {
	if request.PublisherId == "" {
		return status.Error(codes.InvalidArgument, "publisher_id is required")
	}
	return nil

}
