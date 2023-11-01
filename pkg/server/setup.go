package server

import (
	"github.com/gojaguar/jaguar/database"
	"github.com/marcoshuck/todo/pkg/service"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/noop"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"net"
)

// Setup creates a new application using the given Config.
func Setup(cfg Config) (Application, error) {
	logger, err := setupLogger(cfg)
	if err != nil {
		return Application{}, err
	}

	tracerProvider, err := setupTracerProvider(cfg)
	if err != nil {
		return Application{}, err
	}

	meterProvider, err := setupMeterProvider(cfg)
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

	svc := setupServices(db, logger, tracerProvider, meterProvider)

	var opts []grpc.ServerOption

	srv := grpc.NewServer(opts...)

	return Application{
		server:         srv,
		listener:       l,
		logger:         logger,
		tracerProvider: tracerProvider,
		meterProvider:  meterProvider,
		db:             db,
		services:       svc,
	}, nil
}

// setupServices initializes the Application Services.
func setupServices(db *gorm.DB, logger *zap.Logger, tracerProvider trace.TracerProvider, meterProvider metric.MeterProvider) Services {
	logger.Debug("Initializing services")
	tasksService := service.NewTasks(db, logger, tracerProvider.Tracer("todo.huck.com.ar/tasks"))
	return Services{
		Tasks: tasksService,
	}
}

// setupListener initializes a new tcp listener used by a gRPC server.
func setupListener(cfg Config, logger *zap.Logger) (net.Listener, error) {
	protocol, address := cfg.Listener()
	logger.Debug("Initializing listener", zap.String("listener.protocol", protocol), zap.String("listener.address", address))
	l, err := net.Listen(protocol, address)
	if err != nil {
		logger.Error("Failed to initialize listener", zap.Error(err))
		return nil, err
	}
	return l, nil
}

// setupDB initializes a new connection with a DB server.
func setupDB(cfg Config, logger *zap.Logger) (*gorm.DB, error) {
	logger.Debug("Initializing DB connection", zap.String("db.engine", cfg.DB.Engine), zap.String("db.dsn", cfg.DB.DSN()))
	db, err := database.SetupConnectionSQL(cfg.DB)
	if err != nil {
		logger.Error("Failed to initialize DB connection", zap.Error(err))
		return nil, err
	}
	return db, nil
}

// setupLogger initializes a new Zap Logger with the parameters specified by the given Config.
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

func setupTracerProvider(cfg Config) (trace.TracerProvider, error) {
	var tracerProvider trace.TracerProvider
	switch cfg.Environment {
	default:
		tracerProvider = trace.NewNoopTracerProvider()
	}
	return tracerProvider, nil
}

func setupMeterProvider(cfg Config) (metric.MeterProvider, error) {
	var meterProvider metric.MeterProvider
	switch cfg.Environment {
	default:
		meterProvider = noop.NewMeterProvider()
	}
	return meterProvider, nil
}
