package auth

//go:generate moq -fmt goimports -out ./authenticator_mock.go -rm . Authenticator

import (
	"context"

	"github.com/minguu42/simoom/api/domain/model"
)

// Authenticator はユーザ認証を抽象化する
type Authenticator interface {
	CreateAccessToken(ctx context.Context, user model.User, secret string, expiry int) (string, error)
	CreateRefreshToken(ctx context.Context, user model.User, secret string, expiry int) (string, error)
	IsAuthorized(tokenString string, secret string) (bool, error)
	ExtractIDFromToken(tokenString string, secret string) (string, error)
}
