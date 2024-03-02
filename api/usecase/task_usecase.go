package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/minguu42/simoom/api/domain/auth"
	"github.com/minguu42/simoom/api/domain/model"
	"github.com/minguu42/simoom/api/domain/repository"
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

func (in CreateTaskInput) Create(g model.IDGenerator, userID string) model.Task {
	return model.Task{
		ID:        g.Generate(),
		UserID:    userID,
		ProjectID: in.ProjectID,
		Name:      in.Name,
		Priority:  in.Priority,
	}
}

func (uc Task) CreateTask(ctx context.Context, in CreateTaskInput) (TaskOutput, error) {
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

	t := in.Create(uc.idgen, auth.GetUserID(ctx))
	if err := uc.repo.CreateTask(ctx, t); err != nil {
		return TaskOutput{}, fmt.Errorf("failed to create task: %w", err)
	}
	return TaskOutput{Task: t}, nil
}

type ListTasksInput struct {
	Limit     uint
	Offset    uint
	ProjectID *string
	TagID     *string
}

func (uc Task) ListTasks(ctx context.Context, in ListTasksInput) (TasksOutput, error) {
	if in.ProjectID != nil {
		p, err := uc.repo.GetProjectByID(ctx, *in.ProjectID)
		if err != nil {
			if errors.Is(err, repository.ErrModelNotFound) {
				return TasksOutput{}, ErrProjectNotFound
			}
			return TasksOutput{}, fmt.Errorf("failed to get project: %w", err)
		}
		if auth.GetUserID(ctx) != p.UserID {
			return TasksOutput{}, ErrProjectNotFound
		}
	}
	if in.TagID != nil {
		t, err := uc.repo.GetTagByID(ctx, *in.TagID)
		if err != nil {
			if errors.Is(err, repository.ErrModelNotFound) {
				return TasksOutput{}, ErrTagNotFound
			}
			return TasksOutput{}, fmt.Errorf("failed to get tag: %w", err)
		}
		if auth.GetUserID(ctx) != t.UserID {
			return TasksOutput{}, ErrTagNotFound
		}
	}

	ts, err := uc.repo.ListTasks(ctx, in.Limit+1, in.Offset, in.ProjectID, in.TagID)
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

func (uc Task) UpdateTask(ctx context.Context, in UpdateTaskInput) (TaskOutput, error) {
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

func (uc Task) DeleteTask(ctx context.Context, in DeleteTaskInput) error {
	if err := uc.repo.Transaction(ctx, func(ctxWithTx context.Context) error {
		t, err := uc.repo.GetTaskByID(ctxWithTx, in.ID)
		if err != nil {
			if errors.Is(err, repository.ErrModelNotFound) {
				return ErrTaskNotFound
			}
			return fmt.Errorf("failed to get task: %w", err)
		}
		if auth.GetUserID(ctxWithTx) != t.UserID {
			return ErrTaskNotFound
		}

		if err := uc.repo.DeleteTask(ctxWithTx, in.ID); err != nil {
			return fmt.Errorf("failed to delete task: %w", err)
		}
		return nil
	}); err != nil {
		return fmt.Errorf("failed to run transaction: %w", err)
	}
	return nil
}
