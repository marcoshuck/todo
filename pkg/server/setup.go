package server

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gojaguar/jaguar/database"
	grpc_logging "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"github.com/marcoshuck/todo/pkg/service"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/noop"
	metric_sdk "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	trace_sdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"io"
	"net"
)

// Setup creates a new application using the given Config.
func Setup(cfg Config) (Application, error) {
	logger, err := setupLogger(cfg)
	if err != nil {
		return Application{}, err
	}

	tracerProvider, traceExporter, err := setupTracerProvider(cfg)
	if err != nil {
		return Application{}, err
	}

	meterProvider, meterExporter, err := setupMeterProvider(cfg)
	if err != nil {
		return Application{}, err
	}

	logger.Debug("Initializing server", zap.String("server.name", cfg.Name), zap.String("server.environment", cfg.Environment))

	db, dbConn, err := setupDB(cfg, logger, tracerProvider)
	if err != nil {
		return Application{}, err
	}

	l, err := setupListener(cfg, logger)
	if err != nil {
		return Application{}, err
	}

	svc := setupServices(db, logger, tracerProvider, meterProvider)

	opts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			otelgrpc.UnaryServerInterceptor(
				otelgrpc.WithTracerProvider(tracerProvider),
				otelgrpc.WithMeterProvider(meterProvider),
			),
			grpc_logging.UnaryServerInterceptor(InterceptorLogger(logger)),
			grpc_recovery.UnaryServerInterceptor(grpc_recovery.WithRecoveryHandler(grpcPanicRecoveryHandler)),
		),
		grpc.ChainStreamInterceptor(
			otelgrpc.StreamServerInterceptor(
				otelgrpc.WithTracerProvider(tracerProvider),
				otelgrpc.WithMeterProvider(meterProvider),
			),
			grpc_logging.StreamServerInterceptor(InterceptorLogger(logger)),
			grpc_recovery.StreamServerInterceptor(grpc_recovery.WithRecoveryHandler(grpcPanicRecoveryHandler)),
		),
	}

	srv := grpc.NewServer(opts...)

	return Application{
		server:         srv,
		listener:       l,
		logger:         logger,
		tracerProvider: tracerProvider,
		meterProvider:  meterProvider,
		db:             db,
		services:       svc,
		shutdown: []shutDowner{
			traceExporter,
			meterExporter,
		},
		closer: []io.Closer{
			dbConn,
		},
	}, nil
}

// setupServices initializes the Application Services.
func setupServices(db *gorm.DB, logger *zap.Logger, tracerProvider trace.TracerProvider, meterProvider metric.MeterProvider) Services {
	logger.Debug("Initializing services")
	tasksService := service.NewTasks(db, logger, meterProvider.Meter("todo.huck.com.ar/tasks"))
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
func setupDB(cfg Config, logger *zap.Logger, provider trace.TracerProvider) (*gorm.DB, *sql.DB, error) {
	logger.Debug("Initializing DB connection", zap.String("db.engine", cfg.DB.Engine), zap.String("db.dsn", cfg.DB.DSN()))
	db, err := database.SetupConnectionSQL(cfg.DB)
	if err != nil {
		logger.Error("Failed to initialize DB connection", zap.Error(err))
		return nil, nil, err
	}
	err = db.Use(otelgorm.NewPlugin(
		otelgorm.WithDBName(cfg.DB.Name),
		otelgorm.WithAttributes(attribute.String("db.engine", cfg.DB.Engine)),
		otelgorm.WithTracerProvider(provider),
	))
	if err != nil {
		logger.Error("Failed to set up DB plugin", zap.Error(err))
		return nil, nil, err
	}
	return db, nil, nil
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

func setupTracerProvider(cfg Config) (trace.TracerProvider, trace_sdk.SpanExporter, error) {
	var tracerProvider trace.TracerProvider
	var traceExporter trace_sdk.SpanExporter
	switch cfg.Environment {
	case "production", "staging":
		ctx := context.Background()
		res, err := newResource(ctx, cfg)
		if err != nil {
			return nil, nil, err
		}
		conn, err := grpc.DialContext(ctx, cfg.Tracing.Address(),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to create gRPC connection to collector: %w", err)
		}
		traceExporter, err = otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn))
		if err != nil {
			return nil, nil, fmt.Errorf("failed to create trace exporter: %w", err)
		}
		bsp := trace_sdk.NewBatchSpanProcessor(traceExporter)
		tracerProvider = trace_sdk.NewTracerProvider(
			trace_sdk.WithSampler(trace_sdk.AlwaysSample()),
			trace_sdk.WithResource(res),
			trace_sdk.WithSpanProcessor(bsp),
		)
	default:
		tracerProvider = trace.NewNoopTracerProvider()
	}

	return tracerProvider, traceExporter, nil
}

func setupMeterProvider(cfg Config) (metric.MeterProvider, metric_sdk.Exporter, error) {
	var meterProvider metric.MeterProvider
	var meterExporter metric_sdk.Exporter
	switch cfg.Environment {
	case "production", "staging":
		ctx := context.Background()
		res, err := newResource(ctx, cfg)
		if err != nil {
			return nil, nil, err
		}
		conn, err := grpc.DialContext(ctx, cfg.Metrics.Address(),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to create gRPC connection to collector: %w", err)
		}
		meterExporter, err = otlpmetricgrpc.New(ctx, otlpmetricgrpc.WithGRPCConn(conn))
		if err != nil {
			return nil, nil, fmt.Errorf("failed to create metrics exporter: %w", err)
		}
		meterProvider = metric_sdk.NewMeterProvider(
			metric_sdk.WithReader(metric_sdk.NewPeriodicReader(meterExporter)),
			metric_sdk.WithResource(res),
		)
	default:
		meterProvider = noop.NewMeterProvider()
	}
	return meterProvider, meterExporter, nil
}

func newResource(ctx context.Context, cfg Config) (*resource.Resource, error) {
	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceName(cfg.Name),
			semconv.DeploymentEnvironment(cfg.Environment),
		),
	)
	return res, err
}

func InterceptorLogger(l *zap.Logger) grpc_logging.Logger {
	return grpc_logging.LoggerFunc(func(ctx context.Context, lvl grpc_logging.Level, msg string, fields ...any) {
		f := make([]zap.Field, 0, len(fields)/2)

		for i := 0; i < len(fields); i += 2 {
			key := fields[i]
			value := fields[i+1]

			switch v := value.(type) {
			case string:
				f = append(f, zap.String(key.(string), v))
			case int:
				f = append(f, zap.Int(key.(string), v))
			case bool:
				f = append(f, zap.Bool(key.(string), v))
			default:
				f = append(f, zap.Any(key.(string), v))
			}
		}

		logger := l.WithOptions(zap.AddCallerSkip(1)).With(f...)

		switch lvl {
		case grpc_logging.LevelDebug:
			logger.Debug(msg)
		case grpc_logging.LevelInfo:
			logger.Info(msg)
		case grpc_logging.LevelWarn:
			logger.Warn(msg)
		case grpc_logging.LevelError:
			logger.Error(msg)
		default:
			panic(fmt.Sprintf("unknown level %v", lvl))
		}
	})
}

func grpcPanicRecoveryHandler(p any) error {
	return status.Errorf(codes.Unknown, "panic triggered: %v", p)
}
