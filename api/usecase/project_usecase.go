package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/minguu42/simoom/api/domain/auth"
	"github.com/minguu42/simoom/api/domain/model"
	"github.com/minguu42/simoom/api/domain/repository"
)

type Project struct {
	repo  repository.Repository
	idgen model.IDGenerator
}

func NewProject(repo repository.Repository, idgen model.IDGenerator) Project {
	return Project{
		repo:  repo,
		idgen: idgen,
	}
}

type ProjectOutput struct {
	Project model.Project
}

type ProjectsOutput struct {
	Projects []model.Project
	HasNext  bool
}

type CreateProjectInput struct {
	Name  string
	Color string
}

func (in CreateProjectInput) Create(g model.IDGenerator, userID string) model.Project {
	return model.Project{
		ID:         g.Generate(),
		UserID:     userID,
		Name:       in.Name,
		Color:      in.Color,
		IsArchived: false,
	}
}

func (uc Project) CreateProject(ctx context.Context, in CreateProjectInput) (ProjectOutput, error) {
	p := in.Create(uc.idgen, auth.GetUserID(ctx))
	if err := uc.repo.CreateProject(ctx, p); err != nil {
		return ProjectOutput{}, fmt.Errorf("failed to create project: %w", err)
	}
	return ProjectOutput{Project: p}, nil
}

type ListProjectsInput struct {
	Limit  uint
	Offset uint
}

func (uc Project) ListProjects(ctx context.Context, in ListProjectsInput) (ProjectsOutput, error) {
	ps, err := uc.repo.ListProjectsByUserID(ctx, auth.GetUserID(ctx), in.Limit+1, in.Offset)
	if err != nil {
		return ProjectsOutput{}, fmt.Errorf("failed to list projects: %w", err)
	}

	hasNext := false
	if len(ps) == int(in.Limit+1) {
		ps = ps[:in.Limit]
		hasNext = true
	}
	return ProjectsOutput{
		Projects: ps,
		HasNext:  hasNext,
	}, nil
}

type UpdateProjectInput struct {
	ID         string
	Name       *string
	Color      *string
	IsArchived *bool
}

func (uc Project) UpdateProject(ctx context.Context, in UpdateProjectInput) (ProjectOutput, error) {
	p, err := uc.repo.GetProjectByID(ctx, in.ID)
	if err != nil {
		if errors.Is(err, repository.ErrModelNotFound) {
			return ProjectOutput{}, ErrProjectNotFound
		}
		return ProjectOutput{}, fmt.Errorf("failed to get project: %w", err)
	}
	if auth.GetUserID(ctx) != p.UserID {
		return ProjectOutput{}, ErrProjectNotFound
	}

	if in.Name != nil {
		p.Name = *in.Name
	}
	if in.Color != nil {
		p.Color = *in.Color
	}
	if in.IsArchived != nil {
		p.IsArchived = *in.IsArchived
	}
	if err := uc.repo.UpdateProject(ctx, p); err != nil {
		return ProjectOutput{}, fmt.Errorf("failed to update project: %w", err)
	}
	return ProjectOutput{Project: p}, nil
}

type DeleteProjectInput struct {
	ID string
}

func (uc Project) DeleteProject(ctx context.Context, in DeleteProjectInput) error {
	p, err := uc.repo.GetProjectByID(ctx, in.ID)
	if err != nil {
		if errors.Is(err, repository.ErrModelNotFound) {
			return ErrProjectNotFound
		}
		return fmt.Errorf("failed to get project: %w", err)
	}
	if auth.GetUserID(ctx) != p.UserID {
		return ErrProjectNotFound
	}

	if err := uc.repo.DeleteProject(ctx, in.ID); err != nil {
		return fmt.Errorf("failed to delete project: %w", err)
	}
	return nil
}
