package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/minguu42/simoom/pkg/clock"
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

func (uc Tag) CreateTag(ctx context.Context, in CreateTagInput) (TagOutput, error) {
	now := clock.Now(ctx)
	t := model.Tag{
		ID:        uc.idgen.Generate(),
		UserID:    auth.GetUserID(ctx),
		Name:      in.Name,
		CreatedAt: now,
		UpdatedAt: now,
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

func (uc Tag) ListTags(ctx context.Context, in ListTagsInput) (TagsOutput, error) {
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

func (uc Tag) UpdateTag(ctx context.Context, in UpdateTagInput) (TagOutput, error) {
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
	t.UpdatedAt = clock.Now(ctx)
	if err := uc.repo.UpdateTag(ctx, t); err != nil {
		return TagOutput{}, fmt.Errorf("failed to update tag: %w", err)
	}
	return TagOutput{Tag: t}, nil
}

type DeleteTagInput struct {
	ID string
}

func (uc Tag) DeleteTag(ctx context.Context, in DeleteTagInput) error {
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
