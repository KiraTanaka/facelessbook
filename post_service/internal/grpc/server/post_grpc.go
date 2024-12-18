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
	postMessage := request.Post
	post := &models.Post{
		AuthorId: postMessage.AuthorId,
		Text:     postMessage.Text,
	}
	postId, err := s.postService.Create(post)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed create")
	}

	return &pb.CreateResponse{Id: postId}, nil
}

func (s *postServer) PostById(ctx context.Context, request *pb.PostByIdRequest) (*pb.PostByIdResponse, error) {
	post, err := s.postService.PostById(request.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed find post by id")

	}
	postMessage, err := PostModelToPostMessage(post)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed find post by id")

	}

	return &pb.PostByIdResponse{Post: postMessage}, nil
}

func (s *postServer) ListPosts(ctx context.Context, request *pb.ListPostsRequest) (*pb.ListPostsResponse, error) {
	posts, err := s.postService.ListPosts()
	if err != nil {
		return nil, status.Error(codes.Internal, "failed find list posts")
	}

	postMessages := []*pb.PostMessage{}

	for _, post := range posts {
		postMessage, _ := PostModelToPostMessage(post)
		postMessages = append(postMessages, postMessage)
	}

	return &pb.ListPostsResponse{Posts: postMessages}, nil
}

func PostModelToPostMessage(post *models.Post) (*pb.PostMessage, error) {
	return &pb.PostMessage{
		Id:          post.Id,
		CreatedTime: post.CreatedTime.Format(time.DateTime),
		AuthorId:    post.AuthorId,
		Text:        post.Text}, nil

}
