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

func TestHandler_CreateStep(t *testing.T) {
	t.Run("不正なリクエストはバリデーションではじく", func(t *testing.T) {
		if _, err := th.CreateStep(context.Background(), connect.NewRequest(&simoompb.CreateStepRequest{})); assert.Error(t, err) {
			var appErr apperr.Error
			require.ErrorAs(t, err, &appErr)
			assert.Equal(t, connect.CodeInvalidArgument, appErr.ConnectError().Code())
		}
	})
}

func TestHandler_UpdateStep(t *testing.T) {
	t.Run("不正なリクエストはバリデーションではじく", func(t *testing.T) {
		if _, err := th.UpdateStep(context.Background(), connect.NewRequest(&simoompb.UpdateStepRequest{})); assert.Error(t, err) {
			var appErr apperr.Error
			require.ErrorAs(t, err, &appErr)
			assert.Equal(t, connect.CodeInvalidArgument, appErr.ConnectError().Code())
		}
	})
}

func TestHandler_DeleteStep(t *testing.T) {
	t.Run("不正なリクエストはバリデーションではじく", func(t *testing.T) {
		if _, err := th.DeleteStep(context.Background(), connect.NewRequest(&simoompb.DeleteStepRequest{})); assert.Error(t, err) {
			var appErr apperr.Error
			require.ErrorAs(t, err, &appErr)
			assert.Equal(t, connect.CodeInvalidArgument, appErr.ConnectError().Code())
		}
	})
}
