package telemetry

import (
	"context"
	"fmt"
	"github.com/gojaguar/jaguar/config"
	"github.com/marcoshuck/todo/pkg/conf"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	trace_sdk "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

func SetupTracing(cfg config.Config, tracing conf.Tracing) (trace.TracerProvider, trace_sdk.SpanExporter, error) {
	if !tracing.Enabled {
		return trace.NewNoopTracerProvider(), nil, nil
	}

	var tracerProvider trace.TracerProvider
	var traceExporter trace_sdk.SpanExporter
	switch cfg.Environment {
	case "production", "staging":
		var err error
		tracerProvider, traceExporter, err = newTracing(cfg, tracing)
		if err != nil {
			return nil, nil, err
		}
	default:
		tracerProvider = trace.NewNoopTracerProvider()
	}

	return tracerProvider, traceExporter, nil
}

func newTracing(cfg config.Config, tracing conf.Tracing) (trace.TracerProvider, trace_sdk.SpanExporter, error) {
	ctx := context.Background()
	res, err := newResource(ctx, cfg)
	if err != nil {
		return nil, nil, err
	}
	conn, err := newConn(ctx, tracing.Address())
	if err != nil {
		return nil, nil, err
	}
	traceExporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create trace exporter: %w", err)
	}
	bsp := trace_sdk.NewBatchSpanProcessor(traceExporter)
	tracerProvider := trace_sdk.NewTracerProvider(
		trace_sdk.WithSampler(trace_sdk.AlwaysSample()),
		trace_sdk.WithResource(res),
		trace_sdk.WithSpanProcessor(bsp),
	)
	return tracerProvider, traceExporter, nil
}
