package interceptor

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"strings"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/pkg/domain/auth"
)

// NewAuthenticate はユーザ認証を行うインターセプタを返す
func NewAuthenticate(authenticator auth.Authenticator, secret string) connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			excludedProcedures := []string{"CheckHealth", "SignIn", "SignUp", "RefreshToken"}
			if slices.Contains(excludedProcedures, strings.Split(req.Spec().Procedure, "/")[2]) {
				return next(ctx, req)
			}

			authHeader := req.Header().Get("Authorization")
			t := strings.Split(authHeader, " ")
			if len(t) != 2 {
				return nil, errors.New("the Authorization header should include a value in the form 'Bearer xxx'")
			}
			token := t[1]
			authorized, err := authenticator.IsAuthorized(token, secret)
			if err != nil {
				return nil, fmt.Errorf("failed to authenticate user: %w", err)
			}
			if !authorized {
				return nil, errors.New("authentication failed")
			}
			userID, err := authenticator.ExtractIDFromToken(token, secret)
			if err != nil {
				return nil, fmt.Errorf("failed to extract id from token: %w", err)
			}
			ctx = auth.SetUserID(ctx, userID)
			return next(ctx, req)
		}
	}
}
