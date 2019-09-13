package grpc_utils

import (
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/mewil/portal/common/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func NewServer(log logger.Logger) (*grpc.Server, error) {
	return grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_zap.StreamServerInterceptor(log.(*zap.SugaredLogger).Desugar()),
			grpc_recovery.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_zap.UnaryServerInterceptor(log.(*zap.SugaredLogger).Desugar()),
			grpc_recovery.UnaryServerInterceptor(),
		)),
	), nil
}
