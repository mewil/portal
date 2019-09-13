package api

import (
	"context"

	"github.com/mewil/portal/pb"
)

func (s *FrontendSvc) UserSvcGetUser(ctx context.Context, newUserSvcClient UserSvcInjector, userID string) (*pb.User, error) {
	req := &pb.GetUserRequest{UserID: userID}
	return newUserSvcClient().GetUser(ctx, req)
}

func (s *FrontendSvc) UserSvcGetFollowers(ctx context.Context, newUserSvcClient UserSvcInjector, userID string, page uint32) ([]*pb.User, error) {
	req := &pb.GetFollowersRequest{UserID: userID, Page: page}
	res, err := newUserSvcClient().GetFollowers(ctx, req)
	return res.GetFollowers(), err
}

func (s *FrontendSvc) UserSvcGetFollowing(ctx context.Context, newUserSvcClient UserSvcInjector, userID string, page uint32) ([]*pb.User, error) {
	req := &pb.GetFollowingRequest{UserID: userID, Page: page}
	res, err := newUserSvcClient().GetFollowing(ctx, req)
	return res.GetFollowing(), err
}

func (s *FrontendSvc) UserSvcGetProfile(ctx context.Context, newUserSvcClient UserSvcInjector, newPostSvcClient PostSvcInjector, userID string, page uint32) (*pb.User, []*pb.Post, uint32, error) {
	userReq := &pb.GetUserRequest{UserID: userID}
	userRes, err := newUserSvcClient().GetUser(ctx, userReq)
	if err != nil {
		return nil, nil, 0, err
	}
	postReq := &pb.GetPostsRequest{UserID: userID, Page: page}
	postRes, err := newPostSvcClient().GetProfile(ctx, postReq)
	return userRes, postRes.GetPosts(), postRes.GetNextPage(), nil
}

func (s *FrontendSvc) UserSvcCreateFollow(ctx context.Context, newUserSvcClient UserSvcInjector, userID, followID string) (*pb.User, *pb.User, error) {
	req := &pb.FollowRequest{UserID: userID, FollowID: followID}
	res, err := newUserSvcClient().CreateFollow(ctx, req)
	return res.GetUser(), res.GetFollowingUser(), err
}

func (s *FrontendSvc) UserSvcRemoveFollow(ctx context.Context, newUserSvcClient UserSvcInjector, userID, followID string) (*pb.User, *pb.User, error) {
	req := &pb.FollowRequest{UserID: userID, FollowID: followID}
	res, err := newUserSvcClient().RemoveFollow(ctx, req)
	return res.GetUser(), res.GetFollowingUser(), err
}
