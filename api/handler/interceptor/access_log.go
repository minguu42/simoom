package interceptor

import (
	"context"
	"strings"
	"time"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/api/logging"
)

// AccessLog はリクエスト毎のアクセスログを出力するインターセプタを返す
func AccessLog() connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			if strings.Split(req.Spec().Procedure, "/")[2] == "CheckHealth" {
				return next(ctx, req)
			}

			start := time.Now()
			resp, err := next(ctx, req)
			logging.Access(ctx, req.Spec().Procedure, time.Since(start), err)
			return resp, err
		}
	}
}
