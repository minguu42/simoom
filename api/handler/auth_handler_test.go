package handler

import (
	"context"
	"testing"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/api/apperr"
	"github.com/minguu42/simoom/lib/go/simoompb/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandler_SignUp(t *testing.T) {
	t.Run("不正なリクエストはバリデーションではじく", func(t *testing.T) {
		if _, err := th.SignUp(context.Background(), connect.NewRequest(&simoompb.SignUpRequest{})); assert.Error(t, err) {
			var appErr apperr.Error
			require.ErrorAs(t, err, &appErr)
			assert.Equal(t, connect.CodeInvalidArgument, appErr.ConnectError().Code())
		}
	})
}

func TestHandler_SignIn(t *testing.T) {
	t.Run("不正なリクエストはバリデーションではじく", func(t *testing.T) {
		if _, err := th.SignIn(context.Background(), connect.NewRequest(&simoompb.SignInRequest{})); assert.Error(t, err) {
			var appErr apperr.Error
			require.ErrorAs(t, err, &appErr)
			assert.Equal(t, connect.CodeInvalidArgument, appErr.ConnectError().Code())
		}
	})
}

func TestHandler_RefreshToken(t *testing.T) {
	t.Run("不正なリクエストはバリデーションではじく", func(t *testing.T) {
		if _, err := th.RefreshToken(context.Background(), connect.NewRequest(&simoompb.RefreshTokenRequest{})); assert.Error(t, err) {
			var appErr apperr.Error
			require.ErrorAs(t, err, &appErr)
			assert.Equal(t, connect.CodeInvalidArgument, appErr.ConnectError().Code())
		}
	})
}
