package interceptor

import (
	"context"
	"fmt"
	"log/slog"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/pkg/applog"
)

// NewRecordAccess はリクエスト毎のアクセスログ/エラーログを出力するインターセプタを返す
func NewRecordAccess() connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			resp, err := next(ctx, req)
			if err != nil {
				applog.Logger(ctx).LogAttrs(ctx, slog.LevelInfo,
					fmt.Sprintf("%s %s", req.HTTPMethod(), req.Spec().Procedure),
					slog.String("stack_trace", fmt.Sprintf("%+v", err)),
				)
				return resp, err
			}
			applog.Logger(ctx).LogAttrs(ctx, slog.LevelInfo, fmt.Sprintf("200 - %s %s", req.HTTPMethod(), req.Spec().Procedure))
			return resp, err
		}
	}
}
