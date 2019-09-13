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
		token, err = s.GenerateAdminAuthToken(res.GetUserID())
	} else {
		token, err = s.GenerateUserAuthToken(res.GetUserID())
	}
	if err != nil {
		return "", "", status.Errorf(codes.Internal, "failed to create token for user %s %s", res.GetUserID(), err.Error())
	}
	return token, res.GetUserID(), nil
}

func (s *FrontendSvc) AuthSvcSignUp(ctx context.Context, newAuthSvcClient AuthSvcInjector, newUserSvcClient UserSvcInjector, username, name, email, password string) (*pb.User, string, error) {
	userID, err := uuid.NewUUID()
	if err != nil {
		return nil, "", status.Errorf(codes.Internal, "failed to generate UUID for new user %s", err.Error())
	}
	authReq := &pb.SignUpRequest{Email: email, UserID: userID.String(), Password: password}
	authRes, err := newAuthSvcClient().SignUp(ctx, authReq)
	if err != nil {
		return nil, "", err
	}
	token, err := "", *new(error)
	if authRes.GetIsAdmin() {
		token, err = s.GenerateAdminAuthToken(authRes.GetUserID())
	} else {
		token, err = s.GenerateUserAuthToken(authRes.GetUserID())
	}
	userReq := &pb.CreateUserRequest{UserID: userID.String(), Username: username, Email: email, Name: name}
	user, err := newUserSvcClient().CreateUser(ctx, userReq)
	if err != nil {
		return nil, "", err
	}
	return user, token, nil
}
