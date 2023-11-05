// Package applog はアプリケーションのログを扱うパッケージである
package applog

import (
	"context"
	"log/slog"
	"os"
)

type loggerKey struct{}

// InitDefault はデフォルトのロガーをセットする
// しかし、アプリケーションのリクエストスコープではデフォルトのロガーは使用せず、コンテキスト中のロガーを使用する。
// この関数はコンテキスト中のロガーのベースとなるロガーを定義している。
func InitDefault() {
	h := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: false,
		Level:     nil,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.MessageKey {
				a.Key = "message"
			}
			return a
		},
	})
	slog.SetDefault(slog.New(h))
}

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
