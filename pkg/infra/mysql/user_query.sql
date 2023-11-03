-- name: CreateUser :exec
INSERT INTO users (id, name, email, password, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?);

-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = ?;
