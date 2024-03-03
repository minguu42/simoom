package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/minguu42/simoom/api/apperr"
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
	p := in.Create(uc.idgen, auth.User(ctx).ID)
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
	ps, err := uc.repo.ListProjectsByUserID(ctx, auth.User(ctx).ID, in.Limit+1, in.Offset)
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
			return ProjectOutput{}, apperr.ErrProjectNotFound
		}
		return ProjectOutput{}, fmt.Errorf("failed to get project: %w", err)
	}
	if !auth.User(ctx).HasProject(p) {
		return ProjectOutput{}, apperr.ErrProjectNotFound
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
	if err := uc.repo.Transaction(ctx, func(ctxWithTx context.Context) error {
		p, err := uc.repo.GetProjectByID(ctxWithTx, in.ID)
		if err != nil {
			if errors.Is(err, repository.ErrModelNotFound) {
				return apperr.ErrProjectNotFound
			}
			return fmt.Errorf("failed to get project: %w", err)
		}
		if !auth.User(ctxWithTx).HasProject(p) {
			return apperr.ErrProjectNotFound
		}

		if err := uc.repo.DeleteProject(ctxWithTx, in.ID); err != nil {
			return fmt.Errorf("failed to delete project: %w", err)
		}
		return nil
	}); err != nil {
		return fmt.Errorf("failed to run transaction: %w", err)
	}
	return nil
}
