package main

import (
	"context"

	"github.com/mewil/portal/common/validation"
	"github.com/mewil/portal/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type authSvc struct {
	repository AuthRepository
}

func (s *authSvc) SignIn(ctx context.Context, in *pb.SignInRequest) (*pb.SignInResponse, error) {
	email := in.GetEmail()
	if !validation.ValidEmail(email) {
		return nil, status.Error(codes.InvalidArgument, "invalid email format")
	}
	exists, err := s.repository.EmailExists(email)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to fetch if email exists %s", err.Error())
	}
	if !exists {
		return nil, status.Error(codes.InvalidArgument, "email does not exist")
	}
	passwordHash, err := s.repository.LoadPasswordHash(email)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to fetch password hash exists %s", err.Error())
	}
	if !ValidPassword(in.GetPassword(), passwordHash) {
		return nil, status.Error(codes.PermissionDenied, "invalid password")
	}
	userID, isAdmin, err := s.repository.LoadUserIDAndAdminStatus(email)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to load user id and admin status %s", err.Error())
	}
	return &pb.SignInResponse{
		UserID:  userID,
		IsAdmin: isAdmin,
	}, nil
}

func (s *authSvc) SignUp(ctx context.Context, in *pb.SignUpRequest) (*pb.SignInResponse, error) {
	email := in.GetEmail()
	if !validation.ValidEmail(email) {
		return nil, status.Error(codes.InvalidArgument, "invalid email format")
	}
	userID := in.GetUserID()
	if err := validation.ValidUUID(userID); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user id %s", err.Error())
	}
	exists, err := s.repository.EmailExists(email)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to fetch if email exists %s", err.Error())
	}
	if exists {
		return nil, status.Error(codes.InvalidArgument, "email already exists")
	}
	password := in.GetPassword()
	if len(password) >= 8 {
		return nil, status.Error(codes.InvalidArgument, "password is less than 8 characters")
	}
	if err := s.repository.StoreAuthRecord(email, userID, password, false); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to insert auth record into database %s", err.Error())
	}
	userID, isAdmin, err := s.repository.LoadUserIDAndAdminStatus(email)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to load user id and admin status %s", err.Error())
	}
	return &pb.SignInResponse{
		UserID:  userID,
		IsAdmin: isAdmin,
	}, nil
	return nil, nil
}
