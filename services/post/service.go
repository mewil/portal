package main

import (
	"context"
	"fmt"
	"net"
	"os"

	"github.com/mewil/portal/common/grpc_utils"
	"github.com/mewil/portal/common/logger"
	"github.com/mewil/portal/pb"
)

func main() {
	log, err := logger.NewLogger()
	if err != nil {
		panic(err)
	}
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("PORT")))
	if err != nil {
		log.Fatal(err)
	}

	s, err := grpc_utils.NewServer(log)
	if err != nil {
		log.Fatal(err)
	}
	pb.RegisterPostSvcServer(s, &postSvc{})
	s.Serve(listener)
}

type postSvc struct {
}

func (s *postSvc) CreatePost(ctx context.Context, in *pb.CreatePostRequest) (*pb.Post, error) {
	return nil, nil
}

func (s *postSvc) CreatePostLike(ctx context.Context, in *pb.CreatePostLikeRequest) (*pb.Post, error) {
	return nil, nil
}

func (s *postSvc) CreateComment(ctx context.Context, in *pb.CreateCommentRequest) (*pb.CreateCommentResponse, error) {
	return nil, nil
}

func (s *postSvc) CreateCommentLike(ctx context.Context, in *pb.CreateCommentLikeRequest) (*pb.Post, error) {
	return nil, nil
}

func (s *postSvc) GetPost(ctx context.Context, in *pb.GetPostRequest) (*pb.Post, error) {
	return nil, nil
}

func (s *postSvc) GetProfile(ctx context.Context, in *pb.GetPostsRequest) (*pb.GetPostsResponse, error) {
	return nil, nil
}

func (s *postSvc) GetFeed(ctx context.Context, in *pb.GetPostsRequest) (*pb.GetPostsResponse, error) {
	return nil, nil
}

func (s *postSvc) GetPostLikes(ctx context.Context, in *pb.GetPostLikesRequest) (*pb.LikesResponse, error) {
	return nil, nil
}

func (s *postSvc) GetPostComments(ctx context.Context, in *pb.GetPostCommentsRequest) (*pb.CommentsResponse, error) {
	return nil, nil
}

func (s *postSvc) GetCommentLikes(ctx context.Context, in *pb.GetCommentLikesRequest) (*pb.LikesResponse, error) {
	return nil, nil
}

func (s *postSvc) GetComment(ctx context.Context, in *pb.GetCommentRequest) (*pb.Comment, error) {
	return nil, nil
}
