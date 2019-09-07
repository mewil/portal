package grpc_utils

import (
	"io"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-lib/metrics/prometheus"
)

func NewTracer() (opentracing.Tracer, io.Closer, error) {
	cfg, err := config.FromEnv()
	if err != nil {
		return nil, nil, err
	}
	return cfg.NewTracer(
		config.Metrics(prometheus.New()),
	)
}
