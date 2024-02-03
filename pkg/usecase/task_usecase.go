package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/minguu42/simoom/pkg/domain/auth"
	"github.com/minguu42/simoom/pkg/domain/model"
	"github.com/minguu42/simoom/pkg/domain/repository"
)

type Task struct {
	repo  repository.Repository
	idgen model.IDGenerator
}

func NewTask(repo repository.Repository, idgen model.IDGenerator) Task {
	return Task{
		repo:  repo,
		idgen: idgen,
	}
}

type TaskOutput struct {
	Task model.Task
}

type TasksOutput struct {
	Tasks   []model.Task
	HasNext bool
}

type CreateTaskInput struct {
	ProjectID string
	Name      string
	Priority  uint
}

func (in CreateTaskInput) Validate() error {
	if len(in.ProjectID) != 26 {
		return newErrInvalidArgument("project_id is a 26-character string")
	}
	if len(in.Name) < 1 || 80 < len(in.Name) {
		return newErrInvalidArgument("name cannot be an empty string")
	}
	if in.Priority > 3 {
		return newErrInvalidArgument("priority is specified by 0 to 3")
	}
	return nil
}

func (uc Task) CreateTask(ctx context.Context, in CreateTaskInput) (TaskOutput, error) {
	if err := in.Validate(); err != nil {
		return TaskOutput{}, fmt.Errorf("failed to validate input: %w", err)
	}

	p, err := uc.repo.GetProjectByID(ctx, in.ProjectID)
	if err != nil {
		if errors.Is(err, repository.ErrModelNotFound) {
			return TaskOutput{}, ErrProjectNotFound
		}
		return TaskOutput{}, fmt.Errorf("failed to get project: %w", err)
	}
	if auth.GetUserID(ctx) != p.UserID {
		return TaskOutput{}, ErrProjectNotFound
	}

	t := model.Task{
		ID:        uc.idgen.Generate(),
		UserID:    auth.GetUserID(ctx),
		ProjectID: in.ProjectID,
		Name:      in.Name,
		Priority:  in.Priority,
	}
	if err := uc.repo.CreateTask(ctx, t); err != nil {
		return TaskOutput{}, fmt.Errorf("failed to create task: %w", err)
	}
	return TaskOutput{Task: t}, nil
}

type ListTasksByProjectIDInput struct {
	ProjectID string
	Limit     uint
	Offset    uint
}

func (in ListTasksByProjectIDInput) Validate() error {
	if len(in.ProjectID) != 26 {
		return newErrInvalidArgument("project_id is a 26-character string")
	}
	if in.Limit < 1 {
		return newErrInvalidArgument("limit is greater than or equal to 1")
	}
	return nil
}

func (uc Task) ListTasksByProjectID(ctx context.Context, in ListTasksByProjectIDInput) (TasksOutput, error) {
	if err := in.Validate(); err != nil {
		return TasksOutput{}, fmt.Errorf("failed to validate input: %w", err)
	}

	p, err := uc.repo.GetProjectByID(ctx, in.ProjectID)
	if err != nil {
		if errors.Is(err, repository.ErrModelNotFound) {
			return TasksOutput{}, ErrProjectNotFound
		}
		return TasksOutput{}, fmt.Errorf("failed to get project: %w", err)
	}
	if auth.GetUserID(ctx) != p.UserID {
		return TasksOutput{}, ErrProjectNotFound
	}

	ts, err := uc.repo.ListTasksByProjectID(ctx, in.ProjectID, in.Limit+1, in.Offset)
	if err != nil {
		return TasksOutput{}, fmt.Errorf("failed to list tasks: %w", err)
	}

	hasNext := false
	if len(ts) == int(in.Limit+1) {
		ts = ts[:in.Limit]
		hasNext = true
	}
	return TasksOutput{
		Tasks:   ts,
		HasNext: hasNext,
	}, nil
}

type ListTasksByTagIDInput struct {
	TagID  string
	Limit  uint
	Offset uint
}

func (in ListTasksByTagIDInput) Validate() error {
	if len(in.TagID) != 26 {
		return newErrInvalidArgument("tag_id is a 26-character string")
	}
	if in.Limit < 1 {
		return newErrInvalidArgument("limit is greater than or equal to 1")
	}
	return nil
}

