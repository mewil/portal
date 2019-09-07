package api

import (
	"context"

	"github.com/google/uuid"
	"github.com/mewil/portal/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *FrontendSvc) AuthSvcSignIn(ctx context.Context, newAuthSvcClient func() pb.AuthSvcClient, email, password string) (string, string, error) {
	req := &pb.SignInRequest{Email: email, Password: password}
	res, err := newAuthSvcClient().SignIn(ctx, req)
	if err != nil {
		return "", "", err
	}
	token, err := "", *new(error)
	if res.GetIsAdmin() {
		token, err = s.GenerateAdminAuthToken(res.GetUserId())
	} else {
		token, err = s.GenerateUserAuthToken(res.GetUserId())
	}
	if err != nil {
		return "", "", status.Errorf(codes.Internal, "failed to create token for user %s %s", res.GetUserId(), err.Error())
	}
	return token, res.GetUserId(), nil
}

func (s *FrontendSvc) AuthSvcSignUp(ctx context.Context, newAuthSvcClient AuthSvcInjector, newUserSvcClient UserSvcInjector, username, name, email, password string) (*pb.User, string, error) {
	userId, err := uuid.NewUUID()
	if err != nil {
		return nil, "", status.Errorf(codes.Internal, "failed to generate UUID for new user %s", err.Error())
	}
	authReq := &pb.SignUpRequest{Email: email, UserId: userId.String(), Password: password}
	authRes, err := newAuthSvcClient().SignUp(ctx, authReq)
	if err != nil {
		return nil, "", err
	}
	token, err := "", *new(error)
	if authRes.GetIsAdmin() {
		token, err = s.GenerateAdminAuthToken(authRes.GetUserId())
	} else {
		token, err = s.GenerateUserAuthToken(authRes.GetUserId())
	}
	userReq := &pb.CreateUserRequest{UserId: userId.String(), Username: username, Email: email, Name: name}
	user, err := newUserSvcClient().CreateUser(ctx, userReq)
	if err != nil {
		return nil, "", err
	}
	return user, token, nil
}
