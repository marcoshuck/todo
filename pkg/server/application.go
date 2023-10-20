package server

import (
	tasksv1 "github.com/marcoshuck/todo/api/tasks/v1"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"net"
)

// Services groups all the services exposed by a single gRPC Server.
type Services struct {
	Tasks tasksv1.TasksServiceServer
}

// grpcServer holds the method to serve a gRPC server using a net.Listener.
type grpcServer interface {
	// Serve serves a gRPC server through net.Listener until an error occurs.
	Serve(net.Listener) error
}

// Application abstracts all the functional components to be run by the server.
type Application struct {
	server   grpcServer
	listener net.Listener
	logger   *zap.Logger
	db       *gorm.DB
	services Services
}

// Serve serves the application services.
func (app Application) Serve() error {
	return app.server.Serve(app.listener)
}
