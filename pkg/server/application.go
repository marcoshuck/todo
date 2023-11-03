package server

import (
	"context"
	tasksv1 "github.com/marcoshuck/todo/api/tasks/v1"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/multierr"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"io"
	"net"
)

// Services groups all the services exposed by a single gRPC Server.
type Services struct {
	Tasks tasksv1.TasksServiceServer
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
}

// Serve serves the application services.
func (app Application) Serve() error {
	return app.server.Serve(app.listener)
}

// Shutdown releases any held resources by dependencies of this Application.
func (app Application) Shutdown() error {
	ctx := context.Background()
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
