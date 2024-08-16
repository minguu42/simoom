package mysql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/minguu42/simoom/api/adapter/mysql/sqlc"
	"github.com/minguu42/simoom/api/domain/model"
	"github.com/minguu42/simoom/api/domain/repository"
)

func newModelUser(u sqlc.User) model.User {
	return model.User{
		ID:       model.UserID(u.ID),
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	}
}

func (c *Client) CreateUser(ctx context.Context, u model.User) error {
	if err := c.queries(ctx).CreateUser(ctx, sqlc.CreateUserParams{
		ID:       string(u.ID),
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	}); err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

func (c *Client) GetUserByID(ctx context.Context, id model.UserID) (model.User, error) {
	u, err := c.queries(ctx).GetUserByID(ctx, string(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.User{}, repository.ErrModelNotFound
		}
		return model.User{}, fmt.Errorf("failed to get user: %w", err)
	}
	return newModelUser(u), nil
}

func (c *Client) GetUserByName(ctx context.Context, name string) (model.User, error) {
	u, err := c.queries(ctx).GetUserByName(ctx, name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.User{}, repository.ErrModelNotFound
		}
		return model.User{}, fmt.Errorf("failed to get user: %w", err)
	}
	return newModelUser(u), nil
}

func (c *Client) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	u, err := c.queries(ctx).GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.User{}, repository.ErrModelNotFound
		}
		return model.User{}, fmt.Errorf("failed to get user: %w", err)
	}
	return newModelUser(u), nil
}
