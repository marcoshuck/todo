package main

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	tasksv1 "github.com/marcoshuck/todo/api/tasks/v1"
	"github.com/marcoshuck/todo/pkg/conf"
	"github.com/marcoshuck/todo/pkg/interceptors"
	"github.com/marcoshuck/todo/pkg/telemetry"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"time"
)

func main() {
	ctx := context.Background()

	cfg, err := conf.ReadClientConfig()

	t, err := telemetry.SetupTelemetry(cfg.Config, cfg.Tracing, cfg.Metrics)
	if err != nil {
		log.Fatalln("Failed to initialize t")
	}

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalln("Failed to initialize logger")
	}
	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		interceptors.NewClientUnaryInterceptors(t.Logger, t.TracerProvider, t.MeterProvider),
		interceptors.NewClientStreamInterceptors(t.Logger, t.TracerProvider, t.MeterProvider),
	}
	err = tasksv1.RegisterTasksServiceHandlerFromEndpoint(ctx, mux, cfg.ServerAddress, opts)
	if err != nil {
		log.Fatalln("Failed to register tasks service:", err)
	}
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
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
	r.Mount("/", mux)
	addr := fmt.Sprintf(":%d", cfg.Port)
	logger.Info("Listening...", zap.String("address", addr))
	if err := http.ListenAndServe(addr, r); err != nil {
		logger.Fatal("Failed to listen and serve", zap.Error(err))
	}
}
