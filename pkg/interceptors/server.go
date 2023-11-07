package interceptors

import (
	"github.com/bufbuild/protovalidate-go"
	grpc_logging "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	grpc_validate "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/protovalidate"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"github.com/marcoshuck/todo/pkg/telemetry"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

func NewServerInterceptors(telemeter telemetry.Telemetry, validator *protovalidate.Validator) []grpc.ServerOption {
	var opts []grpc.ServerOption
	return append(opts, newServerUnaryInterceptors(telemeter, validator), newServerStreamInterceptors(telemeter, validator))
}

func newServerUnaryInterceptors(telemeter telemetry.Telemetry, validator *protovalidate.Validator) grpc.ServerOption {
	var interceptors []grpc.UnaryServerInterceptor

	interceptors = append(interceptors, otelgrpc.UnaryServerInterceptor(
		otelgrpc.WithTracerProvider(telemeter.TracerProvider),
		otelgrpc.WithMeterProvider(telemeter.MeterProvider),
		otelgrpc.WithPropagators(telemeter.Propagator),
	))

	if validator != nil {
		interceptors = append(interceptors, grpc_validate.UnaryServerInterceptor(validator))
	}

	if telemeter.Logger != nil {
		interceptors = append(interceptors,
			grpc_logging.UnaryServerInterceptor(interceptorLogger(telemeter.Logger)),
			grpc_recovery.UnaryServerInterceptor(grpc_recovery.WithRecoveryHandler(RecoveryHandler(telemeter.Logger))),
		)
	}

	return grpc.ChainUnaryInterceptor(interceptors...)
}

func newServerStreamInterceptors(telemeter telemetry.Telemetry, validator *protovalidate.Validator) grpc.ServerOption {
	var interceptors []grpc.StreamServerInterceptor

	interceptors = append(interceptors, otelgrpc.StreamServerInterceptor(
		otelgrpc.WithTracerProvider(telemeter.TracerProvider),
		otelgrpc.WithMeterProvider(telemeter.MeterProvider),
		otelgrpc.WithPropagators(telemeter.Propagator),
	))

	if validator != nil {
		interceptors = append(interceptors, grpc_validate.StreamServerInterceptor(validator))
	}

	if telemeter.Logger != nil {
		interceptors = append(interceptors,
			grpc_logging.StreamServerInterceptor(interceptorLogger(telemeter.Logger)),
			grpc_recovery.StreamServerInterceptor(grpc_recovery.WithRecoveryHandler(RecoveryHandler(telemeter.Logger))),
		)
	}

	return grpc.ChainStreamInterceptor(interceptors...)
}
