package handler

import (
	"context"
	"testing"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/api/apperr"
	"github.com/minguu42/simoom/lib/go/simoompb/v1"
	"github.com/stretchr/testify/assert"
)

func TestHandler_SignUp(t *testing.T) {
	type args struct {
		ctx context.Context
		req *connect.Request[simoompb.SignUpRequest]
	}
	tests := []struct {
		name    string
		args    args
		want    *connect.Response[simoompb.SignUpResponse]
		wantErr apperr.Error
	}{
		{
			name: "不正なリクエストはバリデーションではじく",
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&simoompb.SignUpRequest{}),
			},
			want:    nil,
			wantErr: apperr.ErrInvalidRequest(nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testHandler.SignUp(tt.args.ctx, tt.args.req)
			assert.Equal(t, tt.want, resp)

			var appErr apperr.Error
			if tt.wantErr.Error() != "" && assert.ErrorAs(t, err, &appErr) {
				assert.Equal(t, tt.wantErr.ConnectError().Code(), appErr.ConnectError().Code())
			}
		})
	}
}

func TestHandler_SignIn(t *testing.T) {
	type args struct {
		ctx context.Context
		req *connect.Request[simoompb.SignInRequest]
	}
	tests := []struct {
		name    string
		args    args
		want    *connect.Response[simoompb.SignInResponse]
		wantErr apperr.Error
	}{
		{
			name: "不正なリクエストはバリデーションではじく",
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&simoompb.SignInRequest{}),
			},
			want:    nil,
			wantErr: apperr.ErrInvalidRequest(nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testHandler.SignIn(tt.args.ctx, tt.args.req)
			assert.Equal(t, tt.want, resp)

			var appErr apperr.Error
			if tt.wantErr.Error() != "" && assert.ErrorAs(t, err, &appErr) {
				assert.Equal(t, tt.wantErr.ConnectError().Code(), appErr.ConnectError().Code())
			}
		})
	}
}

func TestHandler_RefreshToken(t *testing.T) {
	type args struct {
		ctx context.Context
		req *connect.Request[simoompb.RefreshTokenRequest]
	}
	tests := []struct {
		name    string
		args    args
		want    *connect.Response[simoompb.RefreshTokenResponse]
		wantErr apperr.Error
	}{
		{
			name: "不正なリクエストはバリデーションではじく",
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&simoompb.RefreshTokenRequest{}),
			},
			want:    nil,
			wantErr: apperr.ErrInvalidRequest(nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testHandler.RefreshToken(tt.args.ctx, tt.args.req)
			assert.Equal(t, tt.want, resp)

			var appErr apperr.Error
			if tt.wantErr.Error() != "" && assert.ErrorAs(t, err, &appErr) {
				assert.Equal(t, tt.wantErr.ConnectError().Code(), appErr.ConnectError().Code())
			}
		})
	}
}
