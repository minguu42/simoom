package interceptor

import (
	"context"
	"time"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/api/applog"
)

// NewRecordAccess はリクエスト毎のアクセスログを出力するインターセプタを返す
func NewRecordAccess() connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			start := time.Now()
			resp, err := next(ctx, req)
			end := time.Now()
			applog.Access(ctx, req.Spec().Procedure, end.Sub(start), err)
			return resp, err
		}
	}
}
