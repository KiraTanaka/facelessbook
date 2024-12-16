package server_grpc

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"notification_service/internal/services"

	assessment "github.com/KiraTanaka/facelessbook_protos/gen/assessment"
)

type serverAPI struct {
	assessment.UnimplementedPostLikesServer
	notificationService services.NotificationService
}

func (s *serverAPI) GetLikeCount(
	ctx context.Context,
	in *assessment.GetLikeCountRequest,
) (*assessment.GetLikeCountResponse, error) {
	if in.PostId == "" {
		return nil, status.Error(codes.InvalidArgument, "postId is required")
	}

	postLikeCount, err := s.notificationService.GetLikeCount(in.GetPostId())
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get like count")

	}

	return &assessment.GetLikeCountResponse{Cnt: int64(postLikeCount)}, nil
}

func Register(gRPCServer *grpc.Server, notificationService services.PostLikeService) {
	assessment.RegisterPostLikesServer(gRPCServer, &serverAPI{notificationService: notificationService})
}
