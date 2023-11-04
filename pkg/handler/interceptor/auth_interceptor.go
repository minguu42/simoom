package interceptor

import (
	"context"
	"slices"
	"strings"

	"connectrpc.com/connect"
	"github.com/cockroachdb/errors"
	"github.com/minguu42/simoom/pkg/domain/auth"
)

// NewAuth はユーザ認証を行うインターセプタを返す
// secret はアクセスシークレットを受け取る
func NewAuth(secret string) connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			excludedProcedures := []string{"CheckHealth", "SignIn", "SignUp", "RefreshAccessToken"}
			if slices.Contains(excludedProcedures, strings.Split(req.Spec().Procedure, "/")[2]) {
				return next(ctx, req)
			}

			authHeader := req.Header().Get("Authorization")
			t := strings.Split(authHeader, " ")
			if len(t) != 2 {
				return nil, errors.New("the Authorization header should include a value in the form 'Bearer xxx'")
			}
			authToken := t[1]
			authorized, err := auth.IsAuthorized(authToken, secret)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			if !authorized {
				return nil, errors.New("authentication failed")
			}
			userID, err := auth.ExtractIDFromToken(authToken, secret)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			ctx = auth.SetUserID(ctx, userID)
			return next(ctx, req)
		}
	}
}
