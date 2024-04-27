package cmd_test

import (
	"bytes"
	"context"
	"testing"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/cli/api"
	"github.com/minguu42/simoom/cli/cmd"
	"github.com/minguu42/simoom/lib/go/simoompb/v1"
	"github.com/stretchr/testify/require"
)

func TestTaskEditRun(t *testing.T) {
	type args struct {
		ctx  context.Context
		opts cmd.TaskEditOpts
	}
	tests := []struct {
		name    string
		args    args
		wantOut string
	}{
		{
			name: "タスクを編集する",
			args: args{
				ctx: context.Background(),
				opts: cmd.TaskEditOpts{
					Client: &api.ClientMock{
						UpdateTaskFunc: func(_ context.Context, _ *connect.Request[simoompb.UpdateTaskRequest]) (*connect.Response[simoompb.Task], error) {
							return connect.NewResponse(&simoompb.Task{
								Id:        "task-01",
								ProjectId: "project-01",
								Name:      "テストタスク・改",
							}), nil
						},
					},
					ID:   "task-01",
					Name: "テストタスク・改",
				},
			},
			wantOut: "Task テストタスク・改 (task-01) edited\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			err := cmd.TaskEditRun(tt.args.ctx, out, tt.args.opts)
			require.NoError(t, err)
			require.Equal(t, tt.wantOut, out.String())
		})
	}
}
