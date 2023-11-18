package mysql

import (
	"context"
	"database/sql"

	"github.com/cockroachdb/errors"
	"github.com/minguu42/simoom/backend/pkg/domain/model"
	"github.com/minguu42/simoom/backend/pkg/domain/repository"
	sqlc2 "github.com/minguu42/simoom/backend/pkg/infra/mysql/sqlc"
)

func newModelTag(t sqlc2.Tag) model.Tag {
	return model.Tag{
		ID:        t.ID,
		UserID:    t.UserID,
		Name:      t.Name,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}

func newModelTags(ts []sqlc2.Tag) []model.Tag {
	tags := make([]model.Tag, 0, len(ts))
	for _, t := range ts {
		tags = append(tags, newModelTag(t))
	}
	return tags
}

func (c *Client) CreateTag(ctx context.Context, t model.Tag) error {
	if err := sqlc2.New(c.db).CreateTag(ctx, sqlc2.CreateTagParams{
		ID:        t.ID,
		UserID:    t.UserID,
		Name:      t.Name,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (c *Client) ListTagsByUserID(ctx context.Context, userID string, limit, offset uint) ([]model.Tag, error) {
	ts, err := sqlc2.New(c.db).ListTagsByUserID(ctx, sqlc2.ListTagsByUserIDParams{
		UserID: userID,
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return newModelTags(ts), nil
}

func (c *Client) GetTagByID(ctx context.Context, id string) (model.Tag, error) {
	t, err := sqlc2.New(c.db).GetTagByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Tag{}, repository.ErrModelNotFound
		}
		return model.Tag{}, errors.WithStack(err)
	}
	return newModelTag(t), nil
}

func (c *Client) UpdateTag(ctx context.Context, t model.Tag) error {
	if err := sqlc2.New(c.db).UpdateTag(ctx, sqlc2.UpdateTagParams{
		Name: t.Name,
		ID:   t.ID,
	}); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (c *Client) DeleteTag(ctx context.Context, id string) error {
	if err := sqlc2.New(c.db).DeleteTag(ctx, id); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
