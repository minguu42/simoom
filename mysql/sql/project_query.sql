-- name: CreateProject :exec
INSERT INTO project (id, user_id, name, color, is_archived, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: ListProjectsByUserID :many
SELECT * FROM project
WHERE user_id = ?
ORDER BY created_at DESC
LIMIT ? OFFSET ?;

-- name: GetProjectByID :one
SELECT * FROM project
WHERE id = ?;

-- name: UpdateProject :exec
UPDATE project
SET name        = ?,
    color       = ?,
    is_archived = ?
WHERE id = ?;

-- name: DeleteProject :exec
DELETE FROM project WHERE id = ?;
