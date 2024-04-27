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
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestStepDeleteRun(t *testing.T) {
	type args struct {
		ctx  context.Context
		opts cmd.StepDeleteOpts
	}
	tests := []struct {
		name    string
		args    args
		wantOut string
	}{
		{
			name: "ステップを削除する",
			args: args{
				ctx: context.Background(),
				opts: cmd.StepDeleteOpts{
					Client: &api.ClientMock{
						DeleteStepFunc: func(_ context.Context, _ *connect.Request[simoompb.DeleteStepRequest]) (*connect.Response[emptypb.Empty], error) {
							return connect.NewResponse(&emptypb.Empty{}), nil
						},
					},
					ID: "step-01",
				},
			},
			wantOut: "Step step-01 deleted\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			err := cmd.StepDeleteRun(tt.args.ctx, out, tt.args.opts)
			require.NoError(t, err)
			assert.Equal(t, tt.wantOut, out.String())
		})
	}
}
