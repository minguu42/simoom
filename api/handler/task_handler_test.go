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

func TestHandler_CreateTask(t *testing.T) {
	t.Run("不正なリクエストはバリデーションではじく", func(t *testing.T) {
		if _, err := th.CreateTask(context.Background(), connect.NewRequest(&simoompb.CreateTaskRequest{})); assert.Error(t, err) {
			var appErr apperr.Error
			require.ErrorAs(t, err, &appErr)
			assert.Equal(t, connect.CodeInvalidArgument, appErr.ConnectError().Code())
		}
	})
}

func TestHandler_ListTasks(t *testing.T) {
	t.Run("不正なリクエストはバリデーションではじく", func(t *testing.T) {
		if _, err := th.ListTasks(context.Background(), connect.NewRequest(&simoompb.ListTasksRequest{})); assert.Error(t, err) {
			var appErr apperr.Error
			require.ErrorAs(t, err, &appErr)
			assert.Equal(t, connect.CodeInvalidArgument, appErr.ConnectError().Code())
		}
	})
}

func TestHandler_UpdateTask(t *testing.T) {
	t.Run("不正なリクエストはバリデーションではじく", func(t *testing.T) {
		if _, err := th.UpdateTask(context.Background(), connect.NewRequest(&simoompb.UpdateTaskRequest{})); assert.Error(t, err) {
			var appErr apperr.Error
			require.ErrorAs(t, err, &appErr)
			assert.Equal(t, connect.CodeInvalidArgument, appErr.ConnectError().Code())
		}
	})
}

func TestHandler_DeleteTask(t *testing.T) {
	t.Run("不正なリクエストはバリデーションではじく", func(t *testing.T) {
		if _, err := th.DeleteTask(context.Background(), connect.NewRequest(&simoompb.DeleteTaskRequest{})); assert.Error(t, err) {
			var appErr apperr.Error
			require.ErrorAs(t, err, &appErr)
			assert.Equal(t, connect.CodeInvalidArgument, appErr.ConnectError().Code())
		}
	})
}
