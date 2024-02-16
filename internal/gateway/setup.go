package gateway

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/marcoshuck/todo/api/tasks/v1"
	"github.com/marcoshuck/todo/internal/conf"
	"github.com/marcoshuck/todo/internal/interceptors"
	"github.com/marcoshuck/todo/internal/telemetry"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func Setup(ctx context.Context, cfg conf.GatewayConfig) (Gateway, error) {
	telemeter, err := telemetry.SetupTelemetry(cfg.Config, cfg.Tracing, cfg.Metrics)
	if err != nil {
		log.Fatalln("Failed to initialize telemetry:", err)
	}

	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		interceptors.NewClientUnaryInterceptors(telemeter),
		interceptors.NewClientStreamInterceptors(telemeter),
		grpc.WithStatsHandler(
			otelgrpc.NewClientHandler(
				otelgrpc.WithTracerProvider(telemeter.TracerProvider),
				otelgrpc.WithMeterProvider(telemeter.MeterProvider),
				otelgrpc.WithPropagators(telemeter.Propagator),
			),
		),
	}
	err = tasksv1.RegisterTasksWriterServiceHandlerFromEndpoint(ctx, mux, cfg.ServerAddress, opts)
	if err != nil {
		log.Fatalln("Failed to register tasks service:", err)
	}
	err = tasksv1.RegisterTasksReaderServiceHandlerFromEndpoint(ctx, mux, cfg.ServerAddress, opts)
	if err != nil {
		log.Fatalln("Failed to register tasks service:", err)
	}
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(middleware.Heartbeat("/livez"))
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	r.Mount("/metrics", promhttp.Handler())
	r.Mount("/", mux)

	return Gateway{
		Telemeter: telemeter,
		mux:       mux,
		handler:   r,
		Config:    cfg,
	}, nil
}
