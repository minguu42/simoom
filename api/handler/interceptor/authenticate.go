package interceptor

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"strings"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/api/apperr"
	"github.com/minguu42/simoom/api/domain/auth"
	"github.com/minguu42/simoom/api/domain/repository"
)

// NewAuthenticate はユーザ認証を行うインターセプタを返す
func NewAuthenticate(authenticator auth.Authenticator, repo repository.Repository) connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			excludedProcedures := []string{"CheckHealth", "SignIn", "SignUp", "RefreshToken"}
			if slices.Contains(excludedProcedures, strings.Split(req.Spec().Procedure, "/")[2]) {
				return next(ctx, req)
			}

			t := strings.Split(req.Header().Get("Authorization"), " ")
			if len(t) != 2 {
				return nil, apperr.ErrInvalidAuthorizationFormat
			}
			token := t[1]

			if authorized, err := authenticator.IsAuthorized(token); !authorized || err != nil {
				return nil, apperr.ErrAuthenticationFailed
			}

			userID, err := authenticator.ExtractIDFromToken(token)
			if err != nil {
				return nil, fmt.Errorf("failed to extract id from token: %w", err)
			}
			u, err := repo.GetUserByID(ctx, userID)
			if err != nil {
				if errors.Is(err, repository.ErrModelNotFound) {
					return nil, apperr.ErrUserNotFound
				}
				return nil, fmt.Errorf("failed to get user: %w", err)
			}

			return next(auth.WithUser(ctx, u), req)
		}
	}
}
