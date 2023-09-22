-- name: CreateTask :exec
INSERT INTO task (id, user_id, project_id, title, content, priority, due_on, completed_at, created_at, updated_at)
VALUES (?, ?, ?, ?, '', ?, NULL, NULL, ?, ?);

-- name: ListTasksByProjectID :many
SELECT *
FROM task
WHERE project_id = ?
ORDER BY created_at DESC
LIMIT ? OFFSET ?;

-- name: GetTaskByID :one
SELECT t.*,
       s.id           AS step_id,
       s.user_id      AS step_user_id,
       s.task_id      AS step_task_id,
       s.title        AS step_title,
       s.completed_at AS step_completed_at,
       s.created_at   AS step_created_at,
       s.updated_at   AS step_updated_at
FROM task AS t
       INNER JOIN step AS s ON t.id = s.task_id
WHERE t.id = ?;

-- name: UpdateTask :exec
UPDATE task
SET title        = ?,
    content      = ?,
    priority     = ?,
    due_on       = ?,
    completed_at = ?
WHERE id = ?;

-- name: DeleteTask :exec
DELETE
FROM task
WHERE id = ?;
