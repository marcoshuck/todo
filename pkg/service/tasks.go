package service

import (
	"context"
	"errors"
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

// GetTask gets a task by ID.
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
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, "task not found")
		}
		return nil, status.Errorf(codes.Unavailable, "failed to get task: %v", err)
	}
	svc.logger.Debug("Returning task", zap.Int64("task.id", request.GetId()), zap.String("task.title", task.Title))
	return task.API(), nil
}

// ListTasks lists tasks.
func (svc *tasks) ListTasks(ctx context.Context, request *tasksv1.ListTasksRequest) (*tasksv1.ListTasksResponse, error) {
	svc.logger.Debug("Getting task list", zap.Int32("page_size", request.GetPageSize()), zap.String("page_token", request.GetPageToken()))
	span := trace.SpanFromContext(ctx)
	defer span.End()
	span.AddEvent("Getting tasks from the database")
	var out []domain.Task
	err := svc.db.Model(&domain.Task{}).WithContext(ctx).Find(&out).Error
	if err != nil {
		svc.logger.Error("Failed to list tasks", zap.Error(err))
		span.RecordError(err)
		return nil, status.Errorf(codes.Unavailable, "failed to get task: %v", err)
	}
	svc.logger.Debug("Returning task list", zap.Int32("page_size", request.GetPageSize()), zap.String("page_token", request.GetPageToken()), zap.Int("count", len(out)))
	res := tasksv1.ListTasksResponse{
		Tasks:         make([]*tasksv1.Task, len(out)),
		NextPageToken: "",
	}
	for i, task := range out {
		res.Tasks[i] = task.API()
	}
	return &res, nil
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
