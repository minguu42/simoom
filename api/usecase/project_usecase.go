package usecase

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"

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

func (in CreateProjectInput) Validate() error {
	if utf8.RuneCountInString(in.Name) < 1 || 20 < utf8.RuneCountInString(in.Name) {
		return newErrInvalidArgument("name must be at least 1 and no more than 20 characters")
	}
	if len(in.Color) != 7 || !strings.HasPrefix(in.Color, "#") {
		return newErrInvalidArgument("color is specified in the format #000000")
	}
	return nil
}

func (uc Project) CreateProject(ctx context.Context, in CreateProjectInput) (ProjectOutput, error) {
	// if err := in.Validate(); err != nil {
	// 	return ProjectOutput{}, fmt.Errorf("failed to validate input: %w", err)
	// }

	p := model.Project{
		ID:         uc.idgen.Generate(),
		UserID:     auth.GetUserID(ctx),
		Name:       in.Name,
		Color:      in.Color,
		IsArchived: false,
	}
	if err := uc.repo.CreateProject(ctx, p); err != nil {
		return ProjectOutput{}, fmt.Errorf("failed to create project: %w", err)
	}
	return ProjectOutput{Project: p}, nil
}

type ListProjectsInput struct {
	Limit  uint
	Offset uint
}

func (in ListProjectsInput) Validate() error {
	if in.Limit < 1 {
		return newErrInvalidArgument("limit is greater than or equal to 1")
	}
	return nil
}

func (uc Project) ListProjects(ctx context.Context, in ListProjectsInput) (ProjectsOutput, error) {
	// if err := in.Validate(); err != nil {
	// 	return ProjectsOutput{}, fmt.Errorf("failed to validate input: %w", err)
	// }

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

func (in UpdateProjectInput) Validate() error {
	if len(in.ID) != 26 {
		return newErrInvalidArgument("id is a 26-character string")
	}
	if in.Name == nil && in.Color == nil && in.IsArchived == nil {
		return newErrInvalidArgument("must contain some argument other than id")
	}
	if in.Name != nil && (utf8.RuneCountInString(*in.Name) < 1 || 20 < utf8.RuneCountInString(*in.Name)) {
		return newErrInvalidArgument("name must be at least 1 and no more than 20 characters")
	}
	if in.Color != nil && (len(*in.Color) != 7 || !strings.HasPrefix(*in.Color, "#")) {
		return newErrInvalidArgument("color is specified in the format #000000")
	}
	return nil
}

func (uc Project) UpdateProject(ctx context.Context, in UpdateProjectInput) (ProjectOutput, error) {
	// if err := in.Validate(); err != nil {
	// 	return ProjectOutput{}, fmt.Errorf("failed to validate input: %w", err)
	// }

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

func (in DeleteProjectInput) Validate() error {
	if len(in.ID) != 26 {
		return newErrInvalidArgument("id is a 26-character string")
	}
	return nil
}

func (uc Project) DeleteProject(ctx context.Context, in DeleteProjectInput) error {
	// if err := in.Validate(); err != nil {
	// 	return fmt.Errorf("failed to validate input: %w", err)
	// }

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
