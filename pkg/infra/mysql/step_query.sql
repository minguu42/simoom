-- name: CreateStep :exec
INSERT INTO step (id, user_id, task_id, title, completed_at, created_at, updated_at)
VALUES (?, ?, ?, ?, NULL, ?, ?);

-- name: ListStepsByTaskID :many
SELECT * FROM step
WHERE task_id = ?
ORDER BY created_at;

-- name: GetStepByID :one
SELECT * FROM step
WHERE id = ?;

-- name: UpdateStep :exec
UPDATE step
SET title        = ?,
    completed_at = ?
WHERE id = ?;

-- name: DeleteStep :exec
DELETE FROM step WHERE id = ?;