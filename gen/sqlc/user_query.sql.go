// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: user_query.sql

package sqlc

import (
	"context"
	"time"
)

const createUser = `-- name: CreateUser :exec
INSERT INTO users (id, name, email, password, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?)
`

type CreateUserParams struct {
	ID        string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.db.ExecContext(ctx, createUser,
		arg.ID,
		arg.Name,
		arg.Email,
		arg.Password,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	return err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, name, email, password, created_at, updated_at
FROM users
WHERE email = ?
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}