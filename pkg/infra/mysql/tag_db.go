package mysql

import (
	"context"
	"database/sql"

	"github.com/cockroachdb/errors"
	"github.com/minguu42/simoom/pkg/domain/model"
	"github.com/minguu42/simoom/pkg/domain/repository"
	"github.com/minguu42/simoom/pkg/infra/mysql/sqlc"
)

func newModelTag(t sqlc.Tag) model.Tag {
	return model.Tag{
		ID:        t.ID,
		UserID:    t.UserID,
		Name:      t.Name,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}

func newModelTags(ts []sqlc.Tag) []model.Tag {
	tags := make([]model.Tag, 0, len(ts))
	for _, t := range ts {
		tags = append(tags, newModelTag(t))
	}
	return tags
}

func (c *Client) CreateTag(ctx context.Context, t model.Tag) error {
	if err := sqlc.New(c.db).CreateTag(ctx, sqlc.CreateTagParams{
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
	ts, err := sqlc.New(c.db).ListTagsByUserID(ctx, sqlc.ListTagsByUserIDParams{
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
	t, err := sqlc.New(c.db).GetTagByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Tag{}, repository.ErrModelNotFound
		}
		return model.Tag{}, errors.WithStack(err)
	}
	return newModelTag(t), nil
}

func (c *Client) UpdateTag(ctx context.Context, t model.Tag) error {
	if err := sqlc.New(c.db).UpdateTag(ctx, sqlc.UpdateTagParams{
		Name:      t.Name,
		UpdatedAt: t.UpdatedAt,
		ID:        t.ID,
	}); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (c *Client) DeleteTag(ctx context.Context, id string) error {
	if err := sqlc.New(c.db).DeleteTag(ctx, id); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
