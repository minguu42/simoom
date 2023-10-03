// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: tag_query.sql

package sqlc

import (
	"context"
	"time"
)

const createTag = `-- name: CreateTag :exec
INSERT INTO tag (id, user_id, name, created_at, updated_at)
VALUES (?, ?, ?, ?, ?)
`

type CreateTagParams struct {
	ID        string
	UserID    string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (q *Queries) CreateTag(ctx context.Context, arg CreateTagParams) error {
	_, err := q.db.ExecContext(ctx, createTag,
		arg.ID,
		arg.UserID,
		arg.Name,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	return err
}

const deleteTag = `-- name: DeleteTag :exec
DELETE
FROM tag
WHERE id = ?
`

func (q *Queries) DeleteTag(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteTag, id)
	return err
}

const getTagByID = `-- name: GetTagByID :one
SELECT id, user_id, name, created_at, updated_at
FROM tag
WHERE id = ?
`

func (q *Queries) GetTagByID(ctx context.Context, id string) (Tag, error) {
	row := q.db.QueryRowContext(ctx, getTagByID, id)
	var i Tag
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listTagsByTaskID = `-- name: ListTagsByTaskID :many
SELECT t.id, t.user_id, t.name, t.created_at, t.updated_at
FROM tag AS t
       INNER JOIN task_tag AS tt ON t.id = tt.tag_id
WHERE tt.task_id = ?
`

func (q *Queries) ListTagsByTaskID(ctx context.Context, taskID string) ([]Tag, error) {
	rows, err := q.db.QueryContext(ctx, listTagsByTaskID, taskID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Tag
	for rows.Next() {
		var i Tag
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Name,
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

const listTagsByUserID = `-- name: ListTagsByUserID :many
SELECT id, user_id, name, created_at, updated_at
FROM tag
WHERE user_id = ?
LIMIT ? OFFSET ?
`

type ListTagsByUserIDParams struct {
	UserID string
	Limit  int32
	Offset int32
}

func (q *Queries) ListTagsByUserID(ctx context.Context, arg ListTagsByUserIDParams) ([]Tag, error) {
	rows, err := q.db.QueryContext(ctx, listTagsByUserID, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Tag
	for rows.Next() {
		var i Tag
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Name,
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

const updateTag = `-- name: UpdateTag :exec
UPDATE tag
SET name = ?
WHERE id = ?
`

type UpdateTagParams struct {
	Name string
	ID   string
}

func (q *Queries) UpdateTag(ctx context.Context, arg UpdateTagParams) error {
	_, err := q.db.ExecContext(ctx, updateTag, arg.Name, arg.ID)
	return err
}
