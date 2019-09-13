package main

import (
	"context"

	"github.com/mewil/portal/common/validation"
	"github.com/mewil/portal/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type userSvc struct {
	repository UserRepository
}

func (s *userSvc) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.User, error) {
	userID := in.GetUserID()
	if err := s.validateUserIDDoesNotExist(userID); err != nil {
		return nil, err
	}
	username := in.GetUsername()
	if err := s.validateUsernameDoesNotExist(userID); err != nil {
		return nil, err
	}
	email := in.GetEmail()
	if !validation.ValidEmail(email) {
		return nil, status.Error(codes.InvalidArgument, "invalid email format")
	}
	if err := s.repository.CreateUser(userID, username, email); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create user %s", err.Error())
	}
	user, err := s.repository.GetUser(userID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to fetch user %s", err.Error())
	}
	return user, nil
}

func (s *userSvc) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.User, error) {
	userID := in.GetUserID()
	if err := s.validateUserIDExists(userID); err != nil {
		return nil, err
	}
	user, err := s.repository.GetUser(userID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to fetch user %s", err.Error())
	}
	return user, nil
}

func (s *userSvc) GetFollowers(ctx context.Context, in *pb.GetFollowersRequest) (*pb.GetFollowersResponse, error) {
	userID := in.GetUserID()
	if err := s.validateUserIDExists(userID); err != nil {
		return nil, err
	}
	followers, err := s.repository.GetFollowers(userID, in.GetPage())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to fetch followers %s", err.Error())
	}
	return &pb.GetFollowersResponse{
		Followers: followers,
	}, nil
}

func (s *userSvc) GetFollowing(ctx context.Context, in *pb.GetFollowingRequest) (*pb.GetFollowingResponse, error) {
	userID := in.GetUserID()
	if err := s.validateUserIDExists(userID); err != nil {
		return nil, err
	}
	following, err := s.repository.GetFollowing(userID, in.GetPage())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to fetch following %s", err.Error())
	}
	return &pb.GetFollowingResponse{
		Following: following,
	}, nil
}

func (s *userSvc) CreateFollow(ctx context.Context, in *pb.FollowRequest) (*pb.FollowResponse, error) {
	userID := in.GetUserID()
	if err := s.validateUserIDExists(userID); err != nil {
		return nil, err
	}
	followID := in.GetFollowID()
	if err := s.validateUserIDExists(followID); err != nil {
		return nil, err
	}
	user, err := s.repository.GetUser(userID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to fetch user %s", err.Error())
	}
	followed, err := s.repository.GetUser(followID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to fetch user %s", err.Error())
	}
	return &pb.FollowResponse{
		User:       user,
		FollowUser: followed,
	}, nil
}

func (s *userSvc) RemoveFollow(ctx context.Context, in *pb.FollowRequest) (*pb.FollowResponse, error) {
	userID := in.GetUserID()
	if err := s.validateUserIDExists(userID); err != nil {
		return nil, err
	}
	followID := in.GetFollowID()
	if err := s.validateUserIDExists(followID); err != nil {
		return nil, err
	}
	user, err := s.repository.GetUser(userID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to fetch user %s", err.Error())
	}
	unfollowed, err := s.repository.GetUser(followID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to fetch user %s", err.Error())
	}
	return &pb.FollowResponse{
		User:       user,
		FollowUser: unfollowed,
	}, nil
}

func (s *userSvc) validateUserIDExists(userID string) error {
	if err := validation.ValidUUID(userID); err != nil {
		return status.Errorf(codes.InvalidArgument, "invalid user id %s", err.Error())
	}
	if exists, err := s.repository.UserIDExists(userID); err != nil {
		return status.Errorf(codes.Internal, "failed to fetch if user id exists %s", err.Error())
	} else if !exists {
		return status.Error(codes.InvalidArgument, "user id does not exist")
	}
	return nil
}

func (s *userSvc) validateUserIDDoesNotExist(userID string) error {
	if err := validation.ValidUUID(userID); err != nil {
		return status.Errorf(codes.InvalidArgument, "invalid user id %s", err.Error())
	}
	if exists, err := s.repository.UserIDExists(userID); err != nil {
		return status.Errorf(codes.Internal, "failed to fetch if user id exists %s", err.Error())
	} else if exists {
		return status.Error(codes.InvalidArgument, "user id already exists")
	}
	return nil
}

func (s *userSvc) validateUsernameDoesNotExist(username string) error {
	if !validation.ValidUsername(username) {
		return status.Error(codes.InvalidArgument, "invalid username format")
	}
	if exists, err := s.repository.UsernameExists(username); err != nil {
		return status.Errorf(codes.Internal, "failed to fetch if username exists %s", err.Error())
	} else if exists {
		return status.Error(codes.InvalidArgument, "username already exists")
	}
	return nil
}
