package interceptor

import (
	"context"
	"errors"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/api/apperr"
	"github.com/minguu42/simoom/api/applog"
)

// NewRecordAccess はリクエスト毎のアクセスログ/エラーログを出力するインターセプタを返す
func NewRecordAccess() connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			resp, err := next(ctx, req)
			if err == nil {
				applog.LogAccess(ctx, req.Spec().Procedure)
				return resp, nil
			}

			var appErr apperr.Error
			if errors.As(err, &appErr) {
				applog.LogAccessError(ctx, req.Spec().Procedure, appErr)
			} else {
				applog.LogAccessError(ctx, req.Spec().Procedure, apperr.ErrUnknown(err))
			}
			return resp, err
		}
	}
}
