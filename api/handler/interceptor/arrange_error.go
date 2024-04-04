package interceptor

import (
	"context"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/api/apperr"
)

// ArrangeError はレスポンス用のエラーを生成するインターセプタを返す
func ArrangeError() connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			resp, err := next(ctx, req)
			if err == nil { // if NO error
				return resp, nil
			}

			return resp, apperr.NewError(err).ConnectError()
		}
	}
}
