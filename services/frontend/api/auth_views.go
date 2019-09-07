package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// PostAuthSignIn handles POST /v1/auth/signin
func (s *FrontendSvc) PostAuthSignIn(newAuthSvcClient AuthSvcInjector) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := struct {
			Email    string `json:"email" binding:"required"`
			Password string `json:"password" binding:"required"`
		}{}
		if err := c.ShouldBindJSON(&req); err != nil {
			ResponseError(c, http.StatusBadRequest, "please provide an email and password")
			return
		}
		token, id, err := s.AuthSvcSignIn(c.Request.Context(), newAuthSvcClient, req.Email, req.Password)
		st := status.Convert(err)
		switch st.Code() {
		case codes.OK:
			ResponseOK(c, "successfully signed in user", gin.H{
				"token": token,
				"id":    id,
			})
		case codes.Unauthenticated:
			ResponseError(c, http.StatusUnauthorized, "please provide a valid email and password")
		default:
			s.log.Error(st.Err())
			ResponseInternalError(c)
		}
	}
}

// PostAuthSignUp handles POST /v1/auth/signup
func (s *FrontendSvc) PostAuthSignUp(newAuthSvcClient AuthSvcInjector, newUserSvcClient UserSvcInjector) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := struct {
			Username string `json:"username" binding:"required"`
			Name     string `json:"name" binding:"required"`
			Email    string `json:"email" binding:"required"`
			Password string `json:"password" binding:"required"`
		}{}
		if err := c.ShouldBindJSON(&req); err != nil {
			ResponseError(c, http.StatusBadRequest, "please provide an username, name, email, and password")
			return
		}
		user, token, err := s.AuthSvcSignUp(c.Request.Context(), newAuthSvcClient, newUserSvcClient, req.Username, req.Name, req.Email, req.Password)
		st := status.Convert(err)
		switch st.Code() {
		case codes.OK:
			s.log.Info("successfully signed up user " + user.UserId)
			ResponseOK(c, "successfully signed up user", gin.H{
				"user":  user,
				"token": token,
			})
		case codes.InvalidArgument:
			ResponseError(c, http.StatusBadRequest, "please provide a valid username, email, and password")
		default:
			s.log.Error(st.Err())
			ResponseInternalError(c)
		}
	}
}
