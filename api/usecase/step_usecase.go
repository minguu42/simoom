package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/minguu42/simoom/api/apperr"
	"github.com/minguu42/simoom/api/domain/auth"
	"github.com/minguu42/simoom/api/domain/model"
	"github.com/minguu42/simoom/api/domain/repository"
)

type Step struct {
	repo  repository.Repository
	idgen model.IDGenerator
}

func NewStep(repo repository.Repository, idgen model.IDGenerator) Step {
	return Step{
		repo:  repo,
		idgen: idgen,
	}
}

type StepOutput struct {
	Step model.Step
}

type StepsOutput struct {
	Steps   []model.Step
	HasNext bool
}

type CreateStepInput struct {
	TaskID model.TaskID
	Name   string
}

func (in CreateStepInput) Create(g model.IDGenerator, userID model.UserID) model.Step {
	return model.Step{
		ID:     model.StepID(g.Generate()),
		UserID: userID,
		TaskID: in.TaskID,
		Name:   in.Name,
	}
}

func (uc Step) CreateStep(ctx context.Context, in CreateStepInput) (StepOutput, error) {
	t, err := uc.repo.GetTaskByID(ctx, in.TaskID)
	if err != nil {
		if errors.Is(err, repository.ErrModelNotFound) {
			return StepOutput{}, apperr.ErrTaskNotFound(err)
		}
		return StepOutput{}, fmt.Errorf("failed to get task: %w", err)
	}
	user := auth.User(ctx)
	if !user.HasTask(t) {
		return StepOutput{}, apperr.ErrTaskNotFound(err)
	}

	s := in.Create(uc.idgen, user.ID)
	if err := uc.repo.CreateStep(ctx, s); err != nil {
		return StepOutput{}, fmt.Errorf("failed to create step: %w", err)
	}
	return StepOutput{Step: s}, nil
}

type UpdateStepInput struct {
	ID          model.StepID
	Name        *string
	CompletedAt *time.Time
}

func (uc Step) UpdateStep(ctx context.Context, in UpdateStepInput) (StepOutput, error) {
	s, err := uc.repo.GetStepByID(ctx, in.ID)
	if err != nil {
		if errors.Is(err, repository.ErrModelNotFound) {
			return StepOutput{}, apperr.ErrStepNotFound(err)
		}
		return StepOutput{}, fmt.Errorf("failed to get step: %w", err)
	}
	if !auth.User(ctx).HasStep(s) {
		return StepOutput{}, apperr.ErrStepNotFound(err)
	}

	if in.Name != nil {
		s.Name = *in.Name
	}
	if in.CompletedAt != nil {
		s.CompletedAt = in.CompletedAt
	}
	if err := uc.repo.UpdateStep(ctx, s); err != nil {
		return StepOutput{}, fmt.Errorf("failed to update step: %w", err)
	}
	return StepOutput{Step: s}, nil
}

type DeleteStepInput struct {
	ID model.StepID
}

func (uc Step) DeleteStep(ctx context.Context, in DeleteStepInput) error {
	if err := uc.repo.Transaction(ctx, func(ctxWithTx context.Context) error {
		s, err := uc.repo.GetStepByID(ctxWithTx, in.ID)
		if err != nil {
			if errors.Is(err, repository.ErrModelNotFound) {
				return apperr.ErrStepNotFound(err)
			}
			return fmt.Errorf("failed to get step: %w", err)
		}
		if !auth.User(ctx).HasStep(s) {
			return apperr.ErrStepNotFound(err)
		}

		if err := uc.repo.DeleteStep(ctxWithTx, in.ID); err != nil {
			return fmt.Errorf("failed to delete step: %w", err)
		}
		return nil
	}); err != nil {
		return fmt.Errorf("failed to run transaction: %w", err)
	}
	return nil
}
