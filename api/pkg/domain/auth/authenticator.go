package auth

import "github.com/minguu42/simoom/api/pkg/domain/model"

// Authenticator はユーザ認証を抽象化する
type Authenticator interface {
	CreateAccessToken(user model.User, secret string, expiry int) (string, error)
	CreateRefreshToken(user model.User, secret string, expiry int) (string, error)
	IsAuthorized(tokenString string, secret string) (bool, error)
	ExtractIDFromToken(tokenString string, secret string) (string, error)
}
