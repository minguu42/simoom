// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: queries.sql

package sqlc

import (
	"context"
	"database/sql"
)

const createProject = `-- name: CreateProject :exec
INSERT INTO projects (id, user_id, name, color, is_archived)
VALUES (?, ?, ?, ?, ?)
`

type CreateProjectParams struct {
	ID         string
	UserID     string
	Name       string
	Color      string
	IsArchived bool
}

func (q *Queries) CreateProject(ctx context.Context, arg CreateProjectParams) error {
	_, err := q.db.ExecContext(ctx, createProject,
		arg.ID,
		arg.UserID,
		arg.Name,
		arg.Color,
		arg.IsArchived,
	)
	return err
}

const createStep = `-- name: CreateStep :exec
INSERT INTO steps (id, user_id, task_id, name, completed_at)
VALUES (?, ?, ?, ?, NULL)
`

type CreateStepParams struct {
	ID     string
	UserID string
	TaskID string
	Name   string
}

func (q *Queries) CreateStep(ctx context.Context, arg CreateStepParams) error {
	_, err := q.db.ExecContext(ctx, createStep,
		arg.ID,
		arg.UserID,
		arg.TaskID,
		arg.Name,
	)
	return err
}

const createTag = `-- name: CreateTag :exec
INSERT INTO tags (id, user_id, name)
VALUES (?, ?, ?)
`

type CreateTagParams struct {
	ID     string
	UserID string
	Name   string
}

func (q *Queries) CreateTag(ctx context.Context, arg CreateTagParams) error {
	_, err := q.db.ExecContext(ctx, createTag, arg.ID, arg.UserID, arg.Name)
	return err
}

const createTask = `-- name: CreateTask :exec
INSERT INTO tasks (id, user_id, project_id, name, content, priority, due_on, completed_at)
VALUES (?, ?, ?, ?, '', ?, NULL, NULL)
`

type CreateTaskParams struct {
	ID        string
	UserID    string
	ProjectID string
	Name      string
	Priority  uint32
}

func (q *Queries) CreateTask(ctx context.Context, arg CreateTaskParams) error {
	_, err := q.db.ExecContext(ctx, createTask,
		arg.ID,
		arg.UserID,
		arg.ProjectID,
		arg.Name,
		arg.Priority,
	)
	return err
}

const createUser = `-- name: CreateUser :exec
INSERT INTO users (id, name, email, password)
VALUES (?, ?, ?, ?)
`

type CreateUserParams struct {
	ID       string
	Name     string
	Email    string
	Password string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.db.ExecContext(ctx, createUser,
		arg.ID,
		arg.Name,
		arg.Email,
		arg.Password,
	)
	return err
}

const deleteProject = `-- name: DeleteProject :exec
DELETE
FROM projects
WHERE id = ?
`

func (q *Queries) DeleteProject(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteProject, id)
	return err
}

const deleteStep = `-- name: DeleteStep :exec
DELETE
FROM steps
WHERE id = ?
`

func (q *Queries) DeleteStep(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteStep, id)
	return err
}

const deleteTag = `-- name: DeleteTag :exec
DELETE
FROM tags
WHERE id = ?
`

func (q *Queries) DeleteTag(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteTag, id)
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

const getProjectByID = `-- name: GetProjectByID :one
SELECT id, user_id, name, color, is_archived, created_at, updated_at
FROM projects
WHERE id = ?
`

