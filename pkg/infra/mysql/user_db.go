package mysql

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/minguu42/simoom/gen/sqlc"
	"github.com/minguu42/simoom/pkg/domain/model"
)

func newModelUser(u sqlc.User) model.User {
	return model.User{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Password:  u.Password,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

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

func (c *Client) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	u, err := sqlc.New(c.db).GetUserByEmail(ctx, email)
	if err != nil {
		return model.User{}, errors.WithStack(err)
	}
	return newModelUser(u), nil
}
