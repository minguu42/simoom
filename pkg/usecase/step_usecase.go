package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"
	"unicode/utf8"

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

func (in CreateStepInput) Validate() error {
	if len(in.TaskID) != 26 {
		return newErrInvalidArgument("task_id is a 26-character string")
	}
	if utf8.RuneCountInString(in.Name) < 1 || 80 < utf8.RuneCountInString(in.Name) {
		return newErrInvalidArgument("name must be at least 1 and no more than 80 characters")
	}
	return nil
}

func (uc Step) CreateStep(ctx context.Context, in CreateStepInput) (StepOutput, error) {
	// if err := in.Validate(); err != nil {
	// 	return StepOutput{}, fmt.Errorf("failed to validate input: %w", err)
	// }

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

func (in UpdateStepInput) Validate() error {
	if len(in.ID) != 26 {
		return newErrInvalidArgument("id is a 26-character string")
	}
	if in.Name == nil && in.CompletedAt == nil {
		return newErrInvalidArgument("must contain some argument other than id")
	}
	if in.Name != nil && (utf8.RuneCountInString(*in.Name) < 1 || 80 < utf8.RuneCountInString(*in.Name)) {
		return newErrInvalidArgument("name cannot be an empty string")
	}
	return nil
}

func (uc Step) UpdateStep(ctx context.Context, in UpdateStepInput) (StepOutput, error) {
	// if err := in.Validate(); err != nil {
	// 	return StepOutput{}, fmt.Errorf("failed to validate input: %w", err)
	// }

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

func (in DeleteStepInput) Validate() error {
	if len(in.ID) != 26 {
		return newErrInvalidArgument("id is a 26-character string")
	}
	return nil
}

func (uc Step) DeleteStep(ctx context.Context, in DeleteStepInput) error {
	// if err := in.Validate(); err != nil {
	// 	return fmt.Errorf("failed to validate input: %w", err)
	// }

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
