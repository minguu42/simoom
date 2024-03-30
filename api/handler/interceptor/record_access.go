package interceptor

import (
	"context"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/api/applog"
)

// NewRecordAccess はリクエスト毎のアクセスログを出力するインターセプタを返す
func NewRecordAccess() connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			resp, err := next(ctx, req)
			applog.Access(ctx, req.Spec().Procedure, err)
			return resp, err
		}
	}
}
