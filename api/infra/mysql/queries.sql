-- name: CreateProject :exec
INSERT INTO projects (id, user_id, name, color, is_archived)
VALUES (?, ?, ?, ?, ?);

-- name: ListProjectsByUserID :many
SELECT *
FROM projects
WHERE user_id = ?
ORDER BY created_at DESC
LIMIT ? OFFSET ?;

-- name: GetProjectByID :one
SELECT *
FROM projects
WHERE id = ?;

-- name: UpdateProject :exec
UPDATE projects
SET name        = ?,
    color       = ?,
    is_archived = ?
WHERE id = ?;

-- name: DeleteProject :exec
DELETE
FROM projects
WHERE id = ?;

-- name: CreateStep :exec
INSERT INTO steps (id, user_id, task_id, name, completed_at)
VALUES (?, ?, ?, ?, NULL);

-- name: ListStepsByTaskID :many
SELECT *
FROM steps
WHERE task_id = ?
ORDER BY created_at;

-- name: GetStepByID :one
SELECT *
FROM steps
WHERE id = ?;

-- name: UpdateStep :exec
UPDATE steps
SET name         = ?,
    completed_at = ?
WHERE id = ?;

-- name: DeleteStep :exec
DELETE
FROM steps
WHERE id = ?;

-- name: CreateTag :exec
INSERT INTO tags (id, user_id, name)
VALUES (?, ?, ?);

-- name: ListTagsByUserID :many
SELECT *
FROM tags
WHERE user_id = ?
ORDER BY created_at
LIMIT ? OFFSET ?;

-- name: ListTagsByTaskID :many
SELECT t.*
FROM tags AS t
    INNER JOIN tasks_tags AS tt ON t.id = tt.tag_id
WHERE tt.task_id = ?;

-- name: GetTagByID :one
SELECT *
FROM tags
WHERE id = ?;

-- name: UpdateTag :exec
UPDATE tags
SET name = ?
WHERE id = ?;

-- name: DeleteTag :exec
DELETE
FROM tags
WHERE id = ?;

-- name: CreateTask :exec
INSERT INTO tasks (id, user_id, project_id, name, content, priority, due_on, completed_at)
VALUES (?, ?, ?, ?, '', ?, NULL, NULL);

-- name: ListTasksByUserID :many
SELECT *
FROM tasks
WHERE user_id = ?
ORDER BY created_at
LIMIT ? OFFSET ?;

-- name: ListTasksByProjectID :many
SELECT *
FROM tasks
WHERE project_id = ?
ORDER BY created_at
LIMIT ? OFFSET ?;

-- name: ListTasksByTagID :many
SELECT t.*
FROM tasks AS t
    INNER JOIN tasks_tags AS tt ON t.id = tt.task_id
WHERE tt.tag_id = ?
ORDER BY t.created_at
LIMIT ? OFFSET ?;

-- name: ListTasksByProjectIDAndTagID :many
SELECT t.*
FROM tasks AS t
    INNER JOIN tasks_tags AS tt ON t.id = tt.task_id
WHERE t.project_id = ? AND tt.tag_id = ?
ORDER BY t.created_at
LIMIT ? OFFSET ?;

-- name: GetTaskByID :one
SELECT *
FROM tasks
WHERE id = ?;

-- name: UpdateTask :exec
UPDATE tasks
SET name         = ?,
    content      = ?,
    priority     = ?,
    due_on       = ?,
    completed_at = ?
WHERE id = ?;

-- name: DeleteTask :exec
DELETE
FROM tasks
WHERE id = ?;

-- name: CreateUser :exec
INSERT INTO users (id, name, email, password)
VALUES (?, ?, ?, ?);

-- name: GetUserByID :one
SELECT *
FROM users
WHERE id = ?;

-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = ?;
