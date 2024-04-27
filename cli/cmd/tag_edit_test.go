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

func TestTagEditRun(t *testing.T) {
	type args struct {
		ctx  context.Context
		opts cmd.TagEditOpts
	}
	tests := []struct {
		name    string
		args    args
		wantOut string
	}{
		{
			name: "タグを編集する",
			args: args{
				ctx: context.Background(),
				opts: cmd.TagEditOpts{
					Client: &api.ClientMock{
						UpdateTagFunc: func(_ context.Context, _ *connect.Request[simoompb.UpdateTagRequest]) (*connect.Response[simoompb.Tag], error) {
							return connect.NewResponse(&simoompb.Tag{
								Id:   "tag-01",
								Name: "テストタグ・改",
							}), nil
						},
					},
					ID:   "tag-01",
					Name: "テストタグ・改",
				},
			},
			wantOut: "Tag テストタグ・改 (tag-01) edited\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			err := cmd.TagEditRun(tt.args.ctx, out, tt.args.opts)
			require.NoError(t, err)
			require.Equal(t, tt.wantOut, out.String())
		})
	}
}
