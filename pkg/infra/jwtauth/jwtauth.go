// Package jwtauth はJWTを用いて認証器を実装するパッケージ
package jwtauth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/minguu42/simoom/pkg/domain/model"
)

// Authenticator は auth.Authenticator を満たすJWT認証器
type Authenticator struct{}

type accessTokenClaims struct {
	jwt.RegisteredClaims
	Name string `json:"name"`
	ID   string `json:"id"`
}

type refreshTokenClaims struct {
	jwt.RegisteredClaims
	ID string `json:"id"`
}

// CreateAccessToken はアクセスシークレットで署名された、ユーザ名とユーザID、有効期限からなるペイロードをエンコードしてアクセストークンを作成する
func (a Authenticator) CreateAccessToken(user model.User, secret string, expiry int) (string, error) {
	claims := &accessTokenClaims{
		Name: user.Name,
		ID:   user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(expiry))),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", fmt.Errorf("failed to create signed JWT: %w", err)
	}
	return t, nil
}

// CreateRefreshToken は与えられたリフレッシュシークレットで署名された、ユーザIDと有効期限からなるペイロードをエンコードしてリフレッシュトークンを作成する
func (a Authenticator) CreateRefreshToken(user model.User, secret string, expiry int) (string, error) {
	claimsRefresh := &refreshTokenClaims{
		ID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(expiry))),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)
	rt, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", fmt.Errorf("failed to create signed JWT: %w", err)
	}
	return rt, nil
}

// IsAuthorized は requestToken が認可されているかどうかをチェックする
func (a Authenticator) IsAuthorized(tokenString string, secret string) (bool, error) {
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return false, fmt.Errorf("failed to parse token: %w", err)
	}
	return true, nil
}

// ExtractIDFromToken はトークン作成時にエンコードされたIDをデコードして取り出す
func (a Authenticator) ExtractIDFromToken(tokenString string, secret string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return "", fmt.Errorf("failed to parse token: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return "", errors.New("invalid token")
	}
	return claims["id"].(string), nil
}
