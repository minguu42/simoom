package usecase

import (
	"context"
	"errors"
	"fmt"
	"unicode/utf8"

	"github.com/minguu42/simoom/pkg/domain/auth"
	"github.com/minguu42/simoom/pkg/domain/model"
	"github.com/minguu42/simoom/pkg/domain/repository"
)

type Tag struct {
	repo  repository.Repository
	idgen model.IDGenerator
}

func NewTag(repo repository.Repository, idgen model.IDGenerator) Tag {
	return Tag{
		repo:  repo,
		idgen: idgen,
	}
}

type TagOutput struct {
	Tag model.Tag
}

type TagsOutput struct {
	Tags    []model.Tag
	HasNext bool
}

type CreateTagInput struct {
	Name string
}

func (in CreateTagInput) Validate() error {
	if utf8.RuneCountInString(in.Name) < 1 || 20 < utf8.RuneCountInString(in.Name) {
		return newErrInvalidArgument("name must be at least 1 and no more than 20 characters")
	}
	return nil
}

func (uc Tag) CreateTag(ctx context.Context, in CreateTagInput) (TagOutput, error) {
	// if err := in.Validate(); err != nil {
	// 	return TagOutput{}, fmt.Errorf("failed to validate input: %w", err)
	// }

	t := model.Tag{
		ID:     uc.idgen.Generate(),
		UserID: auth.GetUserID(ctx),
		Name:   in.Name,
	}
	if err := uc.repo.CreateTag(ctx, t); err != nil {
		return TagOutput{}, fmt.Errorf("failed to create tag: %w", err)
	}
	return TagOutput{Tag: t}, nil
}

type ListTagsInput struct {
	Limit  uint
	Offset uint
}

func (in ListTagsInput) Validate() error {
	if in.Limit < 1 {
		return newErrInvalidArgument("limit is greater than or equal to 1")
	}
	return nil
}

func (uc Tag) ListTags(ctx context.Context, in ListTagsInput) (TagsOutput, error) {
	// if err := in.Validate(); err != nil {
	// 	return TagsOutput{}, fmt.Errorf("failed to validate input: %w", err)
	// }

	ts, err := uc.repo.ListTagsByUserID(ctx, auth.GetUserID(ctx), in.Limit+1, in.Offset)
	if err != nil {
		return TagsOutput{}, fmt.Errorf("failed to list tags: %w", err)
	}

	hasNext := false
	if len(ts) == int(in.Limit+1) {
		ts = ts[:in.Limit]
		hasNext = true
	}
	return TagsOutput{
		Tags:    ts,
		HasNext: hasNext,
	}, nil
}

type UpdateTagInput struct {
	ID   string
	Name *string
}

func (in UpdateTagInput) Validate() error {
	if len(in.ID) != 26 {
		return newErrInvalidArgument("id is a 26-character string")
	}
	if in.Name == nil {
		return newErrInvalidArgument("must contain some argument other than id")
	}
	if in.Name != nil && (utf8.RuneCountInString(*in.Name) < 1 || 20 < utf8.RuneCountInString(*in.Name)) {
		return newErrInvalidArgument("name cannot be an empty string")
	}
	return nil
}

func (uc Tag) UpdateTag(ctx context.Context, in UpdateTagInput) (TagOutput, error) {
	// if err := in.Validate(); err != nil {
	// 	return TagOutput{}, fmt.Errorf("failed to validate input: %w", err)
	// }

	t, err := uc.repo.GetTagByID(ctx, in.ID)
	if err != nil {
		if errors.Is(err, repository.ErrModelNotFound) {
			return TagOutput{}, ErrTagNotFound
		}
		return TagOutput{}, fmt.Errorf("failed to get tag: %w", err)
	}
	if auth.GetUserID(ctx) != t.UserID {
		return TagOutput{}, ErrTagNotFound
	}

	if in.Name != nil {
		t.Name = *in.Name
	}
	if err := uc.repo.UpdateTag(ctx, t); err != nil {
		return TagOutput{}, fmt.Errorf("failed to update tag: %w", err)
	}
	return TagOutput{Tag: t}, nil
}

type DeleteTagInput struct {
	ID string
}

func (in DeleteTagInput) Validate() error {
	if len(in.ID) != 26 {
		return newErrInvalidArgument("id is a 26-character string")
	}
	return nil
}

func (uc Tag) DeleteTag(ctx context.Context, in DeleteTagInput) error {
	// if err := in.Validate(); err != nil {
	// 	return fmt.Errorf("failed to validate input: %w", err)
	// }

	t, err := uc.repo.GetTagByID(ctx, in.ID)
	if err != nil {
		if errors.Is(err, repository.ErrModelNotFound) {
			return ErrTagNotFound
		}
		return fmt.Errorf("failed to get tag: %w", err)
	}
	if auth.GetUserID(ctx) != t.UserID {
		return ErrTagNotFound
	}

	if err := uc.repo.DeleteTag(ctx, in.ID); err != nil {
		return fmt.Errorf("failed to delete tag: %w", err)
	}
	return nil
}
