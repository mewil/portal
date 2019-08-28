package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/static"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	stdOut               = "stdout"
	serverTimeout        = 10 * time.Second
	serverPort           = 8000
	serverMaxHeaderBytes = 1 << 20
)

func main() {
	l := newLogger(stdOut)
	r := newRouter(l)
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", serverPort),
		Handler:        r,
		ReadTimeout:    serverTimeout,
		WriteTimeout:   serverTimeout,
		MaxHeaderBytes: serverMaxHeaderBytes,
	}
	s.ListenAndServe()
}

func newRouter(logger *zap.Logger) *gin.Engine {
	r := gin.Default()
	r.Use(ginzap.Ginzap(logger, time.RFC3339Nano, true))
	r.Use(ginzap.RecoveryWithZap(logger, true))

	r.Use(static.Serve("/", static.LocalFile("/app", true)))

	api := r.Group("/v1")
	api.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hellow world",
		})
	})
	return r
}

func newLogger(outputDest ...string) *zap.Logger {
	cfg := zap.Config{
		Encoding:         "json",
		Level:            zap.NewAtomicLevelAt(zap.ErrorLevel),
		OutputPaths:      outputDest,
		ErrorOutputPaths: outputDest,
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   "message",
			LevelKey:     "level",
			EncodeLevel:  zapcore.CapitalLevelEncoder,
			TimeKey:      "time",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	return logger
}
