package interceptor

import (
	"context"
	"log/slog"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/pkg/applog"
)

func NewSetContext() connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			ctx = applog.SetLogger(ctx, slog.Default().With(slog.String("procedure", req.Spec().Procedure)))
			return next(ctx, req)
		}
	}
}
