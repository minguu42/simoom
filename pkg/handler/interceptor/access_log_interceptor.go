package interceptor

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/pkg/applog"
)

// NewAccessLog はリクエスト毎のアクセスログ/エラーログを表示するインターセプタを返す
func NewAccessLog() connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			resp, err := next(ctx, req)
			if err != nil {
				code := connect.CodeUnknown
				var connectErr *connect.Error
				if errors.As(err, &connectErr) {
					code = connectErr.Code()
				}
				applog.Logger(ctx).LogAttrs(ctx, slog.LevelInfo,
					fmt.Sprintf("%s - %s %s", code, req.HTTPMethod(), req.Spec().Procedure),
					slog.String("stack_trace", fmt.Sprintf("%+v", err)),
				)
				return resp, err
			}
			applog.Logger(ctx).LogAttrs(ctx, slog.LevelInfo, fmt.Sprintf("200 - %s %s", req.HTTPMethod(), req.Spec().Procedure))
			return resp, err
		}
	}
}
