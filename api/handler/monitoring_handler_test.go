package handler

import (
	"context"
	"testing"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/lib/go/simoompb/v1"
	"github.com/stretchr/testify/assert"
)

func TestHandler_CheckHealth(t *testing.T) {
	th := handler{}
	type args struct {
		ctx context.Context
		req *connect.Request[simoompb.CheckHealthRequest]
	}
	tests := []struct {
		name    string
		args    args
		want    *connect.Response[simoompb.CheckHealthResponse]
		wantErr error
	}{
		{
			name: "ヘルスステータスを取得する",
			args: args{},
			want: connect.NewResponse(&simoompb.CheckHealthResponse{
				Revision: "xxxxxxx",
			}),
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := th.CheckHealth(tt.args.ctx, tt.args.req)
			assert.Equal(t, tt.want, resp)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
