package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/minguu42/simoom/api/apperr"
	"github.com/minguu42/simoom/api/domain"
	"github.com/minguu42/simoom/api/domain/model"
)

type Tag struct {
	repo  domain.Repository
	idgen domain.IDGenerator
}

func NewTag(repo domain.Repository, idgen domain.IDGenerator) Tag {
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

func (in CreateTagInput) Create(g domain.IDGenerator, userID model.UserID) model.Tag {
	return model.Tag{
		ID:     model.TagID(g.Generate()),
		UserID: userID,
		Name:   in.Name,
	}
}

func (uc Tag) CreateTag(ctx context.Context, in CreateTagInput) (TagOutput, error) {
	t := in.Create(uc.idgen, model.UserFromContext(ctx).ID)
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
	ts, err := uc.repo.ListTagsByUserID(ctx, model.UserFromContext(ctx).ID, in.Limit+1, in.Offset)
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
	ID   model.TagID
	Name *string
}

func (uc Tag) UpdateTag(ctx context.Context, in UpdateTagInput) (TagOutput, error) {
	t, err := uc.repo.GetTagByID(ctx, in.ID)
	if err != nil {
		if errors.Is(err, domain.ErrModelNotFound) {
			return TagOutput{}, apperr.ErrTagNotFound(err)
		}
		return TagOutput{}, fmt.Errorf("failed to get tag: %w", err)
	}
	if !model.UserFromContext(ctx).HasTag(t) {
		return TagOutput{}, apperr.ErrTagNotFound(err)
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
	ID model.TagID
}

func (uc Tag) DeleteTag(ctx context.Context, in DeleteTagInput) error {
	if err := uc.repo.Transaction(ctx, func(ctxWithTx context.Context) error {
		t, err := uc.repo.GetTagByID(ctxWithTx, in.ID)
		if err != nil {
			if errors.Is(err, domain.ErrModelNotFound) {
				return apperr.ErrTagNotFound(err)
			}
			return fmt.Errorf("failed to get tag: %w", err)
		}
		if !model.UserFromContext(ctxWithTx).HasTag(t) {
			return apperr.ErrTagNotFound(err)
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
