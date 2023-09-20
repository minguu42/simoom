// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: task_query.sql

package sqlc

import (
	"context"
	"database/sql"
	"time"
)

const createTask = `-- name: CreateTask :exec
INSERT INTO task (id, project_id, title, content, priority, due_on, completed_at, created_at, updated_at)
VALUES (?, ?, ?, '', ?, NULL, NULL, ?, ?)
`

type CreateTaskParams struct {
	ID        string
	ProjectID string
	Title     string
	Priority  uint32
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (q *Queries) CreateTask(ctx context.Context, arg CreateTaskParams) error {
	_, err := q.db.ExecContext(ctx, createTask,
		arg.ID,
		arg.ProjectID,
		arg.Title,
		arg.Priority,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	return err
}

const deleteTask = `-- name: DeleteTask :exec
DELETE
FROM task
WHERE id = ?
`

func (q *Queries) DeleteTask(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteTask, id)
	return err
}

const getTaskByID = `-- name: GetTaskByID :one
SELECT id, project_id, title, content, priority, due_on, completed_at, created_at, updated_at FROM task
WHERE id = ?
LIMIT 1
`

func (q *Queries) GetTaskByID(ctx context.Context, id string) (Task, error) {
	row := q.db.QueryRowContext(ctx, getTaskByID, id)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.ProjectID,
		&i.Title,
		&i.Content,
		&i.Priority,
		&i.DueOn,
		&i.CompletedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listTasksByProjectID = `-- name: ListTasksByProjectID :many
SELECT id, project_id, title, content, priority, due_on, completed_at, created_at, updated_at FROM task
WHERE project_id = ?
ORDER BY created_at DESC
LIMIT ? OFFSET ?
`

type ListTasksByProjectIDParams struct {
	ProjectID string
	Limit     int32
	Offset    int32
}

func (q *Queries) ListTasksByProjectID(ctx context.Context, arg ListTasksByProjectIDParams) ([]Task, error) {
	rows, err := q.db.QueryContext(ctx, listTasksByProjectID, arg.ProjectID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Task
	for rows.Next() {
		var i Task
		if err := rows.Scan(
			&i.ID,
			&i.ProjectID,
			&i.Title,
			&i.Content,
			&i.Priority,
			&i.DueOn,
			&i.CompletedAt,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTask = `-- name: UpdateTask :exec
UPDATE task
SET title        = ?,
    content      = ?,
    priority     = ?,
    due_on       = ?,
    completed_at = ?
WHERE id = ?
`

type UpdateTaskParams struct {
	Title       string
	Content     string
	Priority    uint32
	DueOn       sql.NullTime
	CompletedAt sql.NullTime
	ID          string
}

func (q *Queries) UpdateTask(ctx context.Context, arg UpdateTaskParams) error {
	_, err := q.db.ExecContext(ctx, updateTask,
		arg.Title,
		arg.Content,
		arg.Priority,
		arg.DueOn,
		arg.CompletedAt,
		arg.ID,
	)
	return err
}
