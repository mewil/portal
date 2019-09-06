package api

import (
	"net/http"
	"strconv"

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
		user, err := s.UserGetUser(c.Request.Context(), newUserSvcClient, userId)
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
		page, err := strconv.ParseUint(c.Param("page"), 10, 32)
		if err != nil {
			ResponseError(c, http.StatusBadRequest, "please provide a valid page number")
			return
		}
		followers, err := s.UserGetFollowers(c.Request.Context(), newUserSvcClient, userId, uint32(page))
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
		page, err := strconv.ParseUint(c.Param("page"), 10, 32)
		if err != nil {
			ResponseError(c, http.StatusBadRequest, "please provide a valid page number")
			return
		}
		following, err := s.UserGetFollowing(c.Request.Context(), newUserSvcClient, userId, uint32(page))
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
