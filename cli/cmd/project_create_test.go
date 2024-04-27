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

func TestProjectCreateRun(t *testing.T) {
	type args struct {
		ctx  context.Context
		opts cmd.ProjectCreateOpts
	}
	tests := []struct {
		name    string
		args    args
		wantOut string
	}{
		{
			name: "プロジェクトを作成する",
			args: args{
				ctx: context.Background(),
				opts: cmd.ProjectCreateOpts{
					Client: &api.ClientMock{
						CreateProjectFunc: func(_ context.Context, _ *connect.Request[simoompb.CreateProjectRequest]) (*connect.Response[simoompb.Project], error) {
							return connect.NewResponse(&simoompb.Project{
								Id:    "project-01",
								Name:  "テストプロジェクト1",
								Color: "#123456",
							}), nil
						},
					},
					Name:  "テストプロジェクト1",
					Color: "#123456",
				},
			},
			wantOut: "Project テストプロジェクト1 (project-01) created\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			err := cmd.ProjectCreateRun(tt.args.ctx, out, tt.args.opts)
			require.NoError(t, err)
			assert.Equal(t, tt.wantOut, out.String())
		})
	}
}
