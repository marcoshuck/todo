package server

import (
	"context"
	tasksv1 "github.com/marcoshuck/todo/api/tasks/v1"
	"github.com/marcoshuck/todo/pkg/conf"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/multierr"
	"go.uber.org/zap"
	"google.golang.org/grpc/health"
	healthv1 "google.golang.org/grpc/health/grpc_health_v1"
	"gorm.io/gorm"
	"io"
	"net"
	"net/http"
	"time"
)

// Services groups all the services exposed by a single gRPC Server.
type Services struct {
	TasksWriter tasksv1.TasksWriterServiceServer
	TasksReader tasksv1.TasksReaderServiceServer
	Health      *health.Server
}

// shutDowner holds a method to gracefully shut down a service or integration.
type shutDowner interface {
	// Shutdown releases any held computational resources.
	Shutdown(ctx context.Context) error
}

// grpcServer holds the method to serve a gRPC server using a net.Listener.
type grpcServer interface {
	// Serve serves a gRPC server through net.Listener until an error occurs.
	Serve(net.Listener) error
}

// Application abstracts all the functional components to be run by the server.
type Application struct {
	server         grpcServer
	listener       net.Listener
	logger         *zap.Logger
	db             *gorm.DB
	services       Services
	tracerProvider trace.TracerProvider
	meterProvider  metric.MeterProvider
	shutdown       []shutDowner
	closer         []io.Closer
	cfg            conf.ServerConfig
	metricsServer  *http.Server
}

// Run serves the application services.
func (app Application) Run(ctx context.Context) error {
	go app.checkHealth(ctx)
	go app.serveMetrics()

	app.logger.Info("Running application")
	return app.server.Serve(app.listener)
}

// Shutdown releases any held resources by dependencies of this Application.
func (app Application) Shutdown(ctx context.Context) error {
	var err error
	for _, downer := range app.shutdown {
		if downer == nil {
			continue
		}
		err = multierr.Append(err, downer.Shutdown(ctx))
	}
	for _, closer := range app.closer {
		if closer == nil {
			continue
		}
		err = multierr.Append(err, closer.Close())
	}
	return err
}

func (app Application) checkHealth(ctx context.Context) {
	app.logger.Info("Running health service")
	for {
		if ctx.Err() != nil {
			return
		}
		app.services.Health.SetServingStatus("app.db", app.checkDatabaseHealth())
		time.Sleep(10 * time.Second)
	}
}

func (app Application) checkDatabaseHealth() healthv1.HealthCheckResponse_ServingStatus {
	state := healthv1.HealthCheckResponse_SERVING
	db, err := app.db.DB()
	if err != nil {
		state = healthv1.HealthCheckResponse_NOT_SERVING
	}
	if err = db.Ping(); err != nil {
		state = healthv1.HealthCheckResponse_NOT_SERVING
	}
	return state
}

func (app Application) serveMetrics() {
	if err := app.metricsServer.ListenAndServe(); err != nil {
		app.logger.Error("failed to listen and server to metrics server", zap.Error(err))
	}
}
