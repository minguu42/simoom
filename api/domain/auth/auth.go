// Package auth は認証に関わる処理を扱うパッケージ
package auth

import (
	"context"
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
