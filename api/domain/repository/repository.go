// Package repository はデータベースにデータを保存する
package repository

import (
	"context"
	"errors"

	"github.com/minguu42/simoom/api/domain/model"
)

//go:generate moq -fmt goimports -out ./repository_mock.go -rm . Repository

var ErrModelNotFound = errors.New("model not found in database")

// Repository は DB による永続化を抽象化する
type Repository interface {
	Transaction(ctx context.Context, fn func(ctxWithTx context.Context) error) error

	CreateProject(ctx context.Context, p model.Project) error
	ListProjectsByUserID(ctx context.Context, userID string, limit, offset uint) ([]model.Project, error)
	GetProjectByID(ctx context.Context, id string) (model.Project, error)
	UpdateProject(ctx context.Context, p model.Project) error
	DeleteProject(ctx context.Context, id string) error

	CreateTask(ctx context.Context, t model.Task) error
	ListTasksByUserID(ctx context.Context, userID string, limit, offset uint, projectID, tagID *string) ([]model.Task, error)
	GetTaskByID(ctx context.Context, id string) (model.Task, error)
	UpdateTask(ctx context.Context, t model.Task) error
	DeleteTask(ctx context.Context, id string) error

	CreateStep(ctx context.Context, s model.Step) error
	GetStepByID(ctx context.Context, id string) (model.Step, error)
	UpdateStep(ctx context.Context, s model.Step) error
	DeleteStep(ctx context.Context, id string) error

	CreateTag(ctx context.Context, t model.Tag) error
	ListTagsByUserID(ctx context.Context, userID string, limit, offset uint) ([]model.Tag, error)
	GetTagByID(ctx context.Context, id string) (model.Tag, error)
	UpdateTag(ctx context.Context, t model.Tag) error
	DeleteTag(ctx context.Context, id string) error

	CreateUser(ctx context.Context, u model.User) error
	GetUserByID(ctx context.Context, id string) (model.User, error)
	GetUserByName(ctx context.Context, name string) (model.User, error)
	GetUserByEmail(ctx context.Context, email string) (model.User, error)
}
