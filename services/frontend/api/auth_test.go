package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	// "github.com/mewil/portal/pb"
	"github.com/mewil/portal/common/logger"
	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
)

type mockRequestResponse struct {
	path    string
	method  string
	reqBody gin.H
	status  int
	check   func(gin.H)
}

func TestAuthViews(t *testing.T) {
	tests := []mockRequestResponse{
		{
			"/signup",
			http.MethodPost,
			gin.H{
				"username": "u",
				"name":     "u",
				"email":    "u",
			},
			http.StatusBadRequest,
			func(res gin.H) {
				assert.False(t, res["status"].(bool))
				assert.Equal(t, "please provide an username, name, email, and password", res["message"].(string))
			},
		},
		{
			"/signup",
			http.MethodPost,
			gin.H{
				"username": "user1",
				"name":     "User",
				"email":    "invalid email format",
				"password": "invalid password",
			},
			http.StatusBadRequest,
			func(res gin.H) {
				assert.False(t, res["status"].(bool))
				assert.Equal(t, "please provide a valid username, email, and password", res["message"].(string))
			},
		},
		{
			"/signup",
			http.MethodPost,
			gin.H{
				"username": "user2",
				"name":     "database error",
				"email":    "correct@email.format",
				"password": "valid_password",
			},
			http.StatusInternalServerError,
			func(res gin.H) {
				assert.False(t, res["status"].(bool))
				assert.Equal(t, "an unexpected error occurred, please try again", res["message"].(string))
			},
		},
		{
			"/signup",
			http.MethodPost,
			gin.H{
				"username": "user1",
				"name":     "User",
				"email":    "correct@email.format",
				"password": "valid_password",
			},
			http.StatusOK,
			func(res gin.H) {
				assert.True(t, res["status"].(bool))
				assert.Equal(t, "successfully signed up user", res["message"].(string))
				assert.Contains(t, res, "data")
				data := res["data"].(map[string]interface{})
				assert.Contains(t, data, "user")
				user := data["user"].(map[string]interface{})
				assert.Equal(t, "user1", user["username"])
				assert.Equal(t, "User", user["name"])
				assert.Equal(t, "correct@email.format", user["email"])
				assert.Contains(t, user, "user_id")
				assert.Contains(t, data, "token")
			},
		},
	}

	log, _ := logger.NewLogger()
	s := FrontendSvc{
		log:       log,
		jwtSecret: "secret",
	}
	as := newMockAuthSvc()
	us := newMockUserSvc()

	r := gin.New()
	r.POST("/signin", s.PostAuthSignIn(as.injectMockAuthSvcClient()))
	r.POST("/signup", s.PostAuthSignUp(as.injectMockAuthSvcClient(), us.injectMockUserSvcClient()))

	for _, test := range tests {
		w := httptest.NewRecorder()
		body, _ := json.Marshal(test.reqBody)
		req, _ := http.NewRequest(test.method, test.path, bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(w, req)

		assert.Equal(t, test.status, w.Code)
		res := gin.H{}
		json.Unmarshal(w.Body.Bytes(), &res)
		test.check(res)
	}
}
