package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/static"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/mewil/portal/common/logger"
	"github.com/mewil/portal/frontend/api"
	"go.uber.org/zap"
)

const (
	serverTimeout        = 10 * time.Second
	serverPort           = 8000
	serverMaxHeaderBytes = 1 << 20
)

func main() {
	l, err := logger.NewLogger("frontend_service")
	if err != nil {
		panic(err)
	}
	r := newRouter(l)
	r.Use(static.Serve("/", static.LocalFile("/app", true)))
	api.NewAPI(l, r.Group("/v1"))
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", serverPort),
		Handler:        r,
		ReadTimeout:    serverTimeout,
		WriteTimeout:   serverTimeout,
		MaxHeaderBytes: serverMaxHeaderBytes,
	}
	s.ListenAndServe()
}

func newRouter(log logger.Logger) *gin.Engine {
	r := gin.Default()
	r.Use(ginzap.Ginzap(log.(*zap.SugaredLogger).Desugar(), time.RFC3339Nano, true))
	r.Use(ginzap.RecoveryWithZap(log.(*zap.SugaredLogger).Desugar(), true))
	return r
}
