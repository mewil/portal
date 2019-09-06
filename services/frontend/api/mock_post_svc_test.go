package api_test

import (
	"context"

	. "github.com/mewil/portal/frontend/api"
	"github.com/mewil/portal/pb"
	"google.golang.org/grpc"
)

type mockPostSvc struct {
}

func newMockPostSvc() *mockPostSvc {
	return &mockPostSvc{}
}

func (s *mockPostSvc) injectMockPostSvcClient() PostSvcInjector {
	return func() pb.PostSvcClient {
		return s.newMockPostSvcClient()
	}
}

func (s *mockPostSvc) newMockPostSvcClient() pb.PostSvcClient {
	return &mockPostSvcClient{
		svc: s,
	}
}

type mockPostSvcClient struct {
	svc *mockPostSvc
}

func (s *mockPostSvcClient) CreatePost(ctx context.Context, in *pb.CreatePostRequest, opts ...grpc.CallOption) (*pb.Post, error) {
	return nil, nil
}
func (s *mockPostSvcClient) CreatePostLike(ctx context.Context, in *pb.CreateLikeRequest, opts ...grpc.CallOption) (*pb.Post, error) {
	return nil, nil
}
func (s *mockPostSvcClient) CreateCommentLike(ctx context.Context, in *pb.CreateLikeRequest, opts ...grpc.CallOption) (*pb.Post, error) {
	return nil, nil
}
func (s *mockPostSvcClient) CreateComment(ctx context.Context, in *pb.CreateCommentRequest, opts ...grpc.CallOption) (*pb.Post, error) {
	return nil, nil
}
func (s *mockPostSvcClient) GetPost(ctx context.Context, in *pb.GetPostRequest, opts ...grpc.CallOption) (*pb.Post, error) {
	return nil, nil
}
func (s *mockPostSvcClient) GetFeed(ctx context.Context, in *pb.GetFeedRequest, opts ...grpc.CallOption) (*pb.GetFeedResponse, error) {
	return nil, nil
}
func (s *mockPostSvcClient) GetPostLikes(ctx context.Context, in *pb.PostLikesRequest, opts ...grpc.CallOption) (*pb.LikesResponse, error) {
	return nil, nil
}
func (s *mockPostSvcClient) GetCommentLikes(ctx context.Context, in *pb.CommentLikesRequest, opts ...grpc.CallOption) (*pb.LikesResponse, error) {
	return nil, nil
}
func (s *mockPostSvcClient) GetComments(ctx context.Context, in *pb.CommentsRequest, opts ...grpc.CallOption) (*pb.CommentsResponse, error) {
	return nil, nil
}
