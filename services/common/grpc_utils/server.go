package grpc_utils

import (
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/mewil/portal/common/logger"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func NewServer(log logger.Logger) (*grpc.Server, error) {
	tracer, closer, err := NewTracer()
	defer closer.Close()
	if err != nil {
		return nil, err
	}
	opentracing.SetGlobalTracer(tracer)
	return grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_opentracing.StreamServerInterceptor(grpc_opentracing.WithTracer(tracer)),
			grpc_zap.StreamServerInterceptor(log.(*zap.SugaredLogger).Desugar()),
			grpc_recovery.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_opentracing.UnaryServerInterceptor(grpc_opentracing.WithTracer(tracer)),
			grpc_zap.UnaryServerInterceptor(log.(*zap.SugaredLogger).Desugar()),
			grpc_recovery.UnaryServerInterceptor(),
		)),
	), nil
}
