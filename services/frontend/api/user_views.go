package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetUser handles GET /v1/user/:user_id
func (s *FrontendSvc) GetUser(newUserSvcClient UserSvcInjector) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("user_id")
		_, err := uuid.Parse(userId)
		if err != nil {
			ResponseError(c, http.StatusBadRequest, "please provide a valid user id")
			return
		}
		user, err := s.UserSvcGetUser(c.Request.Context(), newUserSvcClient, userId)
		st := status.Convert(err)
		switch st.Code() {
		case codes.OK:
			ResponseOK(c, "successfully fetched user", gin.H{
				"user": user,
			})
		case codes.NotFound:
			ResponseError(c, http.StatusBadRequest, "please provide a valid user id")
		default:
			s.log.Error(st.Err())
			ResponseInternalError(c)
		}
	}
}

// GetUserFollowers handles GET /v1/user/:user_id/followers
func (s *FrontendSvc) GetUserFollowers(newUserSvcClient UserSvcInjector) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("user_id")
		_, err := uuid.Parse(userId)
		if err != nil {
			ResponseError(c, http.StatusBadRequest, "please provide a valid user id")
			return
		}
		page, err := GetPageQueryParam(c)
		if err != nil {
			ResponseError(c, http.StatusBadRequest, "please provide a valid page number")
			return
		}
		followers, err := s.UserSvcGetFollowers(c.Request.Context(), newUserSvcClient, userId, uint32(page))
		st := status.Convert(err)
		switch st.Code() {
		case codes.OK:
			ResponseOK(c, "successfully fetched followers", gin.H{
				"followers": followers,
			})
		case codes.NotFound:
			ResponseError(c, http.StatusBadRequest, "please provide a valid user id")
		default:
			s.log.Error(st.Err())
			ResponseInternalError(c)
		}
	}
}

// GetUserFollowing handles GET /v1/user/:user_id/following
func (s *FrontendSvc) GetUserFollowing(newUserSvcClient UserSvcInjector) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("user_id")
		_, err := uuid.Parse(userId)
		if err != nil {
			ResponseError(c, http.StatusBadRequest, "please provide a valid user id")
			return
		}
		page, err := GetPageQueryParam(c)
		if err != nil {
			ResponseError(c, http.StatusBadRequest, "please provide a valid page number")
			return
		}
		following, err := s.UserSvcGetFollowing(c.Request.Context(), newUserSvcClient, userId, uint32(page))
		st := status.Convert(err)
		switch st.Code() {
		case codes.OK:
			ResponseOK(c, "successfully fetched following", gin.H{
				"following": following,
			})
		case codes.NotFound:
			ResponseError(c, http.StatusBadRequest, "please provide a valid user id")
		default:
			s.log.Error(st.Err())
			ResponseInternalError(c)
		}
	}
}

// GetUserProfile handles GET /v1/user/:user_id/profile
func (s *FrontendSvc) GetUserProfile(newUserSvcClient UserSvcInjector, newPostSvcClient PostSvcInjector) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("user_id")
		_, err := uuid.Parse(userId)
		if err != nil {
			ResponseError(c, http.StatusBadRequest, "please provide a valid user id")
			return
		}
		page, err := GetPageQueryParam(c)
		if err != nil {
			ResponseError(c, http.StatusBadRequest, "please provide a valid page number")
			return
		}
		user, posts, nextPage, err := s.UserSvcGetProfile(c.Request.Context(), newUserSvcClient, newPostSvcClient, userId, uint32(page))
		st := status.Convert(err)
		switch st.Code() {
		case codes.OK:
			ResponseOK(c, "successfully fetched profile", gin.H{
				"user":      user,
				"posts":     posts,
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

// PostUserFollow handles POST /v1/user/:user_id/follow
func (s *FrontendSvc) PostUserFollow(newUserSvcClient UserSvcInjector) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := GetUserId(c)
		followId := c.Param("user_id")
		_, followIdErr := uuid.Parse(followId)
		_, userIdErr := uuid.Parse(userId)
		if followIdErr != nil || userIdErr != nil {
			ResponseError(c, http.StatusBadRequest, "please provide a valid user and follow id")
			return
		}
		user, following, err := s.UserSvcCreateFollow(c.Request.Context(), newUserSvcClient, userId, followId)
		st := status.Convert(err)
		switch st.Code() {
		case codes.OK:
			ResponseOK(c, "successfully followed user", gin.H{
				"user":           user,
				"following_user": following,
			})
		case codes.NotFound:
			ResponseError(c, http.StatusBadRequest, "please provide a valid user and follow id")
		default:
			s.log.Error(st.Err())
			ResponseInternalError(c)
		}
	}
}

// PostUserUnfollow handles POST /v1/user/:user_id/unfollow
func (s *FrontendSvc) PostUserUnfollow(newUserSvcClient UserSvcInjector) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := GetUserId(c)
		followId := c.Param("user_id")
		_, followIdErr := uuid.Parse(followId)
		_, userIdErr := uuid.Parse(userId)
		if followIdErr != nil || userIdErr != nil {
			ResponseError(c, http.StatusBadRequest, "please provide a valid user and follow id")
			return
		}
		user, following, err := s.UserSvcRemoveFollow(c.Request.Context(), newUserSvcClient, userId, followId)
		st := status.Convert(err)
		switch st.Code() {
		case codes.OK:
			ResponseOK(c, "successfully unfollowed user", gin.H{
				"user":           user,
				"following_user": following,
			})
		case codes.NotFound:
			ResponseError(c, http.StatusBadRequest, "please provide a valid user and follow id")
		default:
			s.log.Error(st.Err())
			ResponseInternalError(c)
		}
	}
}
