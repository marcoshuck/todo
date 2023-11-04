package interceptors

import (
	grpc_logging "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func NewClientUnaryInterceptors(logger *zap.Logger, tracerProvider trace.TracerProvider, meterProvider metric.MeterProvider) grpc.DialOption {
	return grpc.WithChainUnaryInterceptor(
		otelgrpc.UnaryClientInterceptor(
			otelgrpc.WithTracerProvider(tracerProvider),
			otelgrpc.WithMeterProvider(meterProvider),
		),
		grpc_logging.UnaryClientInterceptor(interceptorLogger(logger)),
	)
}

func NewClientStreamInterceptors(logger *zap.Logger, tracerProvider trace.TracerProvider, meterProvider metric.MeterProvider) grpc.DialOption {
	return grpc.WithChainStreamInterceptor(
		otelgrpc.StreamClientInterceptor(
			otelgrpc.WithTracerProvider(tracerProvider),
			otelgrpc.WithMeterProvider(meterProvider),
		),
		grpc_logging.StreamClientInterceptor(interceptorLogger(logger)),
	)
}
