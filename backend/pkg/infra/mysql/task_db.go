package mysql

import (
	"context"
	"database/sql"

	"github.com/cockroachdb/errors"
	"github.com/minguu42/simoom/backend/pkg/domain/model"
	"github.com/minguu42/simoom/backend/pkg/domain/repository"
	"github.com/minguu42/simoom/backend/pkg/infra/mysql/sqlc"
)

func newModelTask(t sqlc.Task, ss []sqlc.Step, ts []sqlc.Tag) model.Task {
	return model.Task{
		ID:          t.ID,
		Steps:       newModelSteps(ss),
		Tags:        newModelTags(ts),
		UserID:      t.UserID,
		ProjectID:   t.ProjectID,
		Title:       t.Title,
		Content:     t.Content,
		Priority:    uint(t.Priority),
		DueOn:       newPtrTime(t.DueOn),
		CompletedAt: newPtrTime(t.CompletedAt),
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}
}

func (c *Client) CreateTask(ctx context.Context, t model.Task) error {
	if err := sqlc.New(c.db).CreateTask(ctx, sqlc.CreateTaskParams{
		ID:        t.ID,
		UserID:    t.UserID,
		ProjectID: t.ProjectID,
		Title:     t.Title,
		Priority:  uint32(t.Priority),
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (c *Client) ListTasksByProjectID(ctx context.Context, projectID string, limit, offset uint) ([]model.Task, error) {
	ts, err := sqlc.New(c.db).ListTasksByProjectID(ctx, sqlc.ListTasksByProjectIDParams{
		ProjectID: projectID,
		Limit:     int32(limit),
		Offset:    int32(offset),
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	tasks := make([]model.Task, 0, len(ts))
	for _, t := range ts {
		ss, err := sqlc.New(c.db).ListStepsByTaskID(ctx, t.ID)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		tags, err := sqlc.New(c.db).ListTagsByTaskID(ctx, t.ID)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		tasks = append(tasks, newModelTask(t, ss, tags))
	}
	return tasks, nil
}

func (c *Client) ListTasksByTagID(ctx context.Context, tagID string, limit, offset uint) ([]model.Task, error) {
	ts, err := sqlc.New(c.db).ListTasksByTagID(ctx, sqlc.ListTasksByTagIDParams{
		TagID:  tagID,
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	tasks := make([]model.Task, 0, len(ts))
	for _, t := range ts {
		ss, err := sqlc.New(c.db).ListStepsByTaskID(ctx, t.ID)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		tags, err := sqlc.New(c.db).ListTagsByTaskID(ctx, t.ID)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		tasks = append(tasks, newModelTask(t, ss, tags))
	}
	return tasks, nil
}

func (c *Client) GetTaskByID(ctx context.Context, id string) (model.Task, error) {
	t, err := sqlc.New(c.db).GetTaskByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Task{}, repository.ErrModelNotFound
		}
		return model.Task{}, errors.WithStack(err)
	}
	ss, err := sqlc.New(c.db).ListStepsByTaskID(ctx, t.ID)
	if err != nil {
		return model.Task{}, errors.WithStack(err)
	}
	ts, err := sqlc.New(c.db).ListTagsByTaskID(ctx, t.ID)
	if err != nil {
		return model.Task{}, errors.WithStack(err)
	}
	return newModelTask(t, ss, ts), nil
}

func (c *Client) UpdateTask(ctx context.Context, t model.Task) error {
	if err := sqlc.New(c.db).UpdateTask(ctx, sqlc.UpdateTaskParams{
		Title:       t.Title,
		Content:     t.Content,
		Priority:    uint32(t.Priority),
		DueOn:       newNullTime(t.DueOn),
		CompletedAt: newNullTime(t.CompletedAt),
		ID:          t.ID,
	}); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (c *Client) DeleteTask(ctx context.Context, id string) error {
	if err := sqlc.New(c.db).DeleteTask(ctx, id); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
