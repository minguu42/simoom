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
	TaskID string
	Name   string
}

func (uc Step) CreateStep(ctx context.Context, in CreateStepInput) (StepOutput, error) {
	t, err := uc.repo.GetTaskByID(ctx, in.TaskID)
	if err != nil {
		if errors.Is(err, repository.ErrModelNotFound) {
			return StepOutput{}, ErrTaskNotFound
		}
		return StepOutput{}, fmt.Errorf("failed to get task: %w", err)
	}
	if auth.GetUserID(ctx) != t.UserID {
		return StepOutput{}, ErrTaskNotFound
	}

	s := model.Step{
		ID:     uc.idgen.Generate(),
		UserID: auth.GetUserID(ctx),
		TaskID: in.TaskID,
		Name:   in.Name,
	}
	if err := uc.repo.CreateStep(ctx, s); err != nil {
		return StepOutput{}, fmt.Errorf("failed to create step: %w", err)
	}
	return StepOutput{Step: s}, nil
}

type UpdateStepInput struct {
	ID          string
	Name        *string
	CompletedAt *time.Time
}

func (uc Step) UpdateStep(ctx context.Context, in UpdateStepInput) (StepOutput, error) {
	s, err := uc.repo.GetStepByID(ctx, in.ID)
	if err != nil {
		if errors.Is(err, repository.ErrModelNotFound) {
			return StepOutput{}, ErrStepNotFound
		}
		return StepOutput{}, fmt.Errorf("failed to get step: %w", err)
	}
	if auth.GetUserID(ctx) != s.UserID {
		return StepOutput{}, ErrStepNotFound
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
	ID string
}

func (uc Step) DeleteStep(ctx context.Context, in DeleteStepInput) error {
	s, err := uc.repo.GetStepByID(ctx, in.ID)
	if err != nil {
		if errors.Is(err, repository.ErrModelNotFound) {
			return ErrStepNotFound
		}
		return fmt.Errorf("failed to get step: %w", err)
	}
	if auth.GetUserID(ctx) != s.UserID {
		return ErrStepNotFound
	}

	if err := uc.repo.DeleteStep(ctx, in.ID); err != nil {
		return fmt.Errorf("failed to delete step: %w", err)
	}
	return nil
}
