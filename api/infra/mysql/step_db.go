package mysql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/minguu42/simoom/api/domain/model"
	"github.com/minguu42/simoom/api/domain/repository"
	"github.com/minguu42/simoom/api/infra/mysql/sqlc"
)

func newModelStep(s sqlc.Step) model.Step {
	return model.Step{
		ID:          s.ID,
		UserID:      s.UserID,
		TaskID:      s.TaskID,
		Name:        s.Name,
		CompletedAt: newPtrTime(s.CompletedAt),
	}
}

func newModelSteps(ss []sqlc.Step) []model.Step {
	steps := make([]model.Step, 0, len(ss))
	for _, s := range ss {
		steps = append(steps, newModelStep(s))
	}
	return steps
}

func (c *Client) CreateStep(ctx context.Context, s model.Step) error {
	if err := c.queries(ctx).CreateStep(ctx, sqlc.CreateStepParams{
		ID:     s.ID,
		UserID: s.UserID,
		TaskID: s.TaskID,
		Name:   s.Name,
	}); err != nil {
		return fmt.Errorf("failed to create step: %w", err)
	}
	return nil
}

func (c *Client) GetStepByID(ctx context.Context, id string) (model.Step, error) {
	s, err := c.queries(ctx).GetStepByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Step{}, repository.ErrModelNotFound
		}
		return model.Step{}, fmt.Errorf("failed to get step: %w", err)
	}
	return newModelStep(s), nil
}

func (c *Client) UpdateStep(ctx context.Context, s model.Step) error {
	if err := c.queries(ctx).UpdateStep(ctx, sqlc.UpdateStepParams{
		Name:        s.Name,
		CompletedAt: newNullTime(s.CompletedAt),
		ID:          s.ID,
	}); err != nil {
		return fmt.Errorf("failed to update step: %w", err)
	}
	return nil
}

func (c *Client) DeleteStep(ctx context.Context, id string) error {
	if err := c.queries(ctx).DeleteStep(ctx, id); err != nil {
		return fmt.Errorf("failed to delete step: %w", err)
	}
	return nil
}
