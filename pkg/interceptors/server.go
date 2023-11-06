package interceptors

import (
	grpc_logging "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"github.com/marcoshuck/todo/pkg/telemetry"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

func NewServerUnaryInterceptors(telemeter telemetry.Telemetry) grpc.ServerOption {
	return grpc.ChainUnaryInterceptor(
		otelgrpc.UnaryServerInterceptor(
			otelgrpc.WithTracerProvider(telemeter.TracerProvider),
			otelgrpc.WithMeterProvider(telemeter.MeterProvider),
		),
		grpc_logging.UnaryServerInterceptor(interceptorLogger(telemeter.Logger)),
		grpc_recovery.UnaryServerInterceptor(grpc_recovery.WithRecoveryHandler(RecoveryHandler(telemeter.Logger))),
	)
}

func NewServerStreamInterceptors(telemeter telemetry.Telemetry) grpc.ServerOption {
	return grpc.ChainStreamInterceptor(
		otelgrpc.StreamServerInterceptor(
			otelgrpc.WithTracerProvider(telemeter.TracerProvider),
			otelgrpc.WithMeterProvider(telemeter.MeterProvider),
		),
		grpc_logging.StreamServerInterceptor(interceptorLogger(telemeter.Logger)),
		grpc_recovery.StreamServerInterceptor(grpc_recovery.WithRecoveryHandler(RecoveryHandler(telemeter.Logger))),
	)
}
