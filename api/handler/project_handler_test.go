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

func TestHandler_CreateProject(t *testing.T) {
	t.Run("不正なリクエストはバリデーションではじく", func(t *testing.T) {
		if _, err := th.CreateProject(context.Background(), connect.NewRequest(&simoompb.CreateProjectRequest{})); assert.Error(t, err) {
			var appErr apperr.Error
			require.ErrorAs(t, err, &appErr)
			assert.Equal(t, connect.CodeInvalidArgument, appErr.ConnectError().Code())
		}
	})
}

func TestHandler_ListProjects(t *testing.T) {
	t.Run("不正なリクエストはバリデーションではじく", func(t *testing.T) {
		if _, err := th.ListProjects(context.Background(), connect.NewRequest(&simoompb.ListProjectsRequest{})); assert.Error(t, err) {
			var appErr apperr.Error
			require.ErrorAs(t, err, &appErr)
			assert.Equal(t, connect.CodeInvalidArgument, appErr.ConnectError().Code())
		}
	})
}

func TestHandler_UpdateProject(t *testing.T) {
	t.Run("不正なリクエストはバリデーションではじく", func(t *testing.T) {
		if _, err := th.UpdateProject(context.Background(), connect.NewRequest(&simoompb.UpdateProjectRequest{})); assert.Error(t, err) {
			var appErr apperr.Error
			require.ErrorAs(t, err, &appErr)
			assert.Equal(t, connect.CodeInvalidArgument, appErr.ConnectError().Code())
		}
	})
}

func TestHandler_DeleteProject(t *testing.T) {
	t.Run("不正なリクエストはバリデーションではじく", func(t *testing.T) {
		if _, err := th.DeleteProject(context.Background(), connect.NewRequest(&simoompb.DeleteProjectRequest{})); assert.Error(t, err) {
			var appErr apperr.Error
			require.ErrorAs(t, err, &appErr)
			assert.Equal(t, connect.CodeInvalidArgument, appErr.ConnectError().Code())
		}
	})
}
