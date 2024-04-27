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

func TestTaskListRun(t *testing.T) {
	type args struct {
		ctx  context.Context
		opts cmd.TaskListOpts
	}
	tests := []struct {
		name    string
		args    args
		wantOut string
	}{
		{
			name: "タスクを一覧で取得する",
			args: args{
				ctx: context.Background(),
				opts: cmd.TaskListOpts{
					Client: &api.ClientMock{
						ListTasksFunc: func(_ context.Context, _ *connect.Request[simoompb.ListTasksRequest]) (*connect.Response[simoompb.Tasks], error) {
							return connect.NewResponse(&simoompb.Tasks{
								Tasks: []*simoompb.Task{
									{
										Id:        "task-01",
										ProjectId: "project-01",
										Name:      "テストタスク1",
									},
									{
										Id:        "task-02",
										ProjectId: "project-01",
										Name:      "テストタスク2",
									},
								},
								HasNext: true,
							}), nil
						},
					},
					Limit:  2,
					Offset: 0,
				},
			},
			wantOut: "task-01 テストタスク1\ntask-02 テストタスク2\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			err := cmd.TaskListRun(tt.args.ctx, out, tt.args.opts)
			require.NoError(t, err)
			require.Equal(t, tt.wantOut, out.String())
		})
	}
}
