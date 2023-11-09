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

type TagUsecase struct {
	repo repository.Repository
}

func NewTag(repo repository.Repository) TagUsecase {
	return TagUsecase{repo: repo}
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

func (uc TagUsecase) CreateTag(ctx context.Context, in CreateTagInput) (TagOutput, error) {
	now := time.Now()
	t := model.Tag{
		ID:        idgen.Generate(),
		UserID:    auth.GetUserID(ctx),
		Name:      in.Name,
		CreatedAt: now,
		UpdatedAt: now,
	}
	if err := uc.repo.CreateTag(ctx, t); err != nil {
		return TagOutput{}, errors.WithStack(err)
	}
	return TagOutput{Tag: t}, nil
}

type ListTagsInput struct {
	Limit  uint
	Offset uint
}

func (uc TagUsecase) ListTags(ctx context.Context, in ListTagsInput) (TagsOutput, error) {
	ts, err := uc.repo.ListTagsByUserID(ctx, auth.GetUserID(ctx), in.Limit, in.Offset)
	if err != nil {
		return TagsOutput{}, errors.WithStack(err)
	}
	return TagsOutput{
		Tags:    ts,
		HasNext: false,
	}, nil
}

type UpdateTagInput struct {
	ID   string
	Name *string
}

func (uc TagUsecase) UpdateTag(ctx context.Context, in UpdateTagInput) (TagOutput, error) {
	t, err := uc.repo.GetTagByID(ctx, in.ID)
	if err != nil {
		if errors.Is(err, repository.ErrModelNotFound) {
			return TagOutput{}, ErrTagNotFound
		}
		return TagOutput{}, errors.WithStack(err)
	}
	if auth.GetUserID(ctx) != t.UserID {
		return TagOutput{}, ErrTagNotFound
	}

	if in.Name != nil {
		t.Name = *in.Name
	}
	if err := uc.repo.UpdateTag(ctx, t); err != nil {
		return TagOutput{}, errors.WithStack(err)
	}
	return TagOutput{Tag: t}, nil
}

type DeleteTagInput struct {
	ID string
}

func (uc TagUsecase) DeleteTag(ctx context.Context, in DeleteTagInput) error {
	t, err := uc.repo.GetTagByID(ctx, in.ID)
	if err != nil {
		if errors.Is(err, repository.ErrModelNotFound) {
			return ErrTagNotFound
		}
		return errors.WithStack(err)
	}
	if auth.GetUserID(ctx) != t.UserID {
		return ErrTagNotFound
	}

	if err := uc.repo.DeleteTag(ctx, in.ID); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
