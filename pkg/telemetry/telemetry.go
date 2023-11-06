package telemetry

import (
	"github.com/gojaguar/jaguar/config"
	"github.com/marcoshuck/todo/pkg/conf"
	"go.opentelemetry.io/otel/metric"
	metric_sdk "go.opentelemetry.io/otel/sdk/metric"
	trace_sdk "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

type Telemetry struct {
	Logger         *zap.Logger
	TracerProvider trace.TracerProvider
	TraceExporter  trace_sdk.SpanExporter
	MeterProvider  metric.MeterProvider
	MeterExporter  metric_sdk.Exporter
}

func SetupTelemetry(cfg config.Config, tracing conf.Tracing, metrics conf.Metrics) (Telemetry, error) {
	var t Telemetry
	var err error
	t.Logger, err = SetupLogger(cfg)
	if err != nil {
		return Telemetry{}, err
	}

	t.TracerProvider, t.TraceExporter, err = SetupTracing(cfg, tracing)
	if err != nil {
		return Telemetry{}, err
	}

	t.MeterProvider, t.MeterExporter, err = SetupMetrics(cfg, metrics)
	if err != nil {
		return Telemetry{}, err
	}
	return t, nil
}
