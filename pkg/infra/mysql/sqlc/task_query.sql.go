// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: task_query.sql

package sqlc

import (
	"context"
	"database/sql"
)

const createTask = `-- name: CreateTask :exec
INSERT INTO tasks (id, user_id, project_id, title, content, priority, due_on, completed_at)
VALUES (?, ?, ?, ?, '', ?, NULL, NULL)
`

type CreateTaskParams struct {
	ID        string
	UserID    string
	ProjectID string
	Title     string
	Priority  uint32
}

func (q *Queries) CreateTask(ctx context.Context, arg CreateTaskParams) error {
	_, err := q.db.ExecContext(ctx, createTask,
		arg.ID,
		arg.UserID,
		arg.ProjectID,
		arg.Title,
		arg.Priority,
	)
	return err
}

const deleteTask = `-- name: DeleteTask :exec
DELETE
FROM tasks
WHERE id = ?
`

func (q *Queries) DeleteTask(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteTask, id)
	return err
}

const getTaskByID = `-- name: GetTaskByID :one
SELECT id, user_id, project_id, title, content, priority, due_on, completed_at, created_at, updated_at
FROM tasks
WHERE id = ?
`

func (q *Queries) GetTaskByID(ctx context.Context, id string) (Task, error) {
	row := q.db.QueryRowContext(ctx, getTaskByID, id)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.UserID,
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
SELECT id, user_id, project_id, title, content, priority, due_on, completed_at, created_at, updated_at
FROM tasks
WHERE project_id = ?
ORDER BY created_at
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
			&i.UserID,
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

const listTasksByTagID = `-- name: ListTasksByTagID :many
SELECT t1.id, t1.user_id, t1.project_id, t1.title, t1.content, t1.priority, t1.due_on, t1.completed_at, t1.created_at, t1.updated_at
FROM tasks AS t1
    INNER JOIN tasks_tags AS tt ON t1.id = tt.task_id
WHERE tt.tag_id = ?
ORDER BY t1.created_at
LIMIT ? OFFSET ?
`

type ListTasksByTagIDParams struct {
	TagID  string
	Limit  int32
	Offset int32
}

func (q *Queries) ListTasksByTagID(ctx context.Context, arg ListTasksByTagIDParams) ([]Task, error) {
	rows, err := q.db.QueryContext(ctx, listTasksByTagID, arg.TagID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Task
	for rows.Next() {
		var i Task
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
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
UPDATE tasks
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
