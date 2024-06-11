package interceptor

import (
	"context"
	"strconv"
	"strings"
	"time"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/api/logging"
)

// AccessLog はリクエスト毎のアクセスログを出力するインターセプタを返す
func AccessLog() connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			method := strings.Split(req.Spec().Procedure, "/")[2]
			if method == "CheckHealth" {
				return next(ctx, req)
			}

			start := time.Now()
			resp, err := next(ctx, req)

			contentLength, _ := strconv.Atoi(req.Header().Get("Content-Length"))
			fields := logging.AccessFields{
				ExecutionTime: time.Since(start),
				Err:           err,
				HTTPMethod:    req.HTTPMethod(),
				Service:       strings.Split(req.Spec().Procedure, "/")[1],
				Method:        method,
				ContentLength: contentLength,
				RemoteAddr:    req.Peer().Addr,
			}
			logging.Access(ctx, fields)
			return resp, err
		}
	}
}
