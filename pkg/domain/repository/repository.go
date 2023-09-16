// Package repository はデータベースにデータを保存する
package repository

import (
	"context"
	"errors"

	"github.com/minguu42/simoom/pkg/domain/model"
)

//go:generate mockgen -source=$GOFILE -destination=../../../gen/mock/$GOFILE -package=mock

var ErrModelNotFound = errors.New("model not found in database")

// Repository は DB による永続化を抽象化する
type Repository interface {
	CreateProject(ctx context.Context, p model.Project) error
	GetProjectsByUserID(ctx context.Context, userID string, limit, offset uint) ([]model.Project, error)
}
