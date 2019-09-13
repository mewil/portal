package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/mewil/portal/common/logger"
	"github.com/mewil/portal/common/middleware"
	"github.com/mewil/portal/frontend_service/api"
)

const (
	serverTimeout        = 10 * time.Second
	serverMaxHeaderBytes = 1 << 20
)

func main() {
	l, err := logger.NewLogger("frontend_service")
	if err != nil {
		panic(err)
	}
	r := newRouter(l)
	r.Use(static.Serve("/", static.LocalFile("/app", true)))
	api.NewAPI(
		l,
		r.Group("/v1"),
		os.Getenv("JWT_SECRET"),
		os.Getenv("AUTH_SVC_ADDR"),
		os.Getenv("USER_SVC_ADDR"),
		os.Getenv("POST_SVC_ADDR"),
		os.Getenv("FILE_SVC_ADDR"),
	)
	s := &http.Server{
		Addr:           fmt.Sprintf(":%s", os.Getenv("PORT")),
		Handler:        r,
		ReadTimeout:    serverTimeout,
		WriteTimeout:   serverTimeout,
		MaxHeaderBytes: serverMaxHeaderBytes,
	}
	s.ListenAndServe()
}

func newRouter(log logger.Logger) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.LogMiddleware(log))
	r.Use(middleware.RecoveryMiddleware(log))
	return r
}
