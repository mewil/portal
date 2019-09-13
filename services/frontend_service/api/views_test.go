package api_test

import (
	"github.com/gin-gonic/gin"
)

type mockRequestResponse struct {
	path    string
	method  string
	reqBody gin.H
	status  int
	check   func(gin.H)
}
