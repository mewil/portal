package api

import (
	"context"
	"io"
	"mime/multipart"

	"github.com/google/uuid"
	"github.com/mewil/portal/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *FrontendSvc) PostSvcGetFeed(ctx context.Context, newPostSvcClient PostSvcInjector, userID string, page uint32) ([]*pb.Post, uint32, error) {
	req := &pb.GetPostsRequest{UserID: userID, Page: page}
	res, err := newPostSvcClient().GetFeed(ctx, req)
	return res.GetPosts(), res.GetNextPage(), err
}

func (s *FrontendSvc) PostSvcGetPost(ctx context.Context, newPostSvcClient PostSvcInjector, postID string) (*pb.Post, error) {
	req := &pb.GetPostRequest{PostID: postID}
	return newPostSvcClient().GetPost(ctx, req)
}

func (s *FrontendSvc) PostSvcGetPostLikes(ctx context.Context, newPostSvcClient PostSvcInjector, postID string, page uint32) ([]*pb.PostLike, uint32, error) {
	req := &pb.GetPostLikesRequest{PostID: postID, Page: page}
	res, err := newPostSvcClient().GetPostLikes(ctx, req)
	return res.GetLikes(), res.GetNextPage(), err
}

func (s *FrontendSvc) PostSvcGetPostComments(ctx context.Context, newPostSvcClient PostSvcInjector, postID string, page uint32) ([]*pb.Comment, uint32, error) {
	req := &pb.GetPostCommentsRequest{PostID: postID, Page: page}
	res, err := newPostSvcClient().GetPostComments(ctx, req)
	return res.GetComments(), res.GetNextPage(), err
}

func (s *FrontendSvc) PostSvcCreatePost(ctx context.Context, newPostSvcClient PostSvcInjector, newFileSvcClient FileSvcInjector, userID, caption string, data multipart.File) (*pb.Post, error) {
	postID, err := uuid.NewUUID()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate UUID for new post %s", err.Error())
	}
	stream, err := newFileSvcClient().Upload(ctx)
	if err != nil {
		return nil, err
	}
	buf := make([]byte, fileBufferSize)
	for i := 0; ; i++ {
		if i*fileBufferSize > fileSizeLimit {
			return nil, status.Error(codes.InvalidArgument, "file too large")
		}
		n, err := data.Read(buf)
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, status.Error(codes.InvalidArgument, "failed to read file")
		}
		stream.Send(&pb.FilePart{
			Content: buf[:n],
		})
	}
	fileRes, err := stream.CloseAndRecv()
	if err != nil {
		return nil, err
	} else if fileRes.GetStatus() != pb.UploadStatusCode_Ok {
		return nil, status.Error(codes.Internal, "failed to upload file to file service")
	}
	req := &pb.CreatePostRequest{PostID: postID.String(), FileID: fileRes.GetFileID(), UserID: userID, Caption: caption}
	return newPostSvcClient().CreatePost(ctx, req)
}

func (s *FrontendSvc) PostSvcCreatePostLike(ctx context.Context, newPostSvcClient PostSvcInjector, postID, userID string) (*pb.Post, error) {
	likeID, err := uuid.NewUUID()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate UUID for new like %s", err.Error())
	}
	req := &pb.CreatePostLikeRequest{PostID: postID, LikeID: likeID.String(), UserID: userID}
	return newPostSvcClient().CreatePostLike(ctx, req)
}

func (s *FrontendSvc) PostSvcCreatePostComment(ctx context.Context, newPostSvcClient PostSvcInjector, postID, userID, text string) (*pb.Post, *pb.Comment, error) {
	commentID, err := uuid.NewUUID()
	if err != nil {
		return nil, nil, status.Errorf(codes.Internal, "failed to generate UUID for new comment %s", err.Error())
	}
	req := &pb.CreateCommentRequest{PostID: postID, CommentID: commentID.String(), UserID: userID, Text: text}
	res, err := newPostSvcClient().CreateComment(ctx, req)
	if err != nil {
		return nil, nil, err
	}
	return res.GetPost(), res.GetComment(), err
}

func (s *FrontendSvc) PostSvcGetComment(ctx context.Context, newPostSvcClient PostSvcInjector, commentID string) (*pb.Comment, error) {
	req := &pb.GetCommentRequest{CommentID: commentID}
	return newPostSvcClient().GetComment(ctx, req)
}

func (s *FrontendSvc) PostSvcGetCommentLikes(ctx context.Context, newPostSvcClient PostSvcInjector, commentID string) ([]*pb.CommentLike, uint32, error) {
	req := &pb.GetCommentLikesRequest{CommentID: commentID}
	res, err := newPostSvcClient().GetCommentLikes(ctx, req)
	return res.GetLikes(), res.GetNextPage(), err
}

func (s *FrontendSvc) PostSvcCreateCommentLike(ctx context.Context, newPostSvcClient PostSvcInjector, commentID, userID string) (*pb.Comment, error) {
	likeID, err := uuid.NewUUID()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate UUID for new like %s", err.Error())
	}
	req := &pb.CreateCommentLikeRequest{CommentID: commentID, LikeID: likeID.String(), UserID: userID}
	return newPostSvcClient().CreateCommentLike(ctx, req)
}
