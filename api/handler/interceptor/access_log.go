package interceptor

import (
	"context"
	"time"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/api/logging"
)

// AccessLog はリクエスト毎のアクセスログを出力するインターセプタを返す
func AccessLog() connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			start := time.Now()
			resp, err := next(ctx, req)
			end := time.Now()

			logging.Access(ctx, req.Spec().Procedure, end.Sub(start), err)
			return resp, err
		}
	}
}