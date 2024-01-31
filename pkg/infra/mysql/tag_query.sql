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
