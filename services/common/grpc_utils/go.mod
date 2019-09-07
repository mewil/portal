module github.com/mewil/portal/common/grpc_utils

go 1.13

require (
	github.com/grpc-ecosystem/go-grpc-middleware v1.1.0
	github.com/mewil/portal/common/logger v0.0.0
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/opentracing/opentracing-go v1.1.0
	github.com/prometheus/client_golang v1.1.0 // indirect
	github.com/uber/jaeger-client-go v2.19.0+incompatible
	github.com/uber/jaeger-lib v2.2.0+incompatible
	go.uber.org/zap v1.10.0
	google.golang.org/grpc v1.19.0
)

replace github.com/mewil/portal/common/logger v0.0.0 => ../logger
