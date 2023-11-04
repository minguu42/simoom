package usecase

import (
	"context"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/minguu42/simoom/pkg/domain/auth"
	"github.com/minguu42/simoom/pkg/domain/idgen"
	"github.com/minguu42/simoom/pkg/domain/model"
	"github.com/minguu42/simoom/pkg/domain/repository"
)

type StepUsecase struct {
	Repo repository.Repository
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

func (uc StepUsecase) CreateStep(ctx context.Context, in CreateStepInput) (StepOutput, error) {
	t, err := uc.Repo.GetTaskByID(ctx, in.TaskID)
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
	if err := uc.Repo.CreateStep(ctx, s); err != nil {
		return StepOutput{}, errors.WithStack(err)
	}
	return StepOutput{Step: s}, nil
}

type UpdateStepInput struct {
	ID          string
	Title       *string
	CompletedAt *time.Time
}

func (uc StepUsecase) UpdateStep(ctx context.Context, in UpdateStepInput) (StepOutput, error) {
	s, err := uc.Repo.GetStepByID(ctx, in.ID)
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
	if err := uc.Repo.UpdateStep(ctx, s); err != nil {
		return StepOutput{}, errors.WithStack(err)
	}
	return StepOutput{Step: s}, nil
}

type DeleteStepInput struct {
	ID string
}

func (uc StepUsecase) DeleteStep(ctx context.Context, in DeleteStepInput) error {
	s, err := uc.Repo.GetStepByID(ctx, in.ID)
	if err != nil {
		if errors.Is(err, repository.ErrModelNotFound) {
			return ErrStepNotFound
		}
		return errors.WithStack(err)
	}
	if auth.GetUserID(ctx) != s.UserID {
		return ErrStepNotFound
	}

	if err := uc.Repo.DeleteStep(ctx, in.ID); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