func (q *Queries) GetProjectByID(ctx context.Context, id string) (Project, error) {
	row := q.db.QueryRowContext(ctx, getProjectByID, id)
	var i Project
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Name,
		&i.Color,
		&i.IsArchived,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getStepByID = `-- name: GetStepByID :one
SELECT id, user_id, task_id, name, completed_at, created_at, updated_at
FROM steps
WHERE id = ?
`

func (q *Queries) GetStepByID(ctx context.Context, id string) (Step, error) {
	row := q.db.QueryRowContext(ctx, getStepByID, id)
	var i Step
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.TaskID,
		&i.Name,
		&i.CompletedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getTagByID = `-- name: GetTagByID :one
SELECT id, user_id, name, created_at, updated_at
FROM tags
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

const getTaskByID = `-- name: GetTaskByID :one
SELECT id, user_id, project_id, name, content, priority, due_on, completed_at, created_at, updated_at
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
		&i.Name,
		&i.Content,
		&i.Priority,
		&i.DueOn,
		&i.CompletedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, name, email, password, created_at, updated_at
FROM users
WHERE email = ?
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, name, email, password, created_at, updated_at
FROM users
WHERE id = ?
`

func (q *Queries) GetUserByID(ctx context.Context, id string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listProjectsByUserID = `-- name: ListProjectsByUserID :many
SELECT id, user_id, name, color, is_archived, created_at, updated_at
FROM projects
WHERE user_id = ?
ORDER BY created_at DESC
LIMIT ? OFFSET ?
`

type ListProjectsByUserIDParams struct {
	UserID string
	Limit  int32
	Offset int32
}

func (q *Queries) ListProjectsByUserID(ctx context.Context, arg ListProjectsByUserIDParams) ([]Project, error) {
	rows, err := q.db.QueryContext(ctx, listProjectsByUserID, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Project
	for rows.Next() {
		var i Project
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Name,
			&i.Color,
			&i.IsArchived,
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

const listStepsByTaskID = `-- name: ListStepsByTaskID :many
SELECT id, user_id, task_id, name, completed_at, created_at, updated_at
FROM steps
WHERE task_id = ?
ORDER BY created_at
`

func (q *Queries) ListStepsByTaskID(ctx context.Context, taskID string) ([]Step, error) {
	rows, err := q.db.QueryContext(ctx, listStepsByTaskID, taskID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Step
	for rows.Next() {
		var i Step
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.TaskID,
			&i.Name,
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

const listTagsByTaskID = `-- name: ListTagsByTaskID :many
SELECT t.id, t.user_id, t.name, t.created_at, t.updated_at
FROM tags AS t
    INNER JOIN tasks_tags AS tt ON t.id = tt.tag_id
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
FROM tags
WHERE user_id = ?
ORDER BY created_at
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

const listTasksByProjectID = `-- name: ListTasksByProjectID :many
SELECT id, user_id, project_id, name, content, priority, due_on, completed_at, created_at, updated_at
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
			&i.Name,
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

const listTasksByProjectIDAndTagID = `-- name: ListTasksByProjectIDAndTagID :many
SELECT t.id, t.user_id, t.project_id, t.name, t.content, t.priority, t.due_on, t.completed_at, t.created_at, t.updated_at
FROM tasks AS t
    INNER JOIN tasks_tags AS tt ON t.id = tt.task_id
WHERE t.project_id = ? AND tt.tag_id = ?
ORDER BY t.created_at
LIMIT ? OFFSET ?
`

type ListTasksByProjectIDAndTagIDParams struct {
	ProjectID string
	TagID     string
	Limit     int32
	Offset    int32
}

func (q *Queries) ListTasksByProjectIDAndTagID(ctx context.Context, arg ListTasksByProjectIDAndTagIDParams) ([]Task, error) {
	rows, err := q.db.QueryContext(ctx, listTasksByProjectIDAndTagID,
		arg.ProjectID,
		arg.TagID,
		arg.Limit,
		arg.Offset,
	)
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
			&i.Name,
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
SELECT t.id, t.user_id, t.project_id, t.name, t.content, t.priority, t.due_on, t.completed_at, t.created_at, t.updated_at
FROM tasks AS t
    INNER JOIN tasks_tags AS tt ON t.id = tt.task_id
WHERE tt.tag_id = ?
ORDER BY t.created_at
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
			&i.Name,
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

const listTasksByUserID = `-- name: ListTasksByUserID :many
SELECT id, user_id, project_id, name, content, priority, due_on, completed_at, created_at, updated_at
FROM tasks
WHERE user_id = ?
ORDER BY created_at
LIMIT ? OFFSET ?
`

type ListTasksByUserIDParams struct {
	UserID string
	Limit  int32
	Offset int32
}

func (q *Queries) ListTasksByUserID(ctx context.Context, arg ListTasksByUserIDParams) ([]Task, error) {
	rows, err := q.db.QueryContext(ctx, listTasksByUserID, arg.UserID, arg.Limit, arg.Offset)
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
			&i.Name,
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

const updateProject = `-- name: UpdateProject :exec
UPDATE projects
SET name        = ?,
    color       = ?,
    is_archived = ?
WHERE id = ?
`

type UpdateProjectParams struct {
	Name       string
	Color      string
	IsArchived bool
	ID         string
}

func (q *Queries) UpdateProject(ctx context.Context, arg UpdateProjectParams) error {
	_, err := q.db.ExecContext(ctx, updateProject,
		arg.Name,
		arg.Color,
		arg.IsArchived,
		arg.ID,
	)
	return err
}

const updateStep = `-- name: UpdateStep :exec
UPDATE steps
SET name         = ?,
    completed_at = ?
WHERE id = ?
`

type UpdateStepParams struct {
	Name        string
	CompletedAt sql.NullTime
	ID          string
}

func (q *Queries) UpdateStep(ctx context.Context, arg UpdateStepParams) error {
	_, err := q.db.ExecContext(ctx, updateStep, arg.Name, arg.CompletedAt, arg.ID)
	return err
}

const updateTag = `-- name: UpdateTag :exec
UPDATE tags
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

const updateTask = `-- name: UpdateTask :exec
UPDATE tasks
SET name         = ?,
    content      = ?,
    priority     = ?,
    due_on       = ?,
    completed_at = ?
WHERE id = ?
`

type UpdateTaskParams struct {
	Name        string
	Content     string
	Priority    uint32
	DueOn       sql.NullTime
	CompletedAt sql.NullTime
	ID          string
}

func (q *Queries) UpdateTask(ctx context.Context, arg UpdateTaskParams) error {
	_, err := q.db.ExecContext(ctx, updateTask,
		arg.Name,
		arg.Content,
		arg.Priority,
		arg.DueOn,
		arg.CompletedAt,
		arg.ID,
	)
	return err
}
