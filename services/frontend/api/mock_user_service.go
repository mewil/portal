package api

import (
	"context"

	"github.com/mewil/portal/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type mockUserSvc struct {
	userStore map[string]pb.User
}

func newMockUserSvc() *mockUserSvc {
	return &mockUserSvc{
		userStore: make(map[string]pb.User, 0),
	}
}

func (s *mockUserSvc) injectMockUserSvcClient() UserSvcInjector {
	return func() pb.UserSvcClient {
		return s.newMockUserSvcClient()
	}
}

func (s *mockUserSvc) newMockUserSvcClient() pb.UserSvcClient {
	return &mockUserSvcClient{
		svc: s,
	}
}

type mockUserSvcClient struct {
	svc *mockUserSvc
}

func newMockUserSvcClient() pb.UserSvcClient {
	return &mockUserSvcClient{}
}

func (s *mockUserSvcClient) CreateUser(ctx context.Context, in *pb.CreateUserRequest, opts ...grpc.CallOption) (*pb.User, error) {
	if in.Name == "database error" {
		return nil, status.Error(codes.Internal, "database error")
	}
	user := pb.User{
		UserId:   in.UserId,
		Username: in.Username,
		Email:    in.Email,
		Name:     in.Name,
	}
	s.svc.userStore[in.UserId] = user
	return &user, nil
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
