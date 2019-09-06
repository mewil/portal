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
			s.log.Info("")
			ResponseError(c, http.StatusBadRequest, "")
			return
		}
		token, id, err := s.AuthSignIn(c.Request.Context(), newAuthSvcClient, req.Email, req.Password)
		st := status.Convert(err)
		switch st.Code() {
		case codes.OK:
			ResponseOK(c, "successfully logged in user", gin.H{
				"token": token,
				"id":    id,
			})
		default:
			ResponseError(c, http.StatusInternalServerError, "")
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
			s.log.Info("")
			ResponseError(c, http.StatusBadRequest, "")
			return
		}
		user, token, err := s.AuthSignUp(c.Request.Context(), newAuthSvcClient, newUserSvcClient, req.Username, req.Name, req.Email, req.Password)
		st := status.Convert(err)
		switch st.Code() {
		case codes.OK:
			ResponseOK(c, "successfully logged in user", gin.H{
				"user":  user,
				"token": token,
			})
		default:
			ResponseError(c, http.StatusInternalServerError, "please try again")
		}
	}
}
