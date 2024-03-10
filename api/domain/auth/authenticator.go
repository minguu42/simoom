package auth

//go:generate moq -fmt goimports -out ./authenticator_mock.go -rm . Authenticator

import (
	"context"

	"github.com/minguu42/simoom/api/domain/model"
)

// Authenticator はユーザ認証を抽象化する
type Authenticator interface {
	CreateAccessToken(ctx context.Context, user model.User) (string, error)
	CreateRefreshToken(ctx context.Context, user model.User) (string, error)
	IsAuthorized(tokenString string) (bool, error)
	ExtractIDFromToken(tokenString string) (string, error)
	ExtractIDFromRefreshToken(tokenString string) (string, error)
}
