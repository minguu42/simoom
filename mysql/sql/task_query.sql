-- name: CreateTask :exec
INSERT INTO tasks (id, project_id, title, content, priority, due_on, completed_at, created_at, updated_at)
VALUES (?, ?, ?, '', ?, NULL, NULL, ?, ?);

-- name: GetTaskByID :one
SELECT * FROM tasks
WHERE id = ?
LIMIT 1;

-- name: ListTasksByProjectID :many
SELECT * FROM tasks
WHERE project_id = ?
ORDER BY created_at DESC
LIMIT ? OFFSET ?;

-- name: UpdateTask :exec
UPDATE tasks SET title = ?, content = ?, priority = ?, due_on = ?, completed_at = ?
WHERE id = ?;

-- name: DeleteTask :exec
DELETE FROM tasks WHERE id = ?;
