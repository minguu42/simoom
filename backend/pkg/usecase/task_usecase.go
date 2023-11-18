package usecase

import (
	"context"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/minguu42/simoom/backend/pkg/domain/auth"
	"github.com/minguu42/simoom/backend/pkg/domain/idgen"
	"github.com/minguu42/simoom/backend/pkg/domain/model"
	"github.com/minguu42/simoom/backend/pkg/domain/repository"
)

type Task struct {
	repo repository.Repository
}

func NewTask(repo repository.Repository) Task {
	return Task{repo: repo}
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
	Title     string
	Priority  uint
}

func (uc Task) CreateTask(ctx context.Context, in CreateTaskInput) (TaskOutput, error) {
	p, err := uc.repo.GetProjectByID(ctx, in.ProjectID)
	if err != nil {
		if errors.Is(err, repository.ErrModelNotFound) {
			return TaskOutput{}, ErrProjectNotFound
		}
		return TaskOutput{}, errors.WithStack(err)
	}
	if auth.GetUserID(ctx) != p.UserID {
		return TaskOutput{}, ErrProjectNotFound
	}

	now := time.Now()
	t := model.Task{
		ID:        idgen.Generate(),
		UserID:    auth.GetUserID(ctx),
		ProjectID: in.ProjectID,
		Title:     in.Title,
		Priority:  in.Priority,
		CreatedAt: now,
		UpdatedAt: now,
	}
	if err := uc.repo.CreateTask(ctx, t); err != nil {
		return TaskOutput{}, errors.WithStack(err)
	}
	return TaskOutput{Task: t}, nil
}

type ListTasksByProjectIDInput struct {
	ProjectID string
	Limit     uint
	Offset    uint
}

func (uc Task) ListTasksByProjectID(ctx context.Context, in ListTasksByProjectIDInput) (TasksOutput, error) {
	p, err := uc.repo.GetProjectByID(ctx, in.ProjectID)
	if err != nil {
		if errors.Is(err, repository.ErrModelNotFound) {
			return TasksOutput{}, ErrProjectNotFound
		}
		return TasksOutput{}, errors.WithStack(err)
	}
	if auth.GetUserID(ctx) != p.UserID {
		return TasksOutput{}, ErrProjectNotFound
	}

	ts, err := uc.repo.ListTasksByProjectID(ctx, in.ProjectID, in.Limit+1, in.Offset)
	if err != nil {
		return TasksOutput{}, errors.WithStack(err)
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

func (uc Task) ListTasksByTagID(ctx context.Context, in ListTasksByTagIDInput) (TasksOutput, error) {
	t, err := uc.repo.GetTagByID(ctx, in.TagID)
	if err != nil {
		if errors.Is(err, repository.ErrModelNotFound) {
			return TasksOutput{}, ErrTagNotFound
		}
		return TasksOutput{}, errors.WithStack(err)
	}
	if auth.GetUserID(ctx) != t.UserID {
		return TasksOutput{}, ErrTagNotFound
	}

	ts, err := uc.repo.ListTasksByTagID(ctx, in.TagID, in.Limit+1, in.Offset)
	if err != nil {
		return TasksOutput{}, errors.WithStack(err)
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
	Title       *string
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
		return TaskOutput{}, errors.WithStack(err)
	}
	if auth.GetUserID(ctx) != t.UserID {
		return TaskOutput{}, ErrTaskNotFound
	}

	if in.Title != nil {
		t.Title = *in.Title
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
		return TaskOutput{}, errors.WithStack(err)
	}
	return TaskOutput{Task: t}, nil
}

type DeleteTaskInput struct {
	ID string
}

func (uc Task) DeleteTask(ctx context.Context, in DeleteTaskInput) error {
	t, err := uc.repo.GetTaskByID(ctx, in.ID)
	if err != nil {
		if errors.Is(err, repository.ErrModelNotFound) {
			return ErrTaskNotFound
		}
		return errors.WithStack(err)
	}
	if auth.GetUserID(ctx) != t.UserID {
		return ErrTaskNotFound
	}

	if err := uc.repo.DeleteTask(ctx, in.ID); err != nil {
		return errors.WithStack(err)
	}
	return nil
}