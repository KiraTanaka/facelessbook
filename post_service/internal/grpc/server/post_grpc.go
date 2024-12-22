package grpc

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"post_service/internal/models"
	"post_service/internal/services"

	pb "github.com/KiraTanaka/facelessbook_protos/gen/post"
)

type postServer struct {
	pb.UnimplementedPostServer
	postService services.PostService
}

func RegisterPostServer(gRPCServer *grpc.Server, postService services.PostService) {
	pb.RegisterPostServer(gRPCServer, &postServer{postService: postService})
}

func (s *postServer) Create(ctx context.Context, request *pb.CreateRequest) (*pb.CreateResponse, error) {
	post := &models.Post{
		AuthorId: request.Post.AuthorId,
		Text:     request.Post.Text,
	}
	postId, err := s.postService.Create(post)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.CreateResponse{Id: postId}, nil
}

func (s *postServer) PostById(ctx context.Context, request *pb.PostByIdRequest) (*pb.PostByIdResponse, error) {
	if err := validatePostByIdRequest(request); err != nil {
		return nil, err
	}
	post, err := s.postService.PostById(request.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())

	}
	postMessage := PostModelToPostMessage(post)

	return &pb.PostByIdResponse{Post: postMessage}, nil
}

func (s *postServer) ListPosts(ctx context.Context, request *pb.ListPostsRequest) (*pb.ListPostsResponse, error) {
	posts, err := s.postService.ListPosts()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	postMessages := make([]*pb.PostInformation, len(posts))

	for i, post := range posts {
		postMessages[i] = PostModelToPostMessage(post)
	}

	return &pb.ListPostsResponse{Posts: postMessages}, nil
}

func (s *postServer) Update(ctx context.Context, request *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	if err := validateUpdateRequest(request); err != nil {
		return nil, err
	}
	err := s.postService.Update(request.Id, request.NewText)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.UpdateResponse{Success: true}, nil
}

func (s *postServer) Delete(ctx context.Context, request *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	if err := validateDeleteRequest(request); err != nil {
		return nil, err
	}
	err := s.postService.Delete(request.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.DeleteResponse{Success: true}, nil
}

func PostModelToPostMessage(post *models.Post) *pb.PostInformation {
	return &pb.PostInformation{
		Id:          post.Id,
		CreatedTime: post.CreatedTime.Format(time.DateTime),
		AuthorId:    post.AuthorId,
		Text:        post.Text}

}

func validatePostByIdRequest(request *pb.PostByIdRequest) error {
	if request.Id == "" {
		return status.Error(codes.InvalidArgument, "id is required")
	}
	return nil
}

func validateUpdateRequest(request *pb.UpdateRequest) error {
	if request.Id == "" {
		return status.Error(codes.InvalidArgument, "id is required")
	}

	if request.NewText == "" {
		return status.Error(codes.InvalidArgument, "new text is required")
	}
	return nil
}

func validateDeleteRequest(request *pb.DeleteRequest) error {
	if request.Id == "" {
		return status.Error(codes.InvalidArgument, "id is required")
	}
	return nil
}
