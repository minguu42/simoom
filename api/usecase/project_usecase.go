package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/minguu42/simoom/api/apperr"
	"github.com/minguu42/simoom/api/domain"
	"github.com/minguu42/simoom/api/domain/model"
)

type Project struct {
	repo  domain.Repository
	idgen domain.IDGenerator
}

func NewProject(repo domain.Repository, idgen domain.IDGenerator) Project {
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

func (in CreateProjectInput) Create(g domain.IDGenerator, userID model.UserID) model.Project {
	return model.Project{
		ID:         model.ProjectID(g.Generate()),
		UserID:     userID,
		Name:       in.Name,
		Color:      in.Color,
		IsArchived: false,
	}
}

func (uc Project) CreateProject(ctx context.Context, in CreateProjectInput) (ProjectOutput, error) {
	p := in.Create(uc.idgen, model.UserFromContext(ctx).ID)
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
	ps, err := uc.repo.ListProjectsByUserID(ctx, model.UserFromContext(ctx).ID, in.Limit+1, in.Offset)
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
	ID         model.ProjectID
	Name       *string
	Color      *string
	IsArchived *bool
}

func (uc Project) UpdateProject(ctx context.Context, in UpdateProjectInput) (ProjectOutput, error) {
	p, err := uc.repo.GetProjectByID(ctx, in.ID)
	if err != nil {
		if errors.Is(err, domain.ErrModelNotFound) {
			return ProjectOutput{}, apperr.ErrProjectNotFound(err)
		}
		return ProjectOutput{}, fmt.Errorf("failed to get project: %w", err)
	}
	if !model.UserFromContext(ctx).HasProject(p) {
		return ProjectOutput{}, apperr.ErrProjectNotFound(err)
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
	ID model.ProjectID
}

func (uc Project) DeleteProject(ctx context.Context, in DeleteProjectInput) error {
	if err := uc.repo.Transaction(ctx, func(ctxWithTx context.Context) error {
		p, err := uc.repo.GetProjectByID(ctxWithTx, in.ID)
		if err != nil {
			if errors.Is(err, domain.ErrModelNotFound) {
				return apperr.ErrProjectNotFound(err)
			}
			return fmt.Errorf("failed to get project: %w", err)
		}
		if !model.UserFromContext(ctxWithTx).HasProject(p) {
			return apperr.ErrProjectNotFound(err)
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
