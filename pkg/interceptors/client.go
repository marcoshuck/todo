package interceptors

import (
	grpc_logging "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"github.com/marcoshuck/todo/pkg/telemetry"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"time"
)

func NewClientUnaryInterceptors(telemeter telemetry.Telemetry) grpc.DialOption {
	return grpc.WithChainUnaryInterceptor(
		otelgrpc.UnaryClientInterceptor(
			otelgrpc.WithTracerProvider(telemeter.TracerProvider),
			otelgrpc.WithMeterProvider(telemeter.MeterProvider),
		),
		grpc_logging.UnaryClientInterceptor(interceptorLogger(telemeter.Logger)),
		retry.UnaryClientInterceptor(
			retry.WithCodes(codes.ResourceExhausted, codes.Unavailable),
			retry.WithMax(10),
			retry.WithBackoff(retry.BackoffExponential(50*time.Millisecond)),
		),
	)
}

func NewClientStreamInterceptors(telemeter telemetry.Telemetry) grpc.DialOption {
	return grpc.WithChainStreamInterceptor(
		otelgrpc.StreamClientInterceptor(
			otelgrpc.WithTracerProvider(telemeter.TracerProvider),
			otelgrpc.WithMeterProvider(telemeter.MeterProvider),
		),
		grpc_logging.StreamClientInterceptor(interceptorLogger(telemeter.Logger)),
		retry.StreamClientInterceptor(
			retry.WithCodes(codes.ResourceExhausted, codes.Unavailable),
			retry.WithMax(10),
			retry.WithBackoff(retry.BackoffExponential(50*time.Millisecond)),
		),
	)
}
