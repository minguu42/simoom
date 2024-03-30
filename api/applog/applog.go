// Package applog はアプリケーションのログを扱うパッケージである
package applog

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/api/apperr"
)

// applicationLogger はリクエストスコープ外でアプリケーションの状況を出力するためのロガー
// リクエストスコープ内ではこのロガーは使用せず、コンテキスト中のリクエストロガーを使用する
var applicationLogger *slog.Logger

func init() {
	applicationLogger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		ReplaceAttr: func(_ []string, a slog.Attr) slog.Attr {
			if a.Key == slog.MessageKey {
				a.Key = "message"
			}
			return a
		},
	}))
}

// LogApplicationEvent はINFOレベルでアプリケーションの状況のログを出力する
func LogApplicationEvent(ctx context.Context, msg string) {
	applicationLogger.Log(ctx, slog.LevelInfo, msg)
}

// LogApplicationError はERRORレベルでアプリケーションエラーのログを出力する
func LogApplicationError(ctx context.Context, msg string) {
	applicationLogger.Log(ctx, slog.LevelError, msg)
}

type loggerKey struct{}

// SetLogger はアプリケーションロガーからリクエストロガーを生成し、コンテキストにリクエストロガーをセットする
func SetLogger(ctx context.Context, method string) context.Context {
	l := applicationLogger.With(slog.String("method", method))
	return context.WithValue(ctx, loggerKey{}, l)
}

// logger はコンテキストからリクエストロガーを取り出す
// コンテキストにリクエストロガーが存在しなければアプリケーションロガーを使用する
func logger(ctx context.Context) *slog.Logger {
	v, ok := ctx.Value(loggerKey{}).(*slog.Logger)
	if ok {
		return v
	}
	return applicationLogger
}

// LogAccess はリクエストが正常に受け付けられた場合のアクセスログを表示する
func LogAccess(ctx context.Context, method string) {
	msg := fmt.Sprintf("ok (0) %s", method)
	logger(ctx).LogAttrs(ctx, slog.LevelInfo, msg)
}

// LogAccessError はリクエストが正常に受け付けられなかった場合のアクセスログを表示する
func LogAccessError(ctx context.Context, method string, err apperr.Error) {
	level := slog.LevelInfo
	code := err.ConnectError().Code()
	if code == connect.CodeUnknown {
		level = slog.LevelError
	}
	msg := fmt.Sprintf("%s (%[1]d) %s", code, method)
	logger(ctx).LogAttrs(ctx, level, msg, slog.String("detail", err.Error()))
}
