package api

import (
	"context"

	"github.com/mewil/portal/pb"
	"google.golang.org/grpc"
)

type mockAuthSvcClient struct {
}

func newMockAuthSvcClient() pb.AuthSvcClient {
	return &mockAuthSvcClient{}
}

func (s *mockAuthSvcClient) SignIn(ctx context.Context, in *pb.SignInRequest, opts ...grpc.CallOption) (*pb.SignInResponse, error) {
	return nil, nil
}

func (s *mockAuthSvcClient) SignUp(ctx context.Context, in *pb.SignUpRequest, opts ...grpc.CallOption) (*pb.SignInResponse, error) {
	return nil, nil
}

func injectMockAuthSvcClient() AuthSvcInjector {
	return func() pb.AuthSvcClient {
		return newMockAuthSvcClient()
	}
}
