-- name: CreateUser :exec
INSERT INTO users (id, name, email, password)
VALUES (?, ?, ?, ?);

-- name: GetUserByID :one
SELECT *
FROM users
WHERE id = ?;

-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = ?;
