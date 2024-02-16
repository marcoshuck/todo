package service

import (
	"context"
	"errors"
	tasksv1 "github.com/marcoshuck/todo/api/tasks/v1"
	"github.com/marcoshuck/todo/internal/domain"
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
		return nil, status.Error(codes.Unavailable, "failed to create task")
	}
	svc.logger.Debug("Returning created task", zap.String("task.title", request.GetTask().GetTitle()))
	return task.API(), nil
}

// DeleteTask deletes a task.
func (svc *tasks) DeleteTask(ctx context.Context, request *tasksv1.DeleteTaskRequest) (*tasksv1.Task, error) {
	svc.logger.Debug("Deleting task", zap.Int64("task.id", request.GetId()))
	span := trace.SpanFromContext(ctx)
	defer span.End()
	task, err := svc.GetTask(ctx, &tasksv1.GetTaskRequest{Id: request.GetId()})
	if err != nil {
		return nil, err
	}
	span.AddEvent("Deleting task from the database")
	err = svc.db.Model(&domain.Task{}).Delete(&domain.Task{}, task.GetId()).Error
	if err != nil {
		svc.logger.Error("Failed to delete task", zap.Error(err))
		span.RecordError(err)
		return nil, status.Error(codes.Unavailable, "failed to delete task")
	}
	svc.logger.Debug("Task was deleted", zap.Int64("task.id", request.GetId()))
	return task, nil
}

// UndeleteTask undeletes a task. Opposite operation to DeleteTask.
func (svc *tasks) UndeleteTask(ctx context.Context, request *tasksv1.UndeleteTaskRequest) (*tasksv1.Task, error) {
	svc.logger.Debug("Undeleting task", zap.Int64("task.id", request.GetId()))
	span := trace.SpanFromContext(ctx)
	defer span.End()
	err := svc.db.Model(&domain.Task{}).Unscoped().Where("id = ?", request.GetId()).Update("deleted_at", nil).Error
	if err != nil {
		svc.logger.Error("Failed to undelete task", zap.Error(err))
		span.RecordError(err)
		return nil, status.Error(codes.Unavailable, "failed to undelete task")
	}
	task, err := svc.GetTask(ctx, &tasksv1.GetTaskRequest{Id: request.GetId()})
	if err != nil {
		return nil, err
	}
	svc.logger.Debug("Task was undeleted", zap.Int64("task.id", request.GetId()))
	return task, nil
}

// UpdateTask updates a task.
func (svc *tasks) UpdateTask(ctx context.Context, request *tasksv1.UpdateTaskRequest) (*tasksv1.Task, error) {
	svc.logger.Debug("Updating task", zap.Int64("task.id", request.GetTask().GetId()))

	span := trace.SpanFromContext(ctx)
	defer span.End()

	_, err := svc.GetTask(ctx, &tasksv1.GetTaskRequest{Id: request.GetTask().GetId()})
	if err != nil {
		return nil, err
	}

	var task domain.Task
	task.FromAPI(request.GetTask())
	m, err := task.ApplyMask(request.GetUpdateMask())
	if err != nil {
		svc.logger.Error("Failed to apply update mask", zap.Error(err))
		span.RecordError(err)
		return nil, status.Error(codes.Internal, "failed to generate update mask")
	}
	err = svc.db.Model(&domain.Task{}).Where("id = ?", request.GetTask().GetId()).Updates(m).Error
	if err != nil {
		svc.logger.Error("Failed to update task", zap.Error(err))
		span.RecordError(err)
		return nil, status.Error(codes.Internal, "failed to update task")
	}
	svc.logger.Debug("Task was updated", zap.Int64("task.id", request.GetTask().GetId()))
	return svc.GetTask(ctx, &tasksv1.GetTaskRequest{Id: request.GetTask().GetId()})
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
