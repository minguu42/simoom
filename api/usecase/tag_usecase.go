package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/minguu42/simoom/api/domain/auth"
	"github.com/minguu42/simoom/api/domain/model"
	"github.com/minguu42/simoom/api/domain/repository"
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

func (in CreateTagInput) Create(g model.IDGenerator, userID string) model.Tag {
	return model.Tag{
		ID:     g.Generate(),
		UserID: userID,
		Name:   in.Name,
	}
}

func (uc Tag) CreateTag(ctx context.Context, in CreateTagInput) (TagOutput, error) {
	t := in.Create(uc.idgen, auth.User(ctx).ID)
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
	ts, err := uc.repo.ListTagsByUserID(ctx, auth.User(ctx).ID, in.Limit+1, in.Offset)
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
	if !auth.User(ctx).HasTag(t) {
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

func (uc Tag) DeleteTag(ctx context.Context, in DeleteTagInput) error {
	if err := uc.repo.Transaction(ctx, func(ctxWithTx context.Context) error {
		t, err := uc.repo.GetTagByID(ctxWithTx, in.ID)
		if err != nil {
			if errors.Is(err, repository.ErrModelNotFound) {
				return ErrTagNotFound
			}
			return fmt.Errorf("failed to get tag: %w", err)
		}
		if !auth.User(ctxWithTx).HasTag(t) {
			return ErrTagNotFound
		}

		if err := uc.repo.DeleteTag(ctxWithTx, in.ID); err != nil {
			return fmt.Errorf("failed to delete tag: %w", err)
		}
		return nil
	}); err != nil {
		return fmt.Errorf("failed to run transaction: %w", err)
	}
	return nil
}
