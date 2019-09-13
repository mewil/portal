package main

import (
	"context"

	"github.com/mewil/portal/common/validation"
	"github.com/mewil/portal/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type postSvc struct {
	repository PostRepository
}

func (s *postSvc) CreatePost(ctx context.Context, in *pb.CreatePostRequest) (*pb.Post, error) {
	postID := in.GetPostID()
	if err := validation.ValidUUID(postID); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid post id %s", err.Error())
	}
	userID := in.GetUserID()
	if err := validation.ValidUUID(userID); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user id %s", err.Error())
	}
	fileID := in.GetFileID()
	if err := validation.ValidUUID(fileID); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid file id %s", err.Error())
	}
	if err := s.repository.CreatePost(postID, userID, fileID, in.GetCaption()); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create post %s", err.Error())
	}
	return s.repository.GetPost(postID)
}

func (s *postSvc) CreatePostLike(ctx context.Context, in *pb.CreatePostLikeRequest) (*pb.Post, error) {
	postID := in.GetPostID()
	if err := validation.ValidUUID(postID); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid post id %s", err.Error())
	}
	userID := in.GetUserID()
	if err := validation.ValidUUID(userID); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user id %s", err.Error())
	}
	likeID := in.GetLikeID()
	if err := validation.ValidUUID(likeID); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid like id %s", err.Error())
	}
	if err := s.repository.CreatePostLike(postID, userID, likeID); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create post like %s", err.Error())
	}
	return s.repository.GetPost(postID)
}

func (s *postSvc) CreateComment(ctx context.Context, in *pb.CreateCommentRequest) (*pb.CreateCommentResponse, error) {
	postID := in.GetPostID()
	if err := validation.ValidUUID(postID); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid post id %s", err.Error())
	}
	userID := in.GetUserID()
	if err := validation.ValidUUID(userID); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user id %s", err.Error())
	}
	commentID := in.GetCommentID()
	if err := validation.ValidUUID(commentID); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid like id %s", err.Error())
	}
	if err := s.repository.CreateComment(postID, userID, commentID, in.GetText()); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create comment %s", err.Error())
	}
	post, err := s.repository.GetPost(postID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to fetch post %s", err.Error())
	}
	comment, err := s.repository.GetComment(commentID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to fetch comment %s", err.Error())
	}
	return &pb.CreateCommentResponse{
		Post:    post,
		Comment: comment,
	}, nil
}

func (s *postSvc) CreateCommentLike(ctx context.Context, in *pb.CreateCommentLikeRequest) (*pb.Comment, error) {
	commentID := in.GetCommentID()
	if err := validation.ValidUUID(commentID); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid comment id %s", err.Error())
	}
	userID := in.GetUserID()
	if err := validation.ValidUUID(userID); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user id %s", err.Error())
	}
	likeID := in.GetLikeID()
	if err := validation.ValidUUID(likeID); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid like id %s", err.Error())
	}
	if err := s.repository.CreateCommentLike(commentID, userID, likeID); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create comment like %s", err.Error())
	}
	return s.repository.GetComment(commentID)
}

func (s *postSvc) DeletePost(ctx context.Context, in *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	postID := in.GetDeleteID()
	if err := validation.ValidUUID(postID); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid post id %s", err.Error())
	}
	err := s.repository.DeletePost(postID)
	return &pb.DeleteResponse{}, err
}

func (s *postSvc) DeletePostLike(ctx context.Context, in *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	likeID := in.GetDeleteID()
	if err := validation.ValidUUID(likeID); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid post like id %s", err.Error())
	}
	err := s.repository.DeletePostLike(likeID)
	return &pb.DeleteResponse{}, err
}

func (s *postSvc) DeleteComment(ctx context.Context, in *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	commentID := in.GetDeleteID()
	if err := validation.ValidUUID(commentID); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid comment id %s", err.Error())
	}
	err := s.repository.DeleteComment(commentID)
	return &pb.DeleteResponse{}, err
}

func (s *postSvc) DeleteCommentLike(ctx context.Context, in *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	likeID := in.GetDeleteID()
	if err := validation.ValidUUID(likeID); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid comment like id %s", err.Error())
	}
	err := s.repository.DeleteCommentLike(likeID)
	return &pb.DeleteResponse{}, err
}

func (s *postSvc) GetPost(ctx context.Context, in *pb.GetPostRequest) (*pb.Post, error) {
	postID := in.GetPostID()
	if err := validation.ValidUUID(postID); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid post id %s", err.Error())
	}
	return s.repository.GetPost(postID)
}

func (s *postSvc) GetProfile(ctx context.Context, in *pb.GetPostsRequest) (*pb.GetPostsResponse, error) {
	userID := in.GetUserID()
	if err := validation.ValidUUID(userID); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user id %s", err.Error())
	}
	posts, err := s.repository.GetProfile(userID, in.GetPage())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to fetch profile %s", err.Error())
	}
	return &pb.GetPostsResponse{
		Posts:    posts,
		NextPage: in.GetPage() + 1,
	}, nil
}

func (s *postSvc) GetFeed(ctx context.Context, in *pb.GetPostsRequest) (*pb.GetPostsResponse, error) {
	userID := in.GetUserID()
	if err := validation.ValidUUID(userID); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user id %s", err.Error())
	}
	posts, err := s.repository.GetFeed(userID, in.GetPage())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to fetch feed %s", err.Error())
	}
	return &pb.GetPostsResponse{
		Posts:    posts,
		NextPage: in.GetPage() + 1,
	}, nil
}

func (s *postSvc) GetPostLikes(ctx context.Context, in *pb.GetPostLikesRequest) (*pb.PostLikesResponse, error) {
	postID := in.GetPostID()
	if err := validation.ValidUUID(postID); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid post id %s", err.Error())
	}
	likes, err := s.repository.GetPostLikes(postID, in.GetPage())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to fetch post likes %s", err.Error())
	}
	return &pb.PostLikesResponse{
		Likes:    likes,
		NextPage: in.GetPage() + 1,
	}, nil
}

func (s *postSvc) GetPostComments(ctx context.Context, in *pb.GetPostCommentsRequest) (*pb.CommentsResponse, error) {
	postID := in.GetPostID()
	if err := validation.ValidUUID(postID); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid post id %s", err.Error())
	}
	comments, err := s.repository.GetPostComments(postID, in.GetPage())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to fetch post comments %s", err.Error())
	}
	return &pb.CommentsResponse{
		Comments: comments,
		NextPage: in.GetPage() + 1,
	}, nil
}

func (s *postSvc) GetCommentLikes(ctx context.Context, in *pb.GetCommentLikesRequest) (*pb.CommentLikesResponse, error) {
	commentID := in.GetCommentID()
	if err := validation.ValidUUID(commentID); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid comment id %s", err.Error())
	}
	likes, err := s.repository.GetCommentLikes(commentID, in.GetPage())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to fetch comment likes %s", err.Error())
	}
	return &pb.CommentLikesResponse{
		Likes:    likes,
		NextPage: in.GetPage() + 1,
	}, nil
}

func (s *postSvc) GetComment(ctx context.Context, in *pb.GetCommentRequest) (*pb.Comment, error) {
	commentID := in.GetCommentID()
	if err := validation.ValidUUID(commentID); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid comment id %s", err.Error())
	}
	return s.repository.GetComment(commentID)
}
