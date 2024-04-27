package cmd_test

import (
	"bytes"
	"context"
	"testing"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/cli/api"
	"github.com/minguu42/simoom/cli/cmd"
	"github.com/minguu42/simoom/lib/go/simoompb/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTaskCreateRun(t *testing.T) {
	type args struct {
		ctx  context.Context
		opts cmd.TaskCreateOpts
	}
	tests := []struct {
		name    string
		args    args
		wantOut string
	}{
		{
			name: "タスクを作成する",
			args: args{
				ctx: context.Background(),
				opts: cmd.TaskCreateOpts{
					Client: &api.ClientMock{
						CreateTaskFunc: func(_ context.Context, _ *connect.Request[simoompb.CreateTaskRequest]) (*connect.Response[simoompb.Task], error) {
							return connect.NewResponse(&simoompb.Task{
								Id:        "task-01",
								ProjectId: "project-01",
								Name:      "テストタスク1",
								Priority:  1,
							}), nil
						},
					},
					ProjectID: "project-01",
					Name:      "テストタスク1",
					Priority:  1,
				},
			},
			wantOut: "Task テストタスク1 (task-01) created\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			err := cmd.TaskCreateRun(tt.args.ctx, out, tt.args.opts)
			require.NoError(t, err)
			assert.Equal(t, tt.wantOut, out.String())
		})
	}
}
