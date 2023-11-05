// Package auth は認証に関わる処理を扱うパッケージ
package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/minguu42/simoom/pkg/domain/model"
)

type userKey struct{}

// SetUserID は ctx にユーザIDをセットする
func SetUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, userKey{}, userID)
}

// GetUserID は ctx からユーザIDを取り出す
// ctx にユーザIDがセットされていない場合は空文字列を返す
func GetUserID(ctx context.Context) string {
	v, _ := ctx.Value(userKey{}).(string)
	return v
}

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
func CreateAccessToken(user model.User, secret string, expiry int) (string, error) {
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
		return "", errors.WithStack(err)
	}
	return t, nil
}

// CreateRefreshToken は与えられたリフレッシュシークレットで署名された、ユーザIDと有効期限からなるペイロードをエンコードしてリフレッシュトークンを作成する
func CreateRefreshToken(user model.User, secret string, expiry int) (string, error) {
	claimsRefresh := &refreshTokenClaims{
		ID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(expiry))),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)
	rt, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", errors.WithStack(err)
	}
	return rt, nil
}

// IsAuthorized は requestToken が認可されているかどうかをチェックする
func IsAuthorized(tokenString string, secret string) (bool, error) {
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New(fmt.Sprintf("unexpected signing method: %s", token.Header["alg"]))
		}
		return []byte(secret), nil
	})
	if err != nil {
		return false, errors.WithStack(err)
	}
	return true, nil
}

// ExtractIDFromToken はトークン作成時にエンコードされたIDをデコードして取り出す
func ExtractIDFromToken(tokenString string, secret string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New(fmt.Sprintf("unexpected signing method: %s", token.Header["alg"]))
		}
		return []byte(secret), nil
	})
	if err != nil {
		return "", errors.WithStack(err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return "", errors.New("invalid token")
	}
	return claims["id"].(string), nil
}
