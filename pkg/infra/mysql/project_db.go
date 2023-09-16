package mysql

import (
	"context"
	"github.com/cockroachdb/errors"
	"github.com/minguu42/simoom/gen/sqlc"
	"github.com/minguu42/simoom/pkg/domain/model"
)

func (c *Client) CreateProject(ctx context.Context, p model.Project) error {
	if err := sqlc.New(c.sqlDB).CreateProject(ctx, sqlc.CreateProjectParams{
		ID:         p.ID,
		UserID:     p.UserID,
		Name:       p.Name,
		Color:      p.Color,
		IsArchived: p.IsArchived,
		CreatedAt:  p.CreatedAt,
		UpdatedAt:  p.UpdatedAt,
	}); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
