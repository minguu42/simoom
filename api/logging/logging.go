// Package logging はアプリケーションのロギングを行う関数を提供する
package logging

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/api/apperr"
)

// applicationLogger はリクエストスコープ外で使用するアプリケーションのデフォルトロガー
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

type loggerKey struct{}

// logger はコンテキストからリクエストロガーを取り出す
// コンテキストにリクエストロガーが存在しなければアプリケーションロガーを使用する
func logger(ctx context.Context) *slog.Logger {
	v, ok := ctx.Value(loggerKey{}).(*slog.Logger)
	if ok {
		return v
	}
	return applicationLogger
}

// ContextWithLogger はリクエストロガーを生成し、コンテキストにリクエストロガーをセットする
func ContextWithLogger(ctx context.Context, method string) context.Context {
	l := applicationLogger.With(slog.String("method", method))
	return context.WithValue(ctx, loggerKey{}, l)
}

// Event はイベントログを出力する
func Event(ctx context.Context, msg string) {
	logger(ctx).Log(ctx, slog.LevelInfo, msg)
}

// Error はエラーログを出力する
func Error(ctx context.Context, msg string) {
	logger(ctx).Log(ctx, slog.LevelError, msg)
}

// Access はアクセスログを出力する
func Access(ctx context.Context, method string, executionTime time.Duration, err error) {
	if err == nil { // if NO error
		logger(ctx).LogAttrs(ctx, slog.LevelInfo, fmt.Sprintf("ok (0) %s", method),
			slog.Int64("execution_time", executionTime.Milliseconds()),
		)
		return
	}

	appErr := apperr.NewError(err)
	level := slog.LevelWarn
	if appErr.Code() == connect.CodeUnknown {
		level = slog.LevelError
	}
	logger(ctx).LogAttrs(ctx, level, fmt.Sprintf("%s (%[1]d) %s", appErr.Code(), method),
		slog.Int64("execution_time", executionTime.Milliseconds()),
		slog.String("detail", err.Error()),
	)
}
