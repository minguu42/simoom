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

func TestTagListRun(t *testing.T) {
	type args struct {
		ctx  context.Context
		opts cmd.TagListOpts
	}
	tests := []struct {
		name    string
		args    args
		wantOut string
	}{
		{
			name: "タグを一覧で取得する",
			args: args{
				ctx: context.Background(),
				opts: cmd.TagListOpts{
					Client: &api.ClientMock{
						ListTagsFunc: func(_ context.Context, _ *connect.Request[simoompb.ListTagsRequest]) (*connect.Response[simoompb.Tags], error) {
							return connect.NewResponse(&simoompb.Tags{
								Tags: []*simoompb.Tag{
									{
										Id:   "tag-01",
										Name: "テストタグ1",
									},
									{
										Id:   "tag-02",
										Name: "テストタグ2",
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
			wantOut: "tag-01 テストタグ1\ntag-02 テストタグ2\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			err := cmd.TagListRun(tt.args.ctx, out, tt.args.opts)
			require.NoError(t, err)
			require.Equal(t, tt.wantOut, out.String())
		})
	}
}
