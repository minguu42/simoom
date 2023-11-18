// Package repository はデータベースにデータを保存する
package repository

import (
	"context"

	"github.com/cockroachdb/errors"
	model2 "github.com/minguu42/simoom/backend/pkg/domain/model"
)

var ErrModelNotFound = errors.New("model not found in database")

// Repository は DB による永続化を抽象化する
type Repository interface {
	CreateProject(ctx context.Context, p model2.Project) error
	ListProjectsByUserID(ctx context.Context, userID string, limit, offset uint) ([]model2.Project, error)
	GetProjectByID(ctx context.Context, id string) (model2.Project, error)
	UpdateProject(ctx context.Context, p model2.Project) error
	DeleteProject(ctx context.Context, id string) error

	CreateTask(ctx context.Context, t model2.Task) error
	ListTasksByProjectID(ctx context.Context, projectID string, limit, offset uint) ([]model2.Task, error)
	ListTasksByTagID(ctx context.Context, tagID string, limit, offset uint) ([]model2.Task, error)
	GetTaskByID(ctx context.Context, id string) (model2.Task, error)
	UpdateTask(ctx context.Context, t model2.Task) error
	DeleteTask(ctx context.Context, id string) error

	CreateStep(ctx context.Context, s model2.Step) error
	GetStepByID(ctx context.Context, id string) (model2.Step, error)
	UpdateStep(ctx context.Context, s model2.Step) error
	DeleteStep(ctx context.Context, id string) error

	CreateTag(ctx context.Context, t model2.Tag) error
	ListTagsByUserID(ctx context.Context, userID string, limit, offset uint) ([]model2.Tag, error)
	GetTagByID(ctx context.Context, id string) (model2.Tag, error)
	UpdateTag(ctx context.Context, t model2.Tag) error
	DeleteTag(ctx context.Context, id string) error

	CreateUser(ctx context.Context, u model2.User) error
	GetUserByID(ctx context.Context, id string) (model2.User, error)
	GetUserByEmail(ctx context.Context, email string) (model2.User, error)
}
