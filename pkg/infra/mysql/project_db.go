package mysql

import (
	"context"
	"database/sql"

	"github.com/cockroachdb/errors"
	"github.com/minguu42/simoom/gen/sqlc"
	"github.com/minguu42/simoom/pkg/domain/model"
	"github.com/minguu42/simoom/pkg/domain/repository"
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

func (c *Client) GetProjectByID(ctx context.Context, id string) (model.Project, error) {
	p, err := sqlc.New(c.sqlDB).GetProjectByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Project{}, repository.ErrModelNotFound
		}
		return model.Project{}, errors.WithStack(err)
	}
	return model.Project{
		ID:         p.ID,
		UserID:     p.UserID,
		Name:       p.Name,
		Color:      p.Color,
		IsArchived: p.IsArchived,
		CreatedAt:  p.CreatedAt,
		UpdatedAt:  p.UpdatedAt,
	}, nil
}

func (c *Client) GetProjectsByUserID(ctx context.Context, userID string, limit, offset uint) ([]model.Project, error) {
	ps, err := sqlc.New(c.sqlDB).GetProjectsByUserID(ctx, sqlc.GetProjectsByUserIDParams{
		UserID: userID,
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	projects := make([]model.Project, 0, len(ps))
	for _, p := range ps {
		project := model.Project{
			ID:         p.ID,
			UserID:     p.UserID,
			Name:       p.Name,
			Color:      p.Color,
			IsArchived: p.IsArchived,
			CreatedAt:  p.CreatedAt,
			UpdatedAt:  p.UpdatedAt,
		}
		projects = append(projects, project)
	}
	return projects, nil
}

func (c *Client) UpdateProject(ctx context.Context, p model.Project) error {
	if err := sqlc.New(c.sqlDB).UpdateProject(ctx, sqlc.UpdateProjectParams{
		Name:       p.Name,
		Color:      p.Color,
		IsArchived: p.IsArchived,
		ID:         p.ID,
	}); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (c *Client) DeleteProject(ctx context.Context, id string) error {
	if err := sqlc.New(c.sqlDB).DeleteProject(ctx, id); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
