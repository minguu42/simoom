package interceptor

import (
	"context"
	"errors"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/api/apperr"
)

// NewArrangeError はレスポンス用のエラーを生成するインターセプタを返す
func NewArrangeError() connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			resp, err := next(ctx, req)
			if err == nil {
				return resp, nil
			}

			var appErr apperr.Error
			if errors.As(err, &appErr) {
				return resp, appErr.ConnectError()
			}
			return resp, apperr.ErrUnknown(err).ConnectError()
		}
	}
}
