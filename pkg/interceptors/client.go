package interceptors

import (
	grpc_logging "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"github.com/marcoshuck/todo/pkg/telemetry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"time"
)

func NewClientUnaryInterceptors(telemeter telemetry.Telemetry) grpc.DialOption {
	return grpc.WithChainUnaryInterceptor(
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
		grpc_logging.StreamClientInterceptor(interceptorLogger(telemeter.Logger)),
		retry.StreamClientInterceptor(
			retry.WithCodes(codes.ResourceExhausted, codes.Unavailable),
			retry.WithMax(10),
			retry.WithBackoff(retry.BackoffExponential(50*time.Millisecond)),
		),
	)
}
