// Package applog はアプリケーションのログを扱うパッケージである
package applog

import (
	"context"
	"errors"
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

// Event はINFOレベルでイベントのログを出力する
func Event(ctx context.Context, msg string) {
	applicationLogger.Log(ctx, slog.LevelInfo, msg)
}

// Error はERRORレベルでエラーログを出力する
func Error(ctx context.Context, msg string) {
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

// Access はアクセスログを出力する
func Access(ctx context.Context, method string, err error) {
	if err == nil { // if NO error
		logger(ctx).LogAttrs(ctx, slog.LevelInfo, fmt.Sprintf("ok (0) %s", method))
		return
	}

	var appErr apperr.Error
	if !errors.As(err, &appErr) {
		appErr = apperr.ErrUnknown(err)
	}
	level := slog.LevelInfo
	if appErr.Code() == connect.CodeUnknown {
		level = slog.LevelError
	}
	msg := fmt.Sprintf("%s (%[1]d) %s", appErr.Code(), method)
	logger(ctx).LogAttrs(ctx, level, msg, slog.String("detail", err.Error()))
}