func (uc Task) ListTasksByTagID(ctx context.Context, in ListTasksByTagIDInput) (TasksOutput, error) {
	if err := in.Validate(); err != nil {
		return TasksOutput{}, fmt.Errorf("failed to validate input: %w", err)
	}

	t, err := uc.repo.GetTagByID(ctx, in.TagID)
	if err != nil {
		if errors.Is(err, repository.ErrModelNotFound) {
			return TasksOutput{}, ErrTagNotFound
		}
		return TasksOutput{}, fmt.Errorf("failed to get tag: %w", err)
	}
	if auth.GetUserID(ctx) != t.UserID {
		return TasksOutput{}, ErrTagNotFound
	}

	ts, err := uc.repo.ListTasksByTagID(ctx, in.TagID, in.Limit+1, in.Offset)
	if err != nil {
		return TasksOutput{}, fmt.Errorf("failed to list tasks: %w", err)
	}

	hasNext := false
	if len(ts) == int(in.Limit+1) {
		ts = ts[:in.Limit]
		hasNext = true
	}
	return TasksOutput{
		Tasks:   ts,
		HasNext: hasNext,
	}, nil
}

type UpdateTaskInput struct {
	ID          string
	Name        *string
	Content     *string
	Priority    *uint
	DueOn       *time.Time
	CompletedAt *time.Time
}

func (in UpdateTaskInput) Validate() error {
	if len(in.ID) != 26 {
		return newErrInvalidArgument("id is a 26-character string")
	}
	if in.Name == nil && in.Content == nil && in.Priority == nil && in.DueOn == nil && in.CompletedAt == nil {
		return newErrInvalidArgument("must contain some argument other than id")
	}
	if in.Name != nil && (len(*in.Name) < 1 || 80 < len(*in.Name)) {
		return newErrInvalidArgument("name cannot be an empty string")
	}
	if in.Priority != nil && *in.Priority > 3 {
		return newErrInvalidArgument("priority is specified by 0 to 3")
	}
	return nil
}

func (uc Task) UpdateTask(ctx context.Context, in UpdateTaskInput) (TaskOutput, error) {
	if err := in.Validate(); err != nil {
		return TaskOutput{}, fmt.Errorf("failed to validate input: %w", err)
	}

	t, err := uc.repo.GetTaskByID(ctx, in.ID)
	if err != nil {
		if errors.Is(err, repository.ErrModelNotFound) {
			return TaskOutput{}, ErrTaskNotFound
		}
		return TaskOutput{}, fmt.Errorf("failed to get task: %w", err)
	}
	if auth.GetUserID(ctx) != t.UserID {
		return TaskOutput{}, ErrTaskNotFound
	}

	if in.Name != nil {
		t.Name = *in.Name
	}
	if in.Content != nil {
		t.Content = *in.Content
	}
	if in.Priority != nil {
		t.Priority = *in.Priority
	}
	if in.DueOn != nil {
		t.DueOn = in.DueOn
	}
	if in.CompletedAt != nil {
		t.CompletedAt = in.CompletedAt
	}
	if err := uc.repo.UpdateTask(ctx, t); err != nil {
		return TaskOutput{}, fmt.Errorf("failed to update task: %w", err)
	}
	return TaskOutput{Task: t}, nil
}

type DeleteTaskInput struct {
	ID string
}

func (in DeleteTaskInput) Validate() error {
	if len(in.ID) != 26 {
		return newErrInvalidArgument("id is a 26-character string")
	}
	return nil
}

func (uc Task) DeleteTask(ctx context.Context, in DeleteTaskInput) error {
	if err := in.Validate(); err != nil {
		return fmt.Errorf("failed to validate input: %w", err)
	}

	t, err := uc.repo.GetTaskByID(ctx, in.ID)
	if err != nil {
		if errors.Is(err, repository.ErrModelNotFound) {
			return ErrTaskNotFound
		}
		return fmt.Errorf("failed to get task: %w", err)
	}
	if auth.GetUserID(ctx) != t.UserID {
		return ErrTaskNotFound
	}

	if err := uc.repo.DeleteTask(ctx, in.ID); err != nil {
		return fmt.Errorf("failed to delete task: %w", err)
	}
	return nil
}
