package service

import (
	"context"
	tasksv1 "github.com/marcoshuck/todo/api/tasks/v1"
	"github.com/marcoshuck/todo/pkg/domain"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

// tasks implements tasksv1.TasksWriterServiceServer.
type tasks struct {
	tasksv1.UnimplementedTasksWriterServiceServer
	tasksv1.UnimplementedTasksReaderServiceServer
	db     *gorm.DB
	logger *zap.Logger
	meter  metric.Meter
}

func (svc *tasks) GetTask(ctx context.Context, request *tasksv1.GetTaskRequest) (*tasksv1.Task, error) {
	svc.logger.Debug("Getting task by ID", zap.Int64("task.id", request.GetId()))
	span := trace.SpanFromContext(ctx)
	defer span.End()

	var task domain.Task
	span.AddEvent("Getting task from the database")
	err := svc.db.Model(&domain.Task{}).WithContext(ctx).First(&task, request.GetId()).Error
	if err != nil {
		svc.logger.Error("Failed to get task", zap.Error(err))
		span.RecordError(err)
		return nil, status.Errorf(codes.Unavailable, "failed to create task")
	}
	svc.logger.Debug("Returning task", zap.Int64("task.id", request.GetId()), zap.String("task.title", task.Title))
	return task.API(), nil
}

// CreateTask creates a Task.
func (svc *tasks) CreateTask(ctx context.Context, request *tasksv1.CreateTaskRequest) (*tasksv1.Task, error) {
	svc.logger.Debug("Creating task", zap.String("task.title", request.GetTask().GetTitle()))
	span := trace.SpanFromContext(ctx)
	defer span.End()

	var task domain.Task
	svc.logger.Debug("Filling out task information")
	span.AddEvent("Parsing task from API request")
	task.FromAPI(request.GetTask())
	span.AddEvent("Persisting task in the database")
	svc.logger.Debug("Persisting task in the database", zap.String("task.title", request.GetTask().GetTitle()))
	err := svc.db.Model(&domain.Task{}).WithContext(ctx).Create(&task).Error
	if err != nil {
		svc.logger.Error("Failed to create task", zap.Error(err))
		span.RecordError(err)
		return nil, status.Errorf(codes.Unavailable, "failed to create task")
	}
	svc.logger.Debug("Returning created task", zap.String("task.title", request.GetTask().GetTitle()))
	return task.API(), nil
}

// NewTasksWriter initializes a new tasksv1.TasksWriterServiceServer implementation.
func NewTasksWriter(db *gorm.DB, logger *zap.Logger, meter metric.Meter) tasksv1.TasksWriterServiceServer {
	tasksLogger := logger.Named("service.tasks.writer")
	return &tasks{
		db:     db,
		logger: tasksLogger,
		meter:  meter,
	}
}

// NewTasksReader initializes a new tasksv1.TasksWriterServiceServer implementation.
func NewTasksReader(db *gorm.DB, logger *zap.Logger, meter metric.Meter) tasksv1.TasksReaderServiceServer {
	tasksLogger := logger.Named("service.tasks.reader")
	return &tasks{
		db:     db,
		logger: tasksLogger,
		meter:  meter,
	}
}
