package api

import (
	"context"

	"github.com/mewil/portal/pb"
)

func (s *FrontendSvc) UserGetUser(ctx context.Context, newUserSvcClient UserSvcInjector, userId string) (*pb.User, error) {
	req := &pb.GetUserRequest{UserId: userId}
	return newUserSvcClient().GetUser(ctx, req)
}

func (s *FrontendSvc) UserGetFollowers(ctx context.Context, newUserSvcClient UserSvcInjector, userId string, page uint32) ([]*pb.User, error) {
	req := &pb.GetFollowersRequest{UserId: userId, Page: page}
	res, err := newUserSvcClient().GetFollowers(ctx, req)
	return res.GetFollowers(), err
}

func (s *FrontendSvc) UserGetFollowing(ctx context.Context, newUserSvcClient UserSvcInjector, userId string, page uint32) ([]*pb.User, error) {
	req := &pb.GetFollowingRequest{UserId: userId, Page: page}
	res, err := newUserSvcClient().GetFollowing(ctx, req)
	return res.GetFollowing(), err
}
