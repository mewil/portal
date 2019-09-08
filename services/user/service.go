package main

import (
	"context"
	"fmt"
	"net"
	"os"

	"github.com/mewil/portal/common/grpc_utils"
	"github.com/mewil/portal/common/logger"
	"github.com/mewil/portal/pb"
)

func main() {
	log, err := logger.NewLogger()
	if err != nil {
		panic(err)
	}
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("PORT")))
	if err != nil {
		log.Fatal(err)
	}

	s, err := grpc_utils.NewServer(log)
	if err != nil {
		log.Fatal(err)
	}
	pb.RegisterUserSvcServer(s, &userSvc{})
	s.Serve(listener)
}

type userSvc struct {
}

func (s *userSvc) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.User, error) {
	return nil, nil
}

func (s *userSvc) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.User, error) {
	return nil, nil
}

func (s *userSvc) GetFollowers(ctx context.Context, in *pb.GetFollowersRequest) (*pb.GetFollowersResponse, error) {
	return nil, nil
}

func (s *userSvc) GetFollowing(ctx context.Context, in *pb.GetFollowingRequest) (*pb.GetFollowingResponse, error) {
	return nil, nil
}

func (s *userSvc) CreateFollow(ctx context.Context, in *pb.FollowRequest) (*pb.FollowResponse, error) {
	return nil, nil
}

func (s *userSvc) RemoveFollow(ctx context.Context, in *pb.FollowRequest) (*pb.FollowResponse, error) {
	return nil, nil
}
