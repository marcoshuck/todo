package domain

import (
	"encoding/json"
	tasksv1 "github.com/marcoshuck/todo/api/tasks/v1"
	"github.com/marcoshuck/todo/pkg/serializer"
	"github.com/marcoshuck/todo/pkg/validator"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gopkg.in/yaml.v3"
	"gorm.io/gorm"
	"time"
)

var _ serializer.JsonSerializer = (*Task)(nil)
var _ serializer.YamlSerializer = (*Task)(nil)
var _ serializer.ApiSerializer[*tasksv1.Task] = (*Task)(nil)

// Task defines the scope of an action a User implements in their tasks dashboard.
type Task struct {
	gorm.Model
	Title       string `validate:"required,min=3"`
	Description string
	Deadline    *time.Time
	CompletedAt *time.Time
}

// API converts this Task to a tasksv1.Task.
func (t *Task) API() *tasksv1.Task {
	var deadline *timestamppb.Timestamp
	if t.Deadline != nil {
		deadline = timestamppb.New(*t.Deadline)
	}
	var completedAt *timestamppb.Timestamp
	if t.CompletedAt != nil {
		completedAt = timestamppb.New(*t.CompletedAt)
	}
	return &tasksv1.Task{
		Id:          int64(t.ID),
		Title:       t.Title,
		Description: t.Description,
		Deadline:    deadline,
		CompletedAt: completedAt,
		CreateTime:  timestamppb.New(t.CreatedAt),
		UpdateTime:  timestamppb.New(t.UpdatedAt),
	}
}

// JSON converts this Task to a slice of bytes in JSON format.
func (t *Task) JSON() ([]byte, error) {
	return json.Marshal(t)
}

// YAML converts this Task to a slice of bytes in YAML format.
func (t *Task) YAML() ([]byte, error) {
	return yaml.Marshal(t)
}

// FromCreateRequest fills out the information of the current Task from the given tasksv1.CreateTaskRequest
func (t *Task) FromCreateRequest(request *tasksv1.CreateTaskRequest) error {
	return t.fromAPI(request.GetTask()).validate()
}

// fromAPI fills out the current Task with the information from tasksv1.Task.
func (t *Task) fromAPI(in *tasksv1.Task) *Task {
	t.Title = in.GetTitle()
	t.Description = in.GetDescription()
	if d := in.GetDeadline(); d != nil {
		td := d.AsTime()
		t.Deadline = &td
	}
	return t
}

// validate validates the current Task
func (t *Task) validate() error {
	return validator.Validator.Struct(t)
}
