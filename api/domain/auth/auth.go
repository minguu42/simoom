// Package auth は認証に関わる処理を扱うパッケージ
package auth

import (
	"context"

	"github.com/minguu42/simoom/api/domain/model"
)

type userKey struct{}

// WithUser は ctx に model.User をセットする
func WithUser(ctx context.Context, u model.User) context.Context {
	return context.WithValue(ctx, userKey{}, u)
}

// User は ctx から model.User を取り出す
// ctx にユーザがセットされていない場合は空の構造体値を返す
func User(ctx context.Context) model.User {
	return ctx.Value(userKey{}).(model.User)
}
