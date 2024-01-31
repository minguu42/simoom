-- name: CreateStep :exec
INSERT INTO steps (id, user_id, task_id, title, completed_at)
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
SET title        = ?,
    completed_at = ?
WHERE id = ?;

-- name: DeleteStep :exec
DELETE
FROM steps
WHERE id = ?;
