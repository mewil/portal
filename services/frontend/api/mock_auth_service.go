package api

import (
	"context"

	"github.com/mewil/portal/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type mockAuthSvc struct {
	store map[string]struct {
		UserId   string
		Password string
	}
}

func newMockAuthSvc() *mockAuthSvc {
	return &mockAuthSvc{
		store: make(map[string]struct {
			UserId   string
			Password string
		}, 0),
	}
}

func (s *mockAuthSvc) injectMockAuthSvcClient() AuthSvcInjector {
	return func() pb.AuthSvcClient {
		return s.newMockAuthSvcClient()
	}
}

func (s *mockAuthSvc) newMockAuthSvcClient() pb.AuthSvcClient {
	return &mockAuthSvcClient{
		svc: s,
	}
}

type mockAuthSvcClient struct {
	svc *mockAuthSvc
}

func (s *mockAuthSvcClient) SignIn(ctx context.Context, in *pb.SignInRequest, opts ...grpc.CallOption) (*pb.SignInResponse, error) {
	if in.Password != s.svc.store[in.Email].Password {
		return nil, status.Error(codes.Unauthenticated, "invalid password")
	}
	return &pb.SignInResponse{
		UserId:  s.svc.store[in.Email].UserId,
		IsAdmin: false,
	}, nil
}

func (s *mockAuthSvcClient) SignUp(ctx context.Context, in *pb.SignUpRequest, opts ...grpc.CallOption) (*pb.SignInResponse, error) {
	if in.Email != "correct@email.format" || in.Password != "valid_password" {
		return nil, status.Error(codes.InvalidArgument, "invalid email or password")
	}

	s.svc.store[in.Email] = struct {
		UserId   string
		Password string
	}{
		UserId:   in.UserId,
		Password: in.Password,
	}
	return &pb.SignInResponse{
		UserId:  in.UserId,
		IsAdmin: false,
	}, nil
}
