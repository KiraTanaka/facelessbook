package grpc

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"user_service/internal/services"

	pb "github.com/KiraTanaka/facelessbook_protos/gen/user"
)

type userServer struct {
	pb.UnimplementedUserServer
	userService services.UserService
}

func RegisterUserServer(gRPCServer *grpc.Server, userService services.UserService) {
	pb.RegisterUserServer(gRPCServer, &userServer{userService: userService})
}

func (s *userServer) NickName(ctx context.Context, request *pb.NickNameRequest) (*pb.NickNameResponse, error) {
	if err := validateNickNameRequest(request); err != nil {
		return nil, err
	}

	nick, err := s.userService.Nickname(request.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.NickNameResponse{Nickname: nick}, nil
}

func validateNickNameRequest(request *pb.NickNameRequest) error {
	if request.Id == "" {
		return status.Error(codes.InvalidArgument, "id is required")
	}
	return nil

}
