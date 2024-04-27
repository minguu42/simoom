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

func TestTagDeleteRun(t *testing.T) {
	type args struct {
		ctx  context.Context
		opts cmd.TagDeleteOpts
	}
	tests := []struct {
		name    string
		args    args
		wantOut string
	}{
		{
			name: "タグを削除する",
			args: args{
				ctx: context.Background(),
				opts: cmd.TagDeleteOpts{
					Client: &api.ClientMock{DeleteTagFunc: func(_ context.Context, _ *connect.Request[simoompb.DeleteTagRequest]) (*connect.Response[emptypb.Empty], error) {
						return connect.NewResponse(&emptypb.Empty{}), nil
					}},
					ID: "tag-01",
				},
			},
			wantOut: "Tag tag-01 deleted\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			err := cmd.TagDeleteRun(tt.args.ctx, out, tt.args.opts)
			require.NoError(t, err)
			assert.Equal(t, tt.wantOut, out.String())
		})
	}
}
