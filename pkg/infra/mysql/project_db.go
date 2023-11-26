package mysql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/minguu42/simoom/pkg/domain/model"
	"github.com/minguu42/simoom/pkg/domain/repository"
	"github.com/minguu42/simoom/pkg/infra/mysql/sqlc"
)

func newModelProject(p sqlc.Project) model.Project {
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

func newModelProjects(ps []sqlc.Project) []model.Project {
	projects := make([]model.Project, 0, len(ps))
	for _, p := range ps {
		projects = append(projects, newModelProject(p))
	}
	return projects
}

func (c *Client) CreateProject(ctx context.Context, p model.Project) error {
	if err := sqlc.New(c.db).CreateProject(ctx, sqlc.CreateProjectParams{
		ID:         p.ID,
		UserID:     p.UserID,
		Name:       p.Name,
		Color:      p.Color,
		IsArchived: p.IsArchived,
		CreatedAt:  p.CreatedAt,
		UpdatedAt:  p.UpdatedAt,
	}); err != nil {
		return fmt.Errorf("failed to create project: %w", err)
	}
	return nil
}

func (c *Client) ListProjectsByUserID(ctx context.Context, userID string, limit, offset uint) ([]model.Project, error) {
	ps, err := sqlc.New(c.db).ListProjectsByUserID(ctx, sqlc.ListProjectsByUserIDParams{
		UserID: userID,
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to list projects: %w", err)
	}
	return newModelProjects(ps), nil
}

func (c *Client) GetProjectByID(ctx context.Context, id string) (model.Project, error) {
	p, err := sqlc.New(c.db).GetProjectByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Project{}, repository.ErrModelNotFound
		}
		return model.Project{}, fmt.Errorf("failed to get project: %w", err)
	}
	return newModelProject(p), nil
}

func (c *Client) UpdateProject(ctx context.Context, p model.Project) error {
	if err := sqlc.New(c.db).UpdateProject(ctx, sqlc.UpdateProjectParams{
		Name:       p.Name,
		Color:      p.Color,
		IsArchived: p.IsArchived,
		UpdatedAt:  p.UpdatedAt,
		ID:         p.ID,
	}); err != nil {
		return fmt.Errorf("failed to update project: %w", err)
	}
	return nil
}

func (c *Client) DeleteProject(ctx context.Context, id string) error {
	if err := sqlc.New(c.db).DeleteProject(ctx, id); err != nil {
		return fmt.Errorf("failed to delete project: %w", err)
	}
	return nil
}
