package mysql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/minguu42/simoom/api/adapter/mysql/sqlc"
	"github.com/minguu42/simoom/api/domain/model"
	"github.com/minguu42/simoom/api/domain/repository"
)

func newModelTag(t sqlc.Tag) model.Tag {
	return model.Tag{
		ID:     model.TagID(t.ID),
		UserID: model.UserID(t.UserID),
		Name:   t.Name,
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
	if err := c.queries(ctx).CreateTag(ctx, sqlc.CreateTagParams{
		ID:     string(t.ID),
		UserID: string(t.UserID),
		Name:   t.Name,
	}); err != nil {
		return fmt.Errorf("failed to create tag: %w", err)
	}
	return nil
}

func (c *Client) ListTagsByUserID(ctx context.Context, userID model.UserID, limit, offset uint) ([]model.Tag, error) {
	ts, err := c.queries(ctx).ListTagsByUserID(ctx, sqlc.ListTagsByUserIDParams{
		UserID: string(userID),
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to list tags: %w", err)
	}
	return newModelTags(ts), nil
}

func (c *Client) GetTagByID(ctx context.Context, id model.TagID) (model.Tag, error) {
	t, err := c.queries(ctx).GetTagByID(ctx, string(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Tag{}, repository.ErrModelNotFound
		}
		return model.Tag{}, fmt.Errorf("failed to get tag: %w", err)
	}
	return newModelTag(t), nil
}

func (c *Client) UpdateTag(ctx context.Context, t model.Tag) error {
	if err := c.queries(ctx).UpdateTag(ctx, sqlc.UpdateTagParams{
		Name: t.Name,
		ID:   string(t.ID),
	}); err != nil {
		return fmt.Errorf("failed to update tag: %w", err)
	}
	return nil
}

func (c *Client) DeleteTag(ctx context.Context, id model.TagID) error {
	if err := c.queries(ctx).DeleteTag(ctx, string(id)); err != nil {
		return fmt.Errorf("failed to delete tag: %w", err)
	}
	return nil
}
