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

func TestStepCreateRun(t *testing.T) {
	type args struct {
		ctx  context.Context
		opts cmd.StepCreateOpts
	}
	tests := []struct {
		name    string
		args    args
		wantOut string
	}{
		{
			name: "ステップを作成する",
			args: args{
				ctx: context.Background(),
				opts: cmd.StepCreateOpts{
					Client: &api.ClientMock{
						CreateStepFunc: func(_ context.Context, _ *connect.Request[simoompb.CreateStepRequest]) (*connect.Response[simoompb.Step], error) {
							return connect.NewResponse(&simoompb.Step{
								Id:     "step-01",
								TaskId: "task-01",
								Name:   "テストステップ1",
							}), nil
						},
					},
					TaskID: "step-01",
					Name:   "テストステップ1",
				},
			},
			wantOut: "Step テストステップ1 (step-01) created\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			err := cmd.StepCreateRun(tt.args.ctx, out, tt.args.opts)
			require.NoError(t, err)
			assert.Equal(t, tt.wantOut, out.String())
		})
	}
}
