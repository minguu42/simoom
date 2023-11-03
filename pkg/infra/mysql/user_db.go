package mysql

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/minguu42/simoom/gen/sqlc"
	"github.com/minguu42/simoom/pkg/domain/model"
)

func (c *Client) CreateUser(ctx context.Context, u model.User) error {
	if err := sqlc.New(c.db).CreateUser(ctx, sqlc.CreateUserParams{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Password:  u.Password,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
