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
    is_archived = ?,
    updated_at  = ?
WHERE id = ?;

-- name: DeleteProject :exec
DELETE
FROM projects
WHERE id = ?;
