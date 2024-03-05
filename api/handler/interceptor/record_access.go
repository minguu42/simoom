package interceptor

import (
	"context"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/api/applog"
)

// NewRecordAccess はリクエスト毎のアクセスログ/エラーログを出力するインターセプタを返す
func NewRecordAccess() connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			resp, err := next(ctx, req)
			if err != nil {
				applog.LogAccessError(ctx, connectError(err).Code(), req.Spec().Procedure, err)
				return resp, err
			}
			applog.LogAccess(ctx, req.Spec().Procedure)
			return resp, nil
		}
	}
}
