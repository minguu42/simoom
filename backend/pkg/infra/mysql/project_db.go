package mysql

import (
	"context"
	"database/sql"

	"github.com/cockroachdb/errors"
	"github.com/minguu42/simoom/backend/pkg/domain/model"
	"github.com/minguu42/simoom/backend/pkg/domain/repository"
	sqlc2 "github.com/minguu42/simoom/backend/pkg/infra/mysql/sqlc"
)

func newModelProject(p sqlc2.Project) model.Project {
	return model.Project{
		ID:         p.ID,
		UserID:     p.UserID,
		Name:       p.Name,
		Color:      p.Color,
		IsArchived: p.IsArchived,
		CreatedAt:  p.CreatedAt,
		UpdatedAt:  p.UpdatedAt,
	}
}

func newModelProjects(ps []sqlc2.Project) []model.Project {
	projects := make([]model.Project, 0, len(ps))
	for _, p := range ps {
		projects = append(projects, newModelProject(p))
	}
	return projects
}

func (c *Client) CreateProject(ctx context.Context, p model.Project) error {
	if err := sqlc2.New(c.db).CreateProject(ctx, sqlc2.CreateProjectParams{
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

func (c *Client) ListProjectsByUserID(ctx context.Context, userID string, limit, offset uint) ([]model.Project, error) {
	ps, err := sqlc2.New(c.db).ListProjectsByUserID(ctx, sqlc2.ListProjectsByUserIDParams{
		UserID: userID,
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return newModelProjects(ps), nil
}

func (c *Client) GetProjectByID(ctx context.Context, id string) (model.Project, error) {
	p, err := sqlc2.New(c.db).GetProjectByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Project{}, repository.ErrModelNotFound
		}
		return model.Project{}, errors.WithStack(err)
	}
	return newModelProject(p), nil
}

func (c *Client) UpdateProject(ctx context.Context, p model.Project) error {
	if err := sqlc2.New(c.db).UpdateProject(ctx, sqlc2.UpdateProjectParams{
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
	if err := sqlc2.New(c.db).DeleteProject(ctx, id); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
