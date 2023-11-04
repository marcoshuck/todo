package interceptors

import (
	grpc_logging "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func NewServerUnaryInterceptors(logger *zap.Logger, tracerProvider trace.TracerProvider, meterProvider metric.MeterProvider) grpc.ServerOption {
	return grpc.ChainUnaryInterceptor(
		otelgrpc.UnaryServerInterceptor(
			otelgrpc.WithTracerProvider(tracerProvider),
			otelgrpc.WithMeterProvider(meterProvider),
		),
		grpc_logging.UnaryServerInterceptor(interceptorLogger(logger)),
		grpc_recovery.UnaryServerInterceptor(grpc_recovery.WithRecoveryHandler(RecoveryHandler)),
	)
}

func NewServerStreamInterceptors(logger *zap.Logger, tracerProvider trace.TracerProvider, meterProvider metric.MeterProvider) grpc.ServerOption {
	return grpc.ChainStreamInterceptor(
		otelgrpc.StreamServerInterceptor(
			otelgrpc.WithTracerProvider(tracerProvider),
			otelgrpc.WithMeterProvider(meterProvider),
		),
		grpc_logging.StreamServerInterceptor(interceptorLogger(logger)),
		grpc_recovery.StreamServerInterceptor(grpc_recovery.WithRecoveryHandler(RecoveryHandler)),
	)
}
