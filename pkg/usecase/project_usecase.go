package usecase

import (
	"context"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/minguu42/simoom/pkg/domain/auth"
	"github.com/minguu42/simoom/pkg/domain/idgen"
	"github.com/minguu42/simoom/pkg/domain/model"
	"github.com/minguu42/simoom/pkg/domain/repository"
)

type Project struct {
	repo repository.Repository
}

func NewProject(repo repository.Repository) Project {
	return Project{repo: repo}
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

func (uc Project) CreateProject(ctx context.Context, in CreateProjectInput) (ProjectOutput, error) {
	now := time.Now()
	p := model.Project{
		ID:         idgen.Generate(),
		UserID:     auth.GetUserID(ctx),
		Name:       in.Name,
		Color:      in.Color,
		IsArchived: false,
		CreatedAt:  now,
		UpdatedAt:  now,
	}
	if err := uc.repo.CreateProject(ctx, p); err != nil {
		return ProjectOutput{}, errors.WithStack(err)
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
		return ProjectsOutput{}, errors.WithStack(err)
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
		return ProjectOutput{}, errors.WithStack(err)
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
		return ProjectOutput{}, errors.WithStack(err)
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
		return errors.WithStack(err)
	}
	if auth.GetUserID(ctx) != p.UserID {
		return ErrProjectNotFound
	}

	if err := uc.repo.DeleteProject(ctx, in.ID); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
