package domain

import (
	"errors"
	"github.com/gojaguar/jaguar/strings"
	tasksv1 "github.com/marcoshuck/todo/api/tasks/v1"
	"github.com/marcoshuck/todo/internal/serializer"
	fieldmask_utils "github.com/mennanov/fieldmask-utils"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
	"time"
)

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

// ApplyMask returns the Map of the current Task with the given mask applied.
func (t *Task) ApplyMask(mask *fieldmaskpb.FieldMask) (map[string]any, error) {
	mask.Normalize()
	if !mask.IsValid(t.API()) {
		return nil, errors.New("invalid mask")
	}
	protoMask, err := fieldmask_utils.MaskFromProtoFieldMask(mask, strings.PascalCase)
	if err != nil {
		return nil, err
	}
	m := make(map[string]any)
	if err = fieldmask_utils.StructToMap(protoMask, t.API(), m); err != nil {
		return nil, err
	}
	return m, nil
}
