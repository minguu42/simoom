package handler

import (
	"context"
	"testing"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/api/apperr"
	"github.com/minguu42/simoom/lib/go/simoompb/v1"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestHandler_CreateStep(t *testing.T) {
	type args struct {
		ctx context.Context
		req *connect.Request[simoompb.CreateStepRequest]
	}
	tests := []struct {
		name    string
		args    args
		want    *connect.Response[simoompb.Step]
		wantErr apperr.Error
	}{
		{
			name: "不正なリクエストはバリデーションではじく",
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&simoompb.CreateStepRequest{}),
			},
			want:    nil,
			wantErr: apperr.ErrInvalidRequest(nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testHandler.CreateStep(tt.args.ctx, tt.args.req)
			assert.Equal(t, tt.want, resp)

			var appErr apperr.Error
			if tt.wantErr.Error() != "" && assert.ErrorAs(t, err, &appErr) {
				assert.Equal(t, tt.wantErr.ConnectError().Code(), appErr.ConnectError().Code())
			}
		})
	}
}

func TestHandler_UpdateStep(t *testing.T) {
	type args struct {
		ctx context.Context
		req *connect.Request[simoompb.UpdateStepRequest]
	}
	tests := []struct {
		name    string
		args    args
		want    *connect.Response[simoompb.Step]
		wantErr apperr.Error
	}{
		{
			name: "不正なリクエストはバリデーションではじく",
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&simoompb.UpdateStepRequest{}),
			},
			want:    nil,
			wantErr: apperr.ErrInvalidRequest(nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testHandler.UpdateStep(tt.args.ctx, tt.args.req)
			assert.Equal(t, tt.want, resp)

			var appErr apperr.Error
			if tt.wantErr.Error() != "" && assert.ErrorAs(t, err, &appErr) {
				assert.Equal(t, tt.wantErr.ConnectError().Code(), appErr.ConnectError().Code())
			}
		})
	}
}

func TestHandler_DeleteStep(t *testing.T) {
	type args struct {
		ctx context.Context
		req *connect.Request[simoompb.DeleteStepRequest]
	}
	tests := []struct {
		name    string
		args    args
		want    *connect.Response[emptypb.Empty]
		wantErr apperr.Error
	}{
		{
			name: "不正なリクエストはバリデーションではじく",
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&simoompb.DeleteStepRequest{}),
			},
			want:    nil,
			wantErr: apperr.ErrInvalidRequest(nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testHandler.DeleteStep(tt.args.ctx, tt.args.req)
			assert.Equal(t, tt.want, resp)

			var appErr apperr.Error
			if tt.wantErr.Error() != "" && assert.ErrorAs(t, err, &appErr) {
				assert.Equal(t, tt.wantErr.ConnectError().Code(), appErr.ConnectError().Code())
			}
		})
	}
}
