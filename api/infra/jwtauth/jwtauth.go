// Package jwtauth はJWTを用いて認証器を実装するパッケージ
package jwtauth

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/minguu42/simoom/api/domain/model"
	"github.com/minguu42/simoom/lib/go/clock"
)

// Authenticator は auth.Authenticator を満たすJWT認証器
type Authenticator struct {
	AccessTokenExpiryHour  int
	RefreshTokenExpiryHour int
	AccessTokenSecret      string
	RefreshTokenSecret     string
}

type accessTokenClaims struct {
	jwt.RegisteredClaims
	Name string       `json:"name"`
	ID   model.UserID `json:"id"`
}

type refreshTokenClaims struct {
	jwt.RegisteredClaims
	ID model.UserID `json:"id"`
}

// CreateAccessToken はアクセスシークレットで署名された、ユーザ名とユーザID、有効期限からなるペイロードをエンコードしてアクセストークンを作成する
func (a Authenticator) CreateAccessToken(ctx context.Context, user model.User) (string, error) {
	claims := &accessTokenClaims{
		Name: user.Name,
		ID:   user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(clock.Now(ctx).Add(time.Hour * time.Duration(a.AccessTokenExpiryHour))),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(a.AccessTokenSecret))
	if err != nil {
		return "", fmt.Errorf("failed to create signed JWT: %w", err)
	}
	return t, nil
}

// CreateRefreshToken は与えられたリフレッシュシークレットで署名された、ユーザIDと有効期限からなるペイロードをエンコードしてリフレッシュトークンを作成する
func (a Authenticator) CreateRefreshToken(ctx context.Context, user model.User) (string, error) {
	claimsRefresh := &refreshTokenClaims{
		ID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(clock.Now(ctx).Add(time.Hour * time.Duration(a.RefreshTokenExpiryHour))),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)
	rt, err := token.SignedString([]byte(a.RefreshTokenSecret))
	if err != nil {
		return "", fmt.Errorf("failed to create signed JWT: %w", err)
	}
	return rt, nil
}

// ExtractIDFromAccessToken はアクセストークン作成時にエンコードされたIDをデコードして取り出す
func (a Authenticator) ExtractIDFromAccessToken(token string) (string, error) {
	jwtToken, err := jwt.Parse(token, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", t.Header["alg"])
		}
		return []byte(a.AccessTokenSecret), nil
	})
	if err != nil {
		return "", fmt.Errorf("failed to parse token: %w", err)
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok && !jwtToken.Valid {
		return "", errors.New("invalid token")
	}
	return claims["id"].(string), nil
}

// ExtractIDFromRefreshToken はリフレッシュトークン作成時にエンコードされたIDをデコードして取り出す
func (a Authenticator) ExtractIDFromRefreshToken(tokenString string) (string, error) {
	jwtToken, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", t.Header["alg"])
		}
		return []byte(a.RefreshTokenSecret), nil
	})
	if err != nil {
		return "", fmt.Errorf("failed to parse token: %w", err)
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok && !jwtToken.Valid {
		return "", errors.New("invalid token")
	}
	return claims["id"].(string), nil
}
