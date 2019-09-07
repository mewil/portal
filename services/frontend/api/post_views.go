package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetFeed handles GET /v1/post/
func (s *FrontendSvc) GetFeed(newPostSvcClient PostSvcInjector) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := GetUserId(c)
		_, userIdErr := uuid.Parse(userId)
		if userIdErr != nil {
			ResponseError(c, http.StatusBadRequest, "please provide a valid user id")
			return
		}
		page, err := GetPageQueryParam(c)
		if err != nil {
			ResponseError(c, http.StatusBadRequest, "please provide a valid page number")
			return
		}
		feed, nextPage, err := s.PostSvcGetFeed(c.Request.Context(), newPostSvcClient, userId, uint32(page))
		st := status.Convert(err)
		switch st.Code() {
		case codes.OK:
			ResponseOK(c, "successfully fetched feed", gin.H{
				"feed":      feed,
				"next_page": nextPage,
			})
		case codes.NotFound:
			ResponseError(c, http.StatusBadRequest, "please provide a valid user id")
		default:
			s.log.Error(st.Err())
			ResponseInternalError(c)
		}
	}
}

// GetPost handles GET /v1/post/:post_id
func (s *FrontendSvc) GetPost(newPostSvcClient PostSvcInjector) gin.HandlerFunc {
	return func(c *gin.Context) {
		postId := c.Param("post_id")
		_, err := uuid.Parse(postId)
		if err != nil {
			ResponseError(c, http.StatusBadRequest, "please provide a valid post id")
			return
		}
		post, err := s.PostSvcGetPost(c.Request.Context(), newPostSvcClient, postId)
		st := status.Convert(err)
		switch st.Code() {
		case codes.OK:
			ResponseOK(c, "successfully fetched post", gin.H{
				"post": post,
			})
		case codes.NotFound:
			ResponseError(c, http.StatusBadRequest, "please provide a valid post id")
		default:
			s.log.Error(st.Err())
			ResponseInternalError(c)
		}
	}
}

// GetPostLikes handles GET /v1/post/:post_id/likes
func (s *FrontendSvc) GetPostLikes(newPostSvcClient PostSvcInjector) gin.HandlerFunc {
	return func(c *gin.Context) {
		postId := c.Param("post_id")
		_, err := uuid.Parse(postId)
		if err != nil {
			ResponseError(c, http.StatusBadRequest, "please provide a valid post id")
			return
		}
		page, err := GetPageQueryParam(c)
		if err != nil {
			ResponseError(c, http.StatusBadRequest, "please provide a valid page number")
			return
		}
		likes, nextPage, err := s.PostSvcGetPostLikes(c.Request.Context(), newPostSvcClient, postId, page)
		st := status.Convert(err)
		switch st.Code() {
		case codes.OK:
			ResponseOK(c, "successfully fetched likes", gin.H{
				"likes":     likes,
				"next_page": nextPage,
			})
		case codes.NotFound:
			ResponseError(c, http.StatusBadRequest, "please provide a valid post id")
		default:
			s.log.Error(st.Err())
			ResponseInternalError(c)
		}
	}
}

// GetPostComments handles GET /v1/post/:post_id/comments
func (s *FrontendSvc) GetPostComments(newPostSvcClient PostSvcInjector) gin.HandlerFunc {
	return func(c *gin.Context) {
		postId := c.Param("post_id")
		_, err := uuid.Parse(postId)
		if err != nil {
			ResponseError(c, http.StatusBadRequest, "please provide a valid post id")
			return
		}
		page, err := GetPageQueryParam(c)
		if err != nil {
			ResponseError(c, http.StatusBadRequest, "please provide a valid page number")
			return
		}
		comments, nextPage, err := s.PostSvcGetPostComments(c.Request.Context(), newPostSvcClient, postId, page)
		st := status.Convert(err)
		switch st.Code() {
		case codes.OK:
			ResponseOK(c, "successfully fetched comments", gin.H{
				"comments":  comments,
				"next_page": nextPage,
			})
		case codes.NotFound:
			ResponseError(c, http.StatusBadRequest, "please provide a valid post id")
		default:
			s.log.Error(st.Err())
			ResponseInternalError(c)
		}
	}
}

// PostPost handles POST /v1/post/
func (s *FrontendSvc) PostPost(newPostSvcClient PostSvcInjector, newFileSvcClient FileSvcInjector) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := GetUserId(c)
		_, err := uuid.Parse(userId)
		if err != nil {
			ResponseError(c, http.StatusBadRequest, "please provide a valid user id")
			return
		}
		req := struct {
			Caption string `json:"caption" binding:"required"`
		}{}
		if err := c.ShouldBindJSON(&req); err != nil {
			ResponseError(c, http.StatusBadRequest, "please provide a caption")
			return
		}
		file, err := c.FormFile("media")
		if err != nil {
			ResponseError(c, http.StatusBadRequest, "please provide a media file")
			return
		}
		content, err := file.Open()
		if file.Size > fileSizeLimit || err != nil {
			ResponseError(c, http.StatusBadRequest, "please provide a valid media file, 8MB or less")
			return
		}
		post, err := s.PostSvcCreatePost(c.Request.Context(), newPostSvcClient, newFileSvcClient, userId, req.Caption, content)
		st := status.Convert(err)
		switch st.Code() {
		case codes.OK:
			ResponseOK(c, "successfully created post", gin.H{
				"post": post,
			})
		case codes.NotFound:
			ResponseError(c, http.StatusBadRequest, "please provide a valid user id")
		default:
			s.log.Error(st.Err())
			ResponseInternalError(c)
		}
	}
}

