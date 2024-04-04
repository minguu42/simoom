package interceptor

import (
	"context"
	"time"

	"connectrpc.com/connect"
)

// Timeout は一定時間でコンテキストをタイムアウトするインターセプタを返す
func Timeout(timeout time.Duration) connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			ctx, cancel := context.WithTimeout(ctx, timeout)
			defer cancel()

			return next(ctx, req)
		}
	}
}
