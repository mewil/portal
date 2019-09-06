package api

import (
	"context"

	"github.com/mewil/portal/pb"
	"google.golang.org/grpc"
)

type mockUserSvcClient struct {
}

func newMockUserSvcClient() pb.UserSvcClient {
	return &mockUserSvcClient{}
}

func (s *mockUserSvcClient) CreateUser(ctx context.Context, in *pb.CreateUserRequest, opts ...grpc.CallOption) (*pb.User, error) {
	return nil, nil
}
func (s *mockUserSvcClient) GetUser(ctx context.Context, in *pb.GetUserRequest, opts ...grpc.CallOption) (*pb.User, error) {
	return nil, nil
}
func (s *mockUserSvcClient) GetFollowers(ctx context.Context, in *pb.FollowersRequest, opts ...grpc.CallOption) (*pb.FollowersResponse, error) {
	return nil, nil
}
func (s *mockUserSvcClient) GetFollowing(ctx context.Context, in *pb.FollowingRequest, opts ...grpc.CallOption) (*pb.FollowingResponse, error) {
	return nil, nil
}

func injectMockUserSvcClient() UserSvcInjector {
	return func() pb.UserSvcClient {
		return newMockUserSvcClient()
	}
}
