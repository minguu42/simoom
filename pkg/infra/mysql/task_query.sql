-- name: CreateTask :exec
INSERT INTO tasks (id, user_id, project_id, title, content, priority, due_on, completed_at)
VALUES (?, ?, ?, ?, '', ?, NULL, NULL);

-- name: ListTasksByProjectID :many
SELECT *
FROM tasks
WHERE project_id = ?
ORDER BY created_at
LIMIT ? OFFSET ?;

-- name: ListTasksByTagID :many
SELECT t1.*
FROM tasks AS t1
    INNER JOIN tasks_tags AS tt ON t1.id = tt.task_id
WHERE tt.tag_id = ?
ORDER BY t1.created_at
LIMIT ? OFFSET ?;

-- name: GetTaskByID :one
SELECT *
FROM tasks
WHERE id = ?;

-- name: UpdateTask :exec
UPDATE tasks
SET title        = ?,
    content      = ?,
    priority     = ?,
    due_on       = ?,
    completed_at = ?,
    updated_at   = ?
WHERE id = ?;

-- name: DeleteTask :exec
DELETE
FROM tasks
WHERE id = ?;
