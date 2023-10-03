package server

import (
	"github.com/gojaguar/jaguar/database"
	tasksv1 "github.com/marcoshuck/todo/api/tasks/v1"
	"github.com/marcoshuck/todo/pkg/service"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"net"
)

func Run(app Application) error {
	app.logger.Debug("Listening...", zap.String("address", app.listener.Addr().String()))
	if err := app.Serve(); err != nil {
		return err
	}
	return nil
}

type Services struct {
	Tasks tasksv1.TasksServiceServer
}

type Application struct {
	server   *grpc.Server
	listener net.Listener
	logger   *zap.Logger
	db       *gorm.DB
	services Services
}

func (app Application) Serve() error {
	return app.server.Serve(app.listener)
}

func Setup(cfg Config) (Application, error) {
	logger, err := setupLogger(cfg)
	if err != nil {
		return Application{}, err
	}

	logger.Debug("Initializing server", zap.String("server.name", cfg.Name), zap.String("server.environment", cfg.Environment))

	db, err := setupDB(cfg, logger)
	if err != nil {
		return Application{}, err
	}

	l, err := setupListener(cfg, logger)
	if err != nil {
		return Application{}, err
	}

	svc := setupServices(db, logger)

	var opts []grpc.ServerOption

	srv := grpc.NewServer(opts...)

	return Application{
		server:   srv,
		listener: l,
		logger:   logger,
		db:       db,
		services: svc,
	}, nil
}

func setupServices(db *gorm.DB, logger *zap.Logger) Services {
	tasksService := service.NewTasks(db, logger)
	return Services{
		Tasks: tasksService,
	}
}

func setupListener(cfg Config, logger *zap.Logger) (net.Listener, error) {
	protocol, address := cfg.Listener()
	logger.Debug("Initializing listener", zap.String("listener.protocol", protocol), zap.String("listener.address", address))
	l, err := net.Listen(protocol, address)
	if err != nil {
		return nil, err
	}
	return l, nil
}

func setupLogger(cfg Config) (*zap.Logger, error) {
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

func setupDB(cfg Config, logger *zap.Logger) (*gorm.DB, error) {
	logger.Debug("Initializing DB connection", zap.String("db.engine", cfg.DB.Engine), zap.String("db.dsn", cfg.DB.ToDNS()))
	db, err := database.SetupConnectionSQL(cfg.DB)
	if err != nil {
		logger.Error("Failed to initialize DB connection", zap.Error(err))
		return nil, err
	}
	return db, nil
}
