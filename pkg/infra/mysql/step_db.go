package mysql

import (
	"context"
	"database/sql"

	"github.com/cockroachdb/errors"
	"github.com/minguu42/simoom/gen/sqlc"
	"github.com/minguu42/simoom/pkg/domain/model"
	"github.com/minguu42/simoom/pkg/domain/repository"
)

func newModelStep(s sqlc.Step) model.Step {
	return model.Step{
		ID:          s.ID,
		UserID:      s.UserID,
		TaskID:      s.TaskID,
		Title:       s.Title,
		CompletedAt: newPtrTime(s.CompletedAt),
		CreatedAt:   s.CreatedAt,
		UpdatedAt:   s.UpdatedAt,
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
	if err := sqlc.New(c.db).CreateStep(ctx, sqlc.CreateStepParams{
		ID:        s.ID,
		UserID:    s.UserID,
		TaskID:    s.TaskID,
		Title:     s.Title,
		CreatedAt: s.CreatedAt,
		UpdatedAt: s.UpdatedAt,
	}); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (c *Client) GetStepByID(ctx context.Context, id string) (model.Step, error) {
	s, err := sqlc.New(c.db).GetStepByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Step{}, repository.ErrModelNotFound
		}
		return model.Step{}, errors.WithStack(err)
	}
	return newModelStep(s), nil
}

func (c *Client) UpdateStep(ctx context.Context, s model.Step) error {
	if err := sqlc.New(c.db).UpdateStep(ctx, sqlc.UpdateStepParams{
		Title:       s.Title,
		CompletedAt: newNullTime(s.CompletedAt),
		ID:          s.ID,
	}); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (c *Client) DeleteStep(ctx context.Context, id string) error {
	if err := sqlc.New(c.db).DeleteStep(ctx, id); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
