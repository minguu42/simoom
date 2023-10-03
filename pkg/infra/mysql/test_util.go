package mysql

import (
	"context"
	"database/sql"

	"github.com/cockroachdb/errors"
	"github.com/minguu42/simoom/gen/sqlc"
)

func initAllData(ctx context.Context, db *sql.DB) error {
	q := sqlc.New(db)
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

func resetProject(ctx context.Context, db *sql.DB) error {
	q := sqlc.New(db)
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
