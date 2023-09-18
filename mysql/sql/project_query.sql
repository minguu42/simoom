-- name: CreateProject :exec
INSERT INTO projects (id, user_id, name, color, is_archived, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: GetProjectByID :one
SELECT * FROM projects
WHERE id = ?
LIMIT 1;

-- name: ListProjectsByUserID :many
SELECT * FROM projects
WHERE user_id = ?
ORDER BY created_at DESC
LIMIT ? OFFSET ?;

-- name: UpdateProject :exec
UPDATE projects SET name = ?, color = ?, is_archived = ?
WHERE id = ?;

-- name: DeleteProject :exec
DELETE FROM projects WHERE id = ?;
