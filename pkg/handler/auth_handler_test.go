package handler

import (
	"context"
	"testing"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/pkg/simoompb/v1"
	"github.com/stretchr/testify/assert"
)

func TestHandler_SignUp(t *testing.T) {
	type args struct {
		ctx context.Context
		req *connect.Request[simoompb.SignUpRequest]
	}
	tests := []struct {
		name     string
		args     args
		hasError bool
	}{
		{
			name: "nameに空文字列は指定できない",
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&simoompb.SignUpRequest{
					Name:     "",
					Email:    "dummy@example.com",
					Password: "password123456",
				}),
			},
			hasError: true,
		},
		{
			name: "emailに空文字列は指定できない",
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&simoompb.SignUpRequest{
					Name:     "テストユーザ",
					Email:    "",
					Password: "password123456",
				}),
			},
			hasError: true,
		},
		{
			name: "passwordは12文字以上、20文字以下である必要がある",
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&simoompb.SignUpRequest{
					Name:     "テストユーザ",
					Email:    "dummy@example.com",
					Password: "password",
				}),
			},
			hasError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := th.SignUp(tt.args.ctx, tt.args.req)
			if tt.hasError {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}

func TestHandler_SignIn(t *testing.T) {
	type args struct {
		ctx context.Context
		req *connect.Request[simoompb.SignInRequest]
	}
	tests := []struct {
		name     string
		args     args
		hasError bool
	}{
		{
			name: "emailに空文字列は指定できない",
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&simoompb.SignInRequest{
					Email:    "",
					Password: "password123456",
				}),
			},
			hasError: true,
		},
		{
			name: "passwordは12文字以上、20文字以下である必要がある",
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&simoompb.SignInRequest{
					Email:    "dummy@example.com",
					Password: "password",
				}),
			},
			hasError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := th.SignIn(tt.args.ctx, tt.args.req)
			if tt.hasError {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}

func TestHandler_RefreshAccessToken(t *testing.T) {
	type args struct {
		ctx context.Context
		req *connect.Request[simoompb.RefreshTokenRequest]
	}
	tests := []struct {
		name     string
		args     args
		hasError bool
	}{
		{
			name: "refresh_tokenに空文字列は指定できない",
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&simoompb.RefreshTokenRequest{
					RefreshToken: "",
				}),
			},
			hasError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := th.RefreshToken(tt.args.ctx, tt.args.req)
			if tt.hasError {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}
