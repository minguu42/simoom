package interceptor

import (
	"context"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/api/logging"
)

func NewSetContext() connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			ctx = logging.SetLogger(ctx, req.Spec().Procedure)
			return next(ctx, req)
		}
	}
}
