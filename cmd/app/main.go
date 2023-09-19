package main

import (
	"fmt"
	"github.com/caarlos0/env/v9"
	"github.com/gojaguar/jaguar/database"
	tasksv1 "github.com/marcoshuck/todo/api/tasks/v1"
	"github.com/marcoshuck/todo/pkg/conf"
	"github.com/marcoshuck/todo/pkg/service"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	var cfg conf.Config
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatalln("Failed to read config:", err)
	}

	logger, err := setupLogger(cfg)
	if err != nil {
		log.Fatalln("Failed to initialize logger:", err)
	}

	logger.Debug("Initializing server:", zap.String("server.name", cfg.Name), zap.String("server.environment", cfg.Environment))

	logger.Debug("Initializing DB connection", zap.String("db.engine", cfg.DB.Engine), zap.String("db.dsn", cfg.DB.ToDNS()))
	db, err := database.SetupConnectionSQL(cfg.DB)
	if err != nil {
		logger.Fatal("Failed to initialize DB connection", zap.Error(err))
	}

	logger.Debug("Listening for incoming TCP connections", zap.Int("port", cfg.Port))
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", cfg.Port))
	if err != nil {
		logger.Fatal("Failed to listen for TCP connections", zap.Error(err))
	}

	server := grpc.NewServer()
	tasksService := service.NewTasks(db, logger)
	tasksv1.RegisterTasksServiceServer(server, tasksService)

	logger.Debug("Listening...", zap.Int("grpc.port", cfg.Port))
	if err := server.Serve(listener); err != nil {
		logger.Fatal("Failed to listen for incoming gRPC requests:", zap.Error(err))
	}
}

func setupLogger(cfg conf.Config) (*zap.Logger, error) {
	var logger *zap.Logger
	var err error
	switch cfg.Environment {
	case "production":
		logger, err = zap.NewProduction()
	case "staging":
		logger, err = zap.NewDevelopment()
	default:
		logger = zap.NewNop()
	}
	if err != nil {
		return nil, err
	}
	logger = logger.Named(cfg.Name)
	return logger, nil
}
