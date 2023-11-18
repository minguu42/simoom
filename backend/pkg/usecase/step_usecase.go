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

type Step struct {
	repo repository.Repository
}

func NewStep(repo repository.Repository) Step {
	return Step{repo: repo}
}

type StepOutput struct {
	Step model.Step
}

type StepsOutput struct {
	Steps   []model.Step
	HasNext bool
}

type CreateStepInput struct {
	TaskID string
	Title  string
}

func (uc Step) CreateStep(ctx context.Context, in CreateStepInput) (StepOutput, error) {
	t, err := uc.repo.GetTaskByID(ctx, in.TaskID)
	if err != nil {
		if errors.Is(err, repository.ErrModelNotFound) {
			return StepOutput{}, ErrTaskNotFound
		}
		return StepOutput{}, errors.WithStack(err)
	}
	if auth.GetUserID(ctx) != t.UserID {
		return StepOutput{}, ErrTaskNotFound
	}

	now := time.Now()
	s := model.Step{
		ID:        idgen.Generate(),
		UserID:    auth.GetUserID(ctx),
		TaskID:    in.TaskID,
		Title:     in.Title,
		CreatedAt: now,
		UpdatedAt: now,
	}
	if err := uc.repo.CreateStep(ctx, s); err != nil {
		return StepOutput{}, errors.WithStack(err)
	}
	return StepOutput{Step: s}, nil
}

type UpdateStepInput struct {
	ID          string
	Title       *string
	CompletedAt *time.Time
}

func (uc Step) UpdateStep(ctx context.Context, in UpdateStepInput) (StepOutput, error) {
	s, err := uc.repo.GetStepByID(ctx, in.ID)
	if err != nil {
		if errors.Is(err, repository.ErrModelNotFound) {
			return StepOutput{}, ErrStepNotFound
		}
		return StepOutput{}, errors.WithStack(err)
	}
	if auth.GetUserID(ctx) != s.UserID {
		return StepOutput{}, ErrStepNotFound
	}

	if in.Title != nil {
		s.Title = *in.Title
	}
	if in.CompletedAt != nil {
		s.CompletedAt = in.CompletedAt
	}
	if err := uc.repo.UpdateStep(ctx, s); err != nil {
		return StepOutput{}, errors.WithStack(err)
	}
	return StepOutput{Step: s}, nil
}

type DeleteStepInput struct {
	ID string
}

func (uc Step) DeleteStep(ctx context.Context, in DeleteStepInput) error {
	s, err := uc.repo.GetStepByID(ctx, in.ID)
	if err != nil {
		if errors.Is(err, repository.ErrModelNotFound) {
			return ErrStepNotFound
		}
		return errors.WithStack(err)
	}
	if auth.GetUserID(ctx) != s.UserID {
		return ErrStepNotFound
	}

	if err := uc.repo.DeleteStep(ctx, in.ID); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
