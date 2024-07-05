package logging

import (
	"context"
	"log/slog"
	"os"
	"time"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/api/apperr"
)

// applicationLogger はリクエストスコープ外で使用するアプリケーションのデフォルトロガー
var applicationLogger *slog.Logger

func init() {
	w := os.Stdout
	opts := &slog.HandlerOptions{
		ReplaceAttr: func(_ []string, a slog.Attr) slog.Attr {
			if a.Key == slog.MessageKey {
				a.Key = "message"
			}
			return a
		},
	}
	if os.Getenv("USE_DEBUG_LOGGER") == "true" {
		applicationLogger = slog.New(NewJSONIndentHandler(w, opts))
	} else {
		applicationLogger = slog.New(slog.NewJSONHandler(w, opts))
	}
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

type AccessFields struct {
	// レスポンス関連のフィールド
	ExecutionTime time.Duration
	Err           error

	// リクエスト関連のフィールド
	HTTPMethod    string
	Service       string
	Method        string
	ContentLength int
	RemoteAddr    string
}

// Access はアクセスログを出力する
func Access(ctx context.Context, fields AccessFields) {
	message := "Request accepted"
	executionTimeAttr := slog.Int64("execution_time", fields.ExecutionTime.Milliseconds())
	requestAttr := slog.Group("request",
		slog.String("http_method", fields.HTTPMethod),
		slog.String("service", fields.Service),
		slog.String("method", fields.Method),
		slog.Int("content_length", fields.ContentLength),
		slog.String("remote_address", fields.RemoteAddr),
	)

	if fields.Err == nil { // if NO error
		logger(ctx).LogAttrs(ctx, slog.LevelInfo, message,
			executionTimeAttr,
			requestAttr,
		)
		return
	}

	appErr := apperr.NewError(fields.Err)
	level := slog.LevelWarn
	if appErr.Code() == connect.CodeUnknown {
		level = slog.LevelError
	}
	logger(ctx).LogAttrs(ctx, level, message,
		executionTimeAttr,
		requestAttr,
		slog.Int("error_code", int(appErr.Code())),
		slog.String("error_text", appErr.Code().String()),
		slog.String("error_message", appErr.Error()),
	)
}

// SQL はクエリログを出力する
func SQL(ctx context.Context, query string, args []any, executionTime time.Duration, err error) {
	level := slog.LevelInfo
	message := "Query executed"
	queryAttr := slog.String("query", query)
	argsAttr := slog.Any("args", args)
	executionTimeAttr := slog.Int64("execution_time", executionTime.Milliseconds())
	if err != nil {
		logger(ctx).LogAttrs(ctx, level, message,
			queryAttr,
			argsAttr,
			executionTimeAttr,
			slog.String("error_message", err.Error()),
		)
	} else {
		logger(ctx).LogAttrs(ctx, level, message,
			queryAttr,
			argsAttr,
			executionTimeAttr,
		)
	}
}
