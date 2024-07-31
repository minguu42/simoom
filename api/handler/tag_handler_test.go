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

func TestHandler_CreateTag(t *testing.T) {
	t.Run("不正なリクエストはバリデーションではじく", func(t *testing.T) {
		if _, err := th.CreateTag(context.Background(), connect.NewRequest(&simoompb.CreateTagRequest{})); assert.Error(t, err) {
			var appErr apperr.Error
			require.ErrorAs(t, err, &appErr)
			assert.Equal(t, connect.CodeInvalidArgument, appErr.ConnectError().Code())
		}
	})
}

func TestHandler_ListTags(t *testing.T) {
	t.Run("不正なリクエストはバリデーションではじく", func(t *testing.T) {
		if _, err := th.ListTags(context.Background(), connect.NewRequest(&simoompb.ListTagsRequest{})); assert.Error(t, err) {
			var appErr apperr.Error
			require.ErrorAs(t, err, &appErr)
			assert.Equal(t, connect.CodeInvalidArgument, appErr.ConnectError().Code())
		}
	})
}

func TestHandler_UpdateTag(t *testing.T) {
	t.Run("不正なリクエストはバリデーションではじく", func(t *testing.T) {
		if _, err := th.UpdateTag(context.Background(), connect.NewRequest(&simoompb.UpdateTagRequest{})); assert.Error(t, err) {
			var appErr apperr.Error
			require.ErrorAs(t, err, &appErr)
			assert.Equal(t, connect.CodeInvalidArgument, appErr.ConnectError().Code())
		}
	})
}

func TestHandler_DeleteTag(t *testing.T) {
	t.Run("不正なリクエストはバリデーションではじく", func(t *testing.T) {
		if _, err := th.DeleteTag(context.Background(), connect.NewRequest(&simoompb.DeleteTagRequest{})); assert.Error(t, err) {
			var appErr apperr.Error
			require.ErrorAs(t, err, &appErr)
			assert.Equal(t, connect.CodeInvalidArgument, appErr.ConnectError().Code())
		}
	})
}
