package mysql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/minguu42/simoom/api/domain/model"
	"github.com/minguu42/simoom/api/domain/repository"
	"github.com/minguu42/simoom/api/infra/mysql/sqlc"
)

func newModelTask(t sqlc.Task, ss []sqlc.Step, ts []sqlc.Tag) model.Task {
	return model.Task{
		ID:          model.TaskID(t.ID),
		Steps:       newModelSteps(ss),
		Tags:        newModelTags(ts),
		UserID:      model.UserID(t.UserID),
		ProjectID:   model.ProjectID(t.ProjectID),
		Name:        t.Name,
		Content:     t.Content,
		Priority:    uint(t.Priority),
		DueOn:       newPtrTime(t.DueOn),
		CompletedAt: newPtrTime(t.CompletedAt),
	}
}

func newModelTasks(ts []sqlc.Task, steps []sqlc.Step, tags []sqlc.ListTagsByTaskIDsRow) []model.Task {
	stepsByTaskID := sqlc.Steps(steps).StepsByTaskID()
	tagsByTaskID := map[string][]sqlc.Tag{}
	for _, t := range tags {
		tagsByTaskID[t.TaskID] = append(tagsByTaskID[t.TaskID], sqlc.Tag{
			ID:        t.ID,
			UserID:    t.UserID,
			Name:      t.Name,
			CreatedAt: t.CreatedAt,
			UpdatedAt: t.UpdatedAt,
		})
	}

	tasks := make([]model.Task, 0, len(ts))
	for _, t := range ts {
		tasks = append(tasks, newModelTask(t, stepsByTaskID[t.ID], tagsByTaskID[t.ID]))
	}
	return tasks
}

func (c *Client) CreateTask(ctx context.Context, t model.Task) error {
	if err := c.queries(ctx).CreateTask(ctx, sqlc.CreateTaskParams{
		ID:        string(t.ID),
		UserID:    string(t.UserID),
		ProjectID: string(t.ProjectID),
		Name:      t.Name,
		Priority:  uint32(t.Priority),
	}); err != nil {
		return fmt.Errorf("failed to create task: %w", err)
	}
	return nil
}

func (c *Client) ListTasksByUserID(ctx context.Context, userID model.UserID, limit, offset uint, projectID *model.ProjectID, tagID *model.TagID) ([]model.Task, error) {
	var ts []sqlc.Task
	var err error
	switch {
	case projectID != nil && tagID != nil:
		ts, err = c.queries(ctx).ListTasksByProjectIDAndTagID(ctx, sqlc.ListTasksByProjectIDAndTagIDParams{
			ProjectID: string(*projectID),
			TagID:     string(*tagID),
			Limit:     int32(limit),
			Offset:    int32(offset),
		})
	case projectID != nil:
		ts, err = c.queries(ctx).ListTasksByProjectID(ctx, sqlc.ListTasksByProjectIDParams{
			ProjectID: string(*projectID),
			Limit:     int32(limit),
			Offset:    int32(offset),
		})
	case tagID != nil:
		ts, err = c.queries(ctx).ListTasksByTagID(ctx, sqlc.ListTasksByTagIDParams{
			TagID:  string(*tagID),
			Limit:  int32(limit),
			Offset: int32(offset),
		})
	default:
		ts, err = c.queries(ctx).ListTasksByUserID(ctx, sqlc.ListTasksByUserIDParams{
			UserID: string(userID),
			Limit:  int32(limit),
			Offset: int32(offset),
		})
	}
	if err != nil {
		return nil, fmt.Errorf("failed to list tasks: %w", err)
	}

	steps, err := c.queries(ctx).ListStepsByTaskIDs(ctx, sqlc.Tasks(ts).IDs())
	if err != nil {
		return nil, fmt.Errorf("failed to list steps: %w", err)
	}
	tags, err := c.queries(ctx).ListTagsByTaskIDs(ctx, sqlc.Tasks(ts).IDs())
	if err != nil {
		return nil, fmt.Errorf("failed to list tags: %w", err)
	}
	return newModelTasks(ts, steps, tags), nil
}

func (c *Client) GetTaskByID(ctx context.Context, id model.TaskID) (model.Task, error) {
	t, err := c.queries(ctx).GetTaskByID(ctx, string(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Task{}, repository.ErrModelNotFound
		}
		return model.Task{}, fmt.Errorf("failed to get task: %w", err)
	}
	ss, err := c.queries(ctx).ListStepsByTaskID(ctx, t.ID)
	if err != nil {
		return model.Task{}, fmt.Errorf("failed to list steps: %w", err)
	}
	ts, err := c.queries(ctx).ListTagsByTaskID(ctx, t.ID)
	if err != nil {
		return model.Task{}, fmt.Errorf("failed to list tags: %w", err)
	}
	return newModelTask(t, ss, ts), nil
}

func (c *Client) UpdateTask(ctx context.Context, t model.Task) error {
	if err := c.queries(ctx).UpdateTask(ctx, sqlc.UpdateTaskParams{
		Name:        t.Name,
		Content:     t.Content,
		Priority:    uint32(t.Priority),
		DueOn:       newNullTime(t.DueOn),
		CompletedAt: newNullTime(t.CompletedAt),
		ID:          string(t.ID),
	}); err != nil {
		return fmt.Errorf("failed to update task: %w", err)
	}
	return nil
}

func (c *Client) DeleteTask(ctx context.Context, id model.TaskID) error {
	if err := c.queries(ctx).DeleteTask(ctx, string(id)); err != nil {
		return fmt.Errorf("failed to delete task: %w", err)
	}
	return nil
}
