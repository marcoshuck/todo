package domain

import (
	"encoding/json"
	tasksv1 "github.com/marcoshuck/todo/api/tasks/v1"
	"github.com/marcoshuck/todo/pkg/serializer"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gopkg.in/yaml.v3"
	"gorm.io/gorm"
	"time"
)

var _ serializer.JSON = (*Task)(nil)
var _ serializer.YAML = (*Task)(nil)
var _ serializer.API[*tasksv1.Task] = (*Task)(nil)

// Task defines the scope of an action a User implements in their tasks dashboard.
type Task struct {
	gorm.Model
	Title       string     `json:"title" validate:"required,min=3"`
	Description string     `json:"description"`
	Deadline    *time.Time `json:"deadline"`
	CompletedAt *time.Time `json:"completed_at"`
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

// FromAPI fills this Task with the data from tasksv1.Task.
func (t *Task) FromAPI(in *tasksv1.Task) {
	t.ID = uint(in.GetId())
	t.Title = in.GetTitle()
	t.Description = in.GetDescription()
	if in.GetDeadline() != nil {
		d := in.GetDeadline().AsTime()
		t.Deadline = &d
	}
	if in.GetCompletedAt() != nil {
		at := in.GetDeadline().AsTime()
		t.CompletedAt = &at
	}
	t.CreatedAt = in.GetCreateTime().AsTime()
	t.UpdatedAt = in.GetUpdateTime().AsTime()
}

// JSON converts this Task to a slice of bytes in JSON format.
func (t *Task) JSON() ([]byte, error) {
	return json.Marshal(t)
}

// FromJSON converts a slice of bytes in JSON format to a Task.
func (t *Task) FromJSON(data []byte) error {
	return json.Unmarshal(data, t)
}

// YAML converts this Task to a slice of bytes in YAML format.
func (t *Task) YAML() ([]byte, error) {
	return yaml.Marshal(t)
}

// FromYAML converts a slice of bytes in YAML format to a Task.
func (t *Task) FromYAML(data []byte) error {
	return yaml.Unmarshal(data, t)
}

// ApplyMask returns the Map of the current Task with the given mask applied.
func (t *Task) ApplyMask(mask *fieldmaskpb.FieldMask) map[string]any {
	m := t.Map()
	mask.Normalize()
	for _, p := range mask.GetPaths() {
		_, ok := m[p]
		if !ok {
			delete(m, p)
		}
	}
	return m
}

// Map converts the current Task to a map.
func (t *Task) Map() map[string]any {
	b, err := t.JSON()
	if err != nil {
		return nil
	}
	m := make(map[string]any)
	err = json.Unmarshal(b, &m)
	if err != nil {
		return nil
	}
	return m
}
