package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/minguu42/simoom/api/apperr"
	"github.com/minguu42/simoom/api/domain"
	"github.com/minguu42/simoom/api/domain/model"
)

type Task struct {
	repo  domain.Repository
	idgen domain.IDGenerator
}

func NewTask(repo domain.Repository, idgen domain.IDGenerator) Task {
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
	ProjectID model.ProjectID
	Name      string
	Priority  uint
}

func (in CreateTaskInput) Create(g domain.IDGenerator, userID model.UserID) model.Task {
	return model.Task{
		ID:        model.TaskID(g.Generate()),
		UserID:    userID,
		ProjectID: in.ProjectID,
		Name:      in.Name,
		Priority:  in.Priority,
	}
}

func (uc Task) CreateTask(ctx context.Context, in CreateTaskInput) (TaskOutput, error) {
	p, err := uc.repo.GetProjectByID(ctx, in.ProjectID)
	if err != nil {
		if errors.Is(err, domain.ErrModelNotFound) {
			return TaskOutput{}, apperr.ErrProjectNotFound(err)
		}
		return TaskOutput{}, fmt.Errorf("failed to get project: %w", err)
	}
	user := model.UserFromContext(ctx)
	if !user.HasProject(p) {
		return TaskOutput{}, apperr.ErrProjectNotFound(err)
	}

	t := in.Create(uc.idgen, user.ID)
	if err := uc.repo.CreateTask(ctx, t); err != nil {
		return TaskOutput{}, fmt.Errorf("failed to create task: %w", err)
	}
	return TaskOutput{Task: t}, nil
}

type ListTasksInput struct {
	Limit     uint
	Offset    uint
	ProjectID *model.ProjectID
	TagID     *model.TagID
}

func (uc Task) ListTasks(ctx context.Context, in ListTasksInput) (TasksOutput, error) {
	user := model.UserFromContext(ctx)
	if in.ProjectID != nil {
		p, err := uc.repo.GetProjectByID(ctx, *in.ProjectID)
		if err != nil {
			if errors.Is(err, domain.ErrModelNotFound) {
				return TasksOutput{}, apperr.ErrProjectNotFound(err)
			}
			return TasksOutput{}, fmt.Errorf("failed to get project: %w", err)
		}
		if !user.HasProject(p) {
			return TasksOutput{}, apperr.ErrProjectNotFound(err)
		}
	}
	if in.TagID != nil {
		t, err := uc.repo.GetTagByID(ctx, *in.TagID)
		if err != nil {
			if errors.Is(err, domain.ErrModelNotFound) {
				return TasksOutput{}, apperr.ErrTagNotFound(err)
			}
			return TasksOutput{}, fmt.Errorf("failed to get tag: %w", err)
		}
		if !user.HasTag(t) {
			return TasksOutput{}, apperr.ErrTagNotFound(err)
		}
	}

	ts, err := uc.repo.ListTasksByUserID(ctx, user.ID, in.Limit+1, in.Offset, in.ProjectID, in.TagID)
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
	ID          model.TaskID
	Name        *string
	Content     *string
	Priority    *uint
	DueOn       *time.Time
	CompletedAt *time.Time
	TagIDs      []model.TagID
}

func (uc Task) UpdateTask(ctx context.Context, in UpdateTaskInput) (TaskOutput, error) {
	t, err := uc.repo.GetTaskByID(ctx, in.ID)
	if err != nil {
		if errors.Is(err, domain.ErrModelNotFound) {
			return TaskOutput{}, apperr.ErrTaskNotFound(err)
		}
		return TaskOutput{}, fmt.Errorf("failed to get task: %w", err)
	}
	if !model.UserFromContext(ctx).HasTask(t) {
		return TaskOutput{}, apperr.ErrTaskNotFound(err)
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
	if err := uc.repo.Transaction(ctx, func(ctxWithTx context.Context) error {
		if err := uc.repo.UpdateTask(ctxWithTx, t); err != nil {
			return fmt.Errorf("failed to update task: %w", err)
		}
		if in.TagIDs != nil {
			if err := uc.repo.UpdateTaskTags(ctxWithTx, in.ID, in.TagIDs); err != nil {
				return fmt.Errorf("failed to attach tags: %w", err)
			}
		}
		return nil
	}); err != nil {
		return TaskOutput{}, fmt.Errorf("failed to run transaction: %w", err)
	}

	updatedTask, err := uc.repo.GetTaskByID(ctx, t.ID)
	if err != nil {
		return TaskOutput{}, fmt.Errorf("failed to get updated task: %w", err)
	}
	return TaskOutput{Task: updatedTask}, nil
}

type DeleteTaskInput struct {
	ID model.TaskID
}

func (uc Task) DeleteTask(ctx context.Context, in DeleteTaskInput) error {
	if err := uc.repo.Transaction(ctx, func(ctxWithTx context.Context) error {
		t, err := uc.repo.GetTaskByID(ctxWithTx, in.ID)
		if err != nil {
			if errors.Is(err, domain.ErrModelNotFound) {
				return apperr.ErrTaskNotFound(err)
			}
			return fmt.Errorf("failed to get task: %w", err)
		}
		if !model.UserFromContext(ctx).HasTask(t) {
			return apperr.ErrTaskNotFound(err)
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
