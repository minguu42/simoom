-- name: CreateTag :exec
INSERT INTO tag (id, user_id, name, created_at, updated_at)
VALUES (?, ?, ?, ?, ?);

-- name: ListTagsByUserID :many
SELECT *
FROM tag
WHERE user_id = ?
ORDER BY created_at
LIMIT ? OFFSET ?;

-- name: ListTagsByTaskID :many
SELECT t.*
FROM tag AS t
       INNER JOIN task_tag AS tt ON t.id = tt.tag_id
WHERE tt.task_id = ?;

-- name: GetTagByID :one
SELECT *
FROM tag
WHERE id = ?;

-- name: UpdateTag :exec
UPDATE tag
SET name = ?
WHERE id = ?;

-- name: DeleteTag :exec
DELETE
FROM tag
WHERE id = ?;
