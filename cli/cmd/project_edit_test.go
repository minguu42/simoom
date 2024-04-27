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

func TestProjectEditRun(t *testing.T) {
	type args struct {
		ctx  context.Context
		opts cmd.ProjectEditOpts
	}
	tests := []struct {
		name    string
		args    args
		wantOut string
	}{
		{
			name: "プロジェクトを編集する",
			args: args{
				ctx: context.Background(),
				opts: cmd.ProjectEditOpts{
					Client: &api.ClientMock{
						UpdateProjectFunc: func(_ context.Context, _ *connect.Request[simoompb.UpdateProjectRequest]) (*connect.Response[simoompb.Project], error) {
							return connect.NewResponse(&simoompb.Project{
								Id:         "project-01",
								Name:       "テストプロジェクト・改",
								Color:      "#123456",
								IsArchived: true,
							}), nil
						},
					},
					ID:         "project-01",
					Name:       "テストプロジェクト・改",
					Color:      "#123456",
					IsArchived: true,
				},
			},
			wantOut: "Project テストプロジェクト・改 (project-01) edited\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			err := cmd.ProjectEditRun(tt.args.ctx, out, tt.args.opts)
			require.NoError(t, err)
			require.Equal(t, tt.wantOut, out.String())
		})
	}
}
