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
	path      string
	method    string
	body      gin.H
	resStatus int
	resBody   gin.H
}

func TestAuthViews(t *testing.T) {
	testRequestResponses := []mockRequestResponse{
		{
			"/signup",
			http.MethodPost,
			gin.H{
				"username": "",
				"name":     "",
				"email":    "",
			},
			http.StatusBadRequest,
			gin.H{},
		},
	}

	log, _ := logger.NewLogger()
	s := FrontendSvc{
		log:       log,
		jwtSecret: "secret",
	}

	r := gin.New()
	r.POST("/signin", s.PostAuthSignIn(injectMockAuthSvcClient()))
	r.POST("/signup", s.PostAuthSignUp(injectMockAuthSvcClient(), injectMockUserSvcClient()))

	for _, test := range testRequestResponses {
		w := httptest.NewRecorder()
		body, _ := json.Marshal(test.body)
		req, _ := http.NewRequest(test.method, test.path, bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(w, req)

		assert.Equal(t, test.resStatus, w.Code)

	}
}