// PostPostLike handles POST /v1/post/:post_id/like
func (s *FrontendSvc) PostPostLike(newPostSvcClient PostSvcInjector) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := GetUserId(c)
		postId := c.Param("post_id")
		_, postIdErr := uuid.Parse(postId)
		_, userIdErr := uuid.Parse(userId)
		if postIdErr != nil || userIdErr != nil {
			ResponseError(c, http.StatusBadRequest, "please provide a valid user id")
			return
		}
		post, err := s.PostSvcCreatePostLike(c.Request.Context(), newPostSvcClient, postId, userId)
		st := status.Convert(err)
		switch st.Code() {
		case codes.OK:
			ResponseOK(c, "successfully liked post", gin.H{
				"post": post,
			})
		case codes.NotFound:
			ResponseError(c, http.StatusBadRequest, "please provide a valid user and post id")
		default:
			s.log.Error(st.Err())
			ResponseInternalError(c)
		}
	}
}

// PostComment handles POST /v1/post/:post_id/comment
func (s *FrontendSvc) PostComment(newPostSvcClient PostSvcInjector) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := GetUserId(c)
		postId := c.Param("post_id")
		_, postIdErr := uuid.Parse(postId)
		_, userIdErr := uuid.Parse(userId)
		if postIdErr != nil || userIdErr != nil {
			ResponseError(c, http.StatusBadRequest, "please provide a valid user and post id")
			return
		}
		req := struct {
			Text string `json:"text" binding:"required"`
		}{}
		if err := c.ShouldBindJSON(&req); err != nil {
			ResponseError(c, http.StatusBadRequest, "please provide a text")
			return
		}
		post, comment, err := s.PostSvcCreatePostComment(c.Request.Context(), newPostSvcClient, postId, userId, req.Text)
		st := status.Convert(err)
		switch st.Code() {
		case codes.OK:
			ResponseOK(c, "successfully commented on post", gin.H{
				"post":    post,
				"comment": comment,
			})
		case codes.NotFound:
			ResponseError(c, http.StatusBadRequest, "please provide a valid user and post id")
		default:
			s.log.Error(st.Err())
			ResponseInternalError(c)
		}
	}
}

// GetComment handles GET /v1/comment/:comment_id
func (s *FrontendSvc) GetComment(newPostSvcClient PostSvcInjector) gin.HandlerFunc {
	return func(c *gin.Context) {
		commentId := c.Param("comment_id")
		_, err := uuid.Parse(commentId)
		if err != nil {
			ResponseError(c, http.StatusBadRequest, "please provide a valid comment id")
			return
		}
		comment, err := s.PostSvcGetComment(c.Request.Context(), newPostSvcClient, commentId)
		st := status.Convert(err)
		switch st.Code() {
		case codes.OK:
			ResponseOK(c, "successfully fetched comment", gin.H{
				"comment": comment,
			})
		case codes.NotFound:
			ResponseError(c, http.StatusBadRequest, "please provide a valid comment id")
		default:
			s.log.Error(st.Err())
			ResponseInternalError(c)
		}
	}
}

// GetCommentLikes handles GET /v1/comment/:comment_id/likes
func (s *FrontendSvc) GetCommentLikes(newPostSvcClient PostSvcInjector) gin.HandlerFunc {
	return func(c *gin.Context) {
		commentId := c.Param("comment_id")
		_, err := uuid.Parse(commentId)
		if err != nil {
			ResponseError(c, http.StatusBadRequest, "please provide a valid comment id")
			return
		}
		likes, nextPage, err := s.PostSvcGetCommentLikes(c.Request.Context(), newPostSvcClient, commentId)
		st := status.Convert(err)
		switch st.Code() {
		case codes.OK:
			ResponseOK(c, "successfully fetched likes", gin.H{
				"likes":     likes,
				"next_page": nextPage,
			})
		case codes.NotFound:
			ResponseError(c, http.StatusBadRequest, "please provide a valid comment id")
		default:
			s.log.Error(st.Err())
			ResponseInternalError(c)
		}
	}
}

// PostCommentLike handles POST /v1/comment/:comment_id/like
func (s *FrontendSvc) PostCommentLike(newPostSvcClient PostSvcInjector) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := GetUserId(c)
		commentId := c.Param("comment_id")
		_, commentIdErr := uuid.Parse(commentId)
		_, userIdErr := uuid.Parse(userId)
		if commentIdErr != nil || userIdErr != nil {
			ResponseError(c, http.StatusBadRequest, "please provide a valid user and comment id")
			return
		}
		comment, err := s.PostSvcCreateCommentLike(c.Request.Context(), newPostSvcClient, commentId, userId)
		st := status.Convert(err)
		switch st.Code() {
		case codes.OK:
			ResponseOK(c, "successfully liked comment", gin.H{
				"comment": comment,
			})
		case codes.NotFound:
			ResponseError(c, http.StatusBadRequest, "please provide a valid comment and user id")
		default:
			s.log.Error(st.Err())
			ResponseInternalError(c)
		}
	}
}
