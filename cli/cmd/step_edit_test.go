package cmd_test

import (
	"bytes"
	"context"
	"testing"
	"time"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/cli/api"
	"github.com/minguu42/simoom/cli/cmd"
	"github.com/minguu42/simoom/lib/go/simoompb/v1"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestStepEditRun(t *testing.T) {
	type args struct {
		ctx  context.Context
		opts cmd.StepEditOpts
	}
	tests := []struct {
		name    string
		args    args
		wantOut string
	}{
		{
			name: "ステップを編集する",
			args: args{
				ctx: context.Background(),
				opts: cmd.StepEditOpts{
					Client: &api.ClientMock{
						UpdateStepFunc: func(_ context.Context, _ *connect.Request[simoompb.UpdateStepRequest]) (*connect.Response[simoompb.Step], error) {
							return connect.NewResponse(&simoompb.Step{
								Id:          "step-01",
								TaskId:      "task-01",
								Name:        "テストステップ・改",
								CompletedAt: timestamppb.New(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)),
							}), nil
						},
					},
					ID:        "step-01",
					Name:      "テストステップ・改",
					Completed: true,
				},
			},
			wantOut: "Step テストステップ・改 (step-01) edited\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			err := cmd.StepEditRun(tt.args.ctx, out, tt.args.opts)
			require.NoError(t, err)
			require.Equal(t, tt.wantOut, out.String())
		})
	}
}
