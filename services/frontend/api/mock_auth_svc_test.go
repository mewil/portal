package api_test

import (
	"context"

	. "github.com/mewil/portal/frontend/api"
	"github.com/mewil/portal/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type mockAuthSvc struct {
	store map[string]struct {
		UserID   string
		Password string
	}
}

func newMockAuthSvc() *mockAuthSvc {
	return &mockAuthSvc{
		store: make(map[string]struct {
			UserID   string
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
		UserID:  s.svc.store[in.Email].UserID,
		IsAdmin: false,
	}, nil
}

func (s *mockAuthSvcClient) SignUp(ctx context.Context, in *pb.SignUpRequest, opts ...grpc.CallOption) (*pb.SignInResponse, error) {
	if in.Email != "correct@email.format" || in.Password != "valid_password" {
		return nil, status.Error(codes.InvalidArgument, "invalid email or password")
	}

	s.svc.store[in.Email] = struct {
		UserID   string
		Password string
	}{
		UserID:   in.UserID,
		Password: in.Password,
	}
	return &pb.SignInResponse{
		UserID:  in.UserID,
		IsAdmin: false,
	}, nil
}
