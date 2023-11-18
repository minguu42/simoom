package interceptor

import (
	"context"

	"connectrpc.com/connect"
	"github.com/cockroachdb/errors"
	"github.com/minguu42/simoom/pkg/usecase"
)

// NewJudgeError はハンドラから返されたエラーの種類を判定し、適切なConnectエラーに変換するインターセプトを返す
func NewJudgeError() connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			res, err := next(ctx, req)
			if err != nil {
				var connectErr *connect.Error
				switch {
				case errors.As(err, &connectErr):
					return res, err
				case errors.Is(err, usecase.ErrProjectNotFound) ||
					errors.Is(err, usecase.ErrStepNotFound) ||
					errors.Is(err, usecase.ErrTagNotFound) ||
					errors.Is(err, usecase.ErrTaskNotFound):
					return res, connect.NewError(connect.CodeNotFound, err)
				}
				return res, connect.NewError(connect.CodeUnknown, err)
			}
			return res, nil
		}
	}
}
