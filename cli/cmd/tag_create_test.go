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

func TestTagCreateRun(t *testing.T) {
	type args struct {
		ctx  context.Context
		opts cmd.TagCreateOpts
	}
	tests := []struct {
		name    string
		args    args
		wantOut string
	}{
		{
			name: "タグを作成する",
			args: args{
				ctx: context.Background(),
				opts: cmd.TagCreateOpts{
					Client: &api.ClientMock{
						CreateTagFunc: func(_ context.Context, _ *connect.Request[simoompb.CreateTagRequest]) (*connect.Response[simoompb.Tag], error) {
							return connect.NewResponse(&simoompb.Tag{
								Id:   "tag-01",
								Name: "テストタグ1",
							}), nil
						},
					},
					Name: "テストタグ1",
				},
			},
			wantOut: "Tag テストタグ1 (tag-01) created\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			err := cmd.TagCreateRun(tt.args.ctx, out, tt.args.opts)
			require.NoError(t, err)
			assert.Equal(t, tt.wantOut, out.String())
		})
	}
}
