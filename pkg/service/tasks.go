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

// tasks implements tasksv1.TasksServiceServer.
type tasks struct {
	tasksv1.UnimplementedTasksServiceServer
	db     *gorm.DB
	logger *zap.Logger
	meter  metric.Meter
}

// CreateTask creates a Task.
func (svc *tasks) CreateTask(ctx context.Context, request *tasksv1.CreateTaskRequest) (*tasksv1.Task, error) {
	svc.logger.Debug("Creating task", zap.String("task.title", request.GetTask().GetTitle()))
	span := trace.SpanFromContext(ctx)
	defer span.End()

	var task domain.Task
	svc.logger.Debug("Filling out task information")
	span.AddEvent("Parsing task from API request")
	err := task.FromCreateRequest(request)
	if err != nil {
		svc.logger.Warn("User submitted an invalid task", zap.Error(err))
		span.RecordError(err)
		return nil, status.Errorf(codes.InvalidArgument, "invalid task: %s", err)
	}
	span.AddEvent("Persisting task in the database")
	svc.logger.Debug("Persisting task in the database", zap.String("task.title", request.GetTask().GetTitle()))
	err = svc.db.Model(&domain.Task{}).WithContext(ctx).Create(&task).Error
	if err != nil {
		svc.logger.Error("Failed to create task", zap.Error(err))
		span.RecordError(err)
		return nil, status.Errorf(codes.Unavailable, "failed to create task")
	}
	svc.logger.Debug("Returning created task", zap.String("task.title", request.GetTask().GetTitle()))
	return task.API(), nil
}

// NewTasks initializes a new tasksv1.TasksServiceServer implementation.
func NewTasks(db *gorm.DB, logger *zap.Logger, meter metric.Meter) tasksv1.TasksServiceServer {
	tasksLogger := logger.Named("service.tasks")
	return &tasks{
		db:     db,
		logger: tasksLogger,
		meter:  meter,
	}
}
