package usecase

import (
	"context"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/minguu42/simoom/pkg/domain/idgen"
	"github.com/minguu42/simoom/pkg/domain/model"
	"github.com/minguu42/simoom/pkg/domain/repository"
)

type TaskUsecase struct {
	Repo repository.Repository
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

func (uc TaskUsecase) CreateTask(ctx context.Context, in CreateTaskInput) (TaskOutput, error) {
	p, err := uc.Repo.GetProjectByID(ctx, in.ProjectID)
	if err != nil {
		if errors.Is(err, repository.ErrModelNotFound) {
			return TaskOutput{}, ErrProjectNotFound
		}
		return TaskOutput{}, ErrUnkown
	}
	if userID != p.UserID {
		return TaskOutput{}, ErrProjectNotFound
	}

	now := time.Now()
	t := model.Task{
		ID:        idgen.Generate(),
		UserID:    userID,
		ProjectID: in.ProjectID,
		Title:     in.Title,
		Priority:  in.Priority,
		CreatedAt: now,
		UpdatedAt: now,
	}
	if err := uc.Repo.CreateTask(ctx, t); err != nil {
		return TaskOutput{}, ErrUnkown
	}
	return TaskOutput{Task: t}, nil
}

type ListTasksByProjectIDInput struct {
	ProjectID string
	Limit     uint
	Offset    uint
}

func (uc TaskUsecase) ListTasksByProjectID(ctx context.Context, in ListTasksByProjectIDInput) (TasksOutput, error) {
	p, err := uc.Repo.GetProjectByID(ctx, in.ProjectID)
	if err != nil {
		if errors.Is(err, repository.ErrModelNotFound) {
			return TasksOutput{}, ErrProjectNotFound
		}
		return TasksOutput{}, ErrUnkown
	}
	if userID != p.UserID {
		return TasksOutput{}, ErrProjectNotFound
	}

	ts, err := uc.Repo.ListTasksByProjectID(ctx, in.ProjectID, in.Limit, in.Offset)
	if err != nil {
		return TasksOutput{}, ErrUnkown
	}
	return TasksOutput{
		Tasks:   ts,
		HasNext: false,
	}, nil
}

type ListTasksByTagIDInput struct {
	TagID  string
	Limit  uint
	Offset uint
}

func (uc TaskUsecase) ListTasksByTagID(ctx context.Context, in ListTasksByTagIDInput) (TasksOutput, error) {
	t, err := uc.Repo.GetTagByID(ctx, in.TagID)
	if err != nil {
		if errors.Is(err, repository.ErrModelNotFound) {
			return TasksOutput{}, ErrTagNotFound
		}
		return TasksOutput{}, ErrUnkown
	}
	if userID != t.UserID {
		return TasksOutput{}, ErrTagNotFound
	}

	ts, err := uc.Repo.ListTasksByTagID(ctx, in.TagID, in.Limit, in.Offset)
	if err != nil {
		return TasksOutput{}, ErrUnkown
	}
	return TasksOutput{
		Tasks:   ts,
		HasNext: false,
	}, nil
}

type UpdateTaskInput struct {
	ID          string
	ProjectID   string
	Title       *string
	Content     *string
	Priority    *uint
	DueOn       *time.Time
	CompletedAt *time.Time
}

func (uc TaskUsecase) UpdateTask(ctx context.Context, in UpdateTaskInput) (TaskOutput, error) {
	t, err := uc.Repo.GetTaskByID(ctx, in.ID)
	if err != nil {
		if errors.Is(err, repository.ErrModelNotFound) {
			return TaskOutput{}, ErrTaskNotFound
		}
		return TaskOutput{}, ErrUnkown
	}
	if userID != t.UserID {
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
	if err := uc.Repo.UpdateTask(ctx, t); err != nil {
		return TaskOutput{}, ErrUnkown
	}
	return TaskOutput{Task: t}, nil
}

type DeleteTaskInput struct {
	ID string
}

func (uc TaskUsecase) DeleteTask(ctx context.Context, in DeleteTaskInput) error {
	t, err := uc.Repo.GetTaskByID(ctx, in.ID)
	if err != nil {
		if errors.Is(err, repository.ErrModelNotFound) {
			return ErrTaskNotFound
		}
		return ErrUnkown
	}
	if userID != t.UserID {
		return ErrTaskNotFound
	}

	if err := uc.Repo.DeleteTask(ctx, in.ID); err != nil {
		return ErrUnkown
	}
	return nil
}
