// Package applog はアプリケーションのログを扱うパッケージである
package applog

import (
	"context"
	"log/slog"
	"os"
)

// applicationLogger はリクエストスコープ外でアプリケーションの状況を出力するためのロガー
// リクエストスコープ内ではこのロガーは使用せず、コンテキスト中のリクエストロガーを使用する。
var applicationLogger *slog.Logger

// Init はアプリケーションロガーを初期化する
func Init() {
	applicationLogger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
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

// SetLogger は ctx にリクエストスコープのロガーをセットする
func SetLogger(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, loggerKey{}, logger)
}

// Logger は ctx からリクエストスコープのロガーを取り出す
func Logger(ctx context.Context) *slog.Logger {
	v, ok := ctx.Value(loggerKey{}).(*slog.Logger)
	if !ok {
		return slog.Default()
	}
	return v
}
