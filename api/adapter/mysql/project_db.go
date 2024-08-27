package mysql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/minguu42/simoom/api/adapter/mysql/sqlc"
	"github.com/minguu42/simoom/api/domain"
	"github.com/minguu42/simoom/api/domain/model"
)

func newModelProject(p sqlc.Project) model.Project {
	return model.Project{
		ID:         model.ProjectID(p.ID),
		UserID:     model.UserID(p.UserID),
		Name:       p.Name,
		Color:      p.Color,
		IsArchived: p.IsArchived,
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
	if err := c.queries(ctx).CreateProject(ctx, sqlc.CreateProjectParams{
		ID:         string(p.ID),
		UserID:     string(p.UserID),
		Name:       p.Name,
		Color:      p.Color,
		IsArchived: p.IsArchived,
	}); err != nil {
		return fmt.Errorf("failed to create project: %w", err)
	}
	return nil
}

func (c *Client) ListProjectsByUserID(ctx context.Context, userID model.UserID, limit, offset uint) ([]model.Project, error) {
	ps, err := c.queries(ctx).ListProjectsByUserID(ctx, sqlc.ListProjectsByUserIDParams{
		UserID: string(userID),
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to list projects: %w", err)
	}
	return newModelProjects(ps), nil
}

func (c *Client) GetProjectByID(ctx context.Context, id model.ProjectID) (model.Project, error) {
	p, err := c.queries(ctx).GetProjectByID(ctx, string(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Project{}, domain.ErrModelNotFound
		}
		return model.Project{}, fmt.Errorf("failed to get project: %w", err)
	}
	return newModelProject(p), nil
}

func (c *Client) UpdateProject(ctx context.Context, p model.Project) error {
	if err := c.queries(ctx).UpdateProject(ctx, sqlc.UpdateProjectParams{
		Name:       p.Name,
		Color:      p.Color,
		IsArchived: p.IsArchived,
		ID:         string(p.ID),
	}); err != nil {
		return fmt.Errorf("failed to update project: %w", err)
	}
	return nil
}

func (c *Client) DeleteProject(ctx context.Context, id model.ProjectID) error {
	if err := c.queries(ctx).DeleteProject(ctx, string(id)); err != nil {
		return fmt.Errorf("failed to delete project: %w", err)
	}
	return nil
}
