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

func TestProjectListRun(t *testing.T) {
	type args struct {
		ctx  context.Context
		opts cmd.ProjectListOpts
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
				opts: cmd.ProjectListOpts{
					Client: &api.ClientMock{
						ListProjectsFunc: func(_ context.Context, _ *connect.Request[simoompb.ListProjectsRequest]) (*connect.Response[simoompb.Projects], error) {
							return connect.NewResponse(&simoompb.Projects{
								Projects: []*simoompb.Project{
									{
										Id:         "project-01",
										Name:       "テストプロジェクト1",
										Color:      "#000000",
										IsArchived: true,
									},
									{
										Id:         "project-02",
										Name:       "テストプロジェクト2",
										Color:      "#123456",
										IsArchived: false,
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
			wantOut: "project-01 テストプロジェクト1\nproject-02 テストプロジェクト2\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			err := cmd.ProjectListRun(tt.args.ctx, out, tt.args.opts)
			require.NoError(t, err)
			require.Equal(t, tt.wantOut, out.String())
		})
	}
}
