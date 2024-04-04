package interceptor

import (
	"context"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/api/logging"
)

func WithLogger() connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			ctx = logging.ContextWithLogger(ctx, req.Spec().Procedure)
			return next(ctx, req)
		}
	}
}
