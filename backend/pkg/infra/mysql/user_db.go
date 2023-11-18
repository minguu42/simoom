package mysql

import (
	"context"
	"database/sql"

	"github.com/cockroachdb/errors"
	"github.com/minguu42/simoom/backend/pkg/domain/model"
	"github.com/minguu42/simoom/backend/pkg/domain/repository"
	"github.com/minguu42/simoom/backend/pkg/infra/mysql/sqlc"
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

func (c *Client) GetUserByID(ctx context.Context, id string) (model.User, error) {
	u, err := sqlc.New(c.db).GetUserByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.User{}, repository.ErrModelNotFound
		}
		return model.User{}, errors.WithStack(err)
	}
	return newModelUser(u), nil
}

func (c *Client) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	u, err := sqlc.New(c.db).GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.User{}, repository.ErrModelNotFound
		}
		return model.User{}, errors.WithStack(err)
	}
	return newModelUser(u), nil
}
