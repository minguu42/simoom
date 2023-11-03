-- name: CreateUser :exec
INSERT INTO users (id, name, email, password, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?);

-- name: GetUserByID :one
SELECT *
FROM users
WHERE id = ?;

-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = ?;
