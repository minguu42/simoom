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

func TestProjectDeleteRun(t *testing.T) {
	type args struct {
		ctx  context.Context
		opts cmd.ProjectDeleteOpts
	}
	tests := []struct {
		name    string
		args    args
		wantOut string
	}{
		{
			name: "プロジェクトを削除する",
			args: args{
				ctx: context.Background(),
				opts: cmd.ProjectDeleteOpts{
					Client: &api.ClientMock{
						DeleteProjectFunc: func(_ context.Context, _ *connect.Request[simoompb.DeleteProjectRequest]) (*connect.Response[emptypb.Empty], error) {
							return connect.NewResponse(&emptypb.Empty{}), nil
						},
					},
					ID: "project-01",
				},
			},
			wantOut: "Project project-01 deleted\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			err := cmd.ProjectDeleteRun(tt.args.ctx, out, tt.args.opts)
			require.NoError(t, err)
			assert.Equal(t, tt.wantOut, out.String())
		})
	}
}
