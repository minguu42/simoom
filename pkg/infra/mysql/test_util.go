package mysql

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/minguu42/simoom/pkg/infra/mysql/sqlc"
)

func InitAllData(ctx context.Context, client *Client) error {
	q := sqlc.New(client.db)
	if err := q.DeleteAllUsers(ctx); err != nil {
		return errors.WithStack(err)
	}

	if err := q.ImportUser(ctx); err != nil {
		return errors.WithStack(err)
	}
	if err := q.ImportProject(ctx); err != nil {
		return errors.WithStack(err)
	}
	if err := q.ImportTag(ctx); err != nil {
		return errors.WithStack(err)
	}
	if err := q.ImportTask(ctx); err != nil {
		return errors.WithStack(err)
	}
	if err := q.ImportTaskTag(ctx); err != nil {
		return errors.WithStack(err)
	}
	if err := q.ImportStep(ctx); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func ResetProject(ctx context.Context, client *Client) error {
	q := sqlc.New(client.db)
	if err := q.DeleteAllProjects(ctx); err != nil {
		return errors.WithStack(err)
	}

	if err := q.ImportProject(ctx); err != nil {
		return errors.WithStack(err)
	}
	if err := q.ImportTask(ctx); err != nil {
		return errors.WithStack(err)
	}
	if err := q.ImportTaskTag(ctx); err != nil {
		return errors.WithStack(err)
	}
	if err := q.ImportStep(ctx); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func ResetStep(ctx context.Context, client *Client) error {
	q := sqlc.New(client.db)
	if err := q.DeleteAllSteps(ctx); err != nil {
		return errors.WithStack(err)
	}
	if err := q.ImportStep(ctx); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func ResetTag(ctx context.Context, client *Client) error {
	q := sqlc.New(client.db)
	if err := q.DeleteAllTags(ctx); err != nil {
		return errors.WithStack(err)
	}

	if err := q.ImportTag(ctx); err != nil {
		return errors.WithStack(err)
	}
	if err := q.ImportTaskTag(ctx); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func ResetTask(ctx context.Context, client *Client) error {
	q := sqlc.New(client.db)
	if err := q.DeleteAllTasks(ctx); err != nil {
		return errors.WithStack(err)
	}

	if err := q.ImportTask(ctx); err != nil {
		return errors.WithStack(err)
	}
	if err := q.ImportStep(ctx); err != nil {
		return errors.WithStack(err)
	}
	if err := q.ImportTaskTag(ctx); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
