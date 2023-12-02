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
		name string
		args args
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
		},
		{
			name: "nameに16文字以上の文字列は指定できない",
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&simoompb.SignUpRequest{
					Name:     "some-some-some-a",
					Email:    "dummy@example.com",
					Password: "password123456",
				}),
			},
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
		},
		{
			name: "emailに255文字以上の文字列は指定できない",
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&simoompb.SignUpRequest{
					Name:     "テストユーザ",
					Email:    "very-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-lon@example.com",
					Password: "password123456",
				}),
			},
		},
		{
			name: "passwordに11文字以下の文字列は指定できない",
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&simoompb.SignUpRequest{
					Name:     "テストユーザ",
					Email:    "dummy@example.com",
					Password: "short-pw123",
				}),
			},
		},
		{
			name: "passwordに21文字以上の文字列は指定できない",
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&simoompb.SignUpRequest{
					Name:     "テストユーザ",
					Email:    "dummy@example.com",
					Password: "long-long-password123",
				}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := th.SignUp(tt.args.ctx, tt.args.req)
			assert.Error(t, err)
		})
	}
}

func TestHandler_SignIn(t *testing.T) {
	type args struct {
		ctx context.Context
		req *connect.Request[simoompb.SignInRequest]
	}
	tests := []struct {
		name string
		args args
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
		},
		{
			name: "emailに255文字以上の文字列は指定できない",
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&simoompb.SignInRequest{
					Email:    "very-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-lon@example.com",
					Password: "password123456",
				}),
			},
		},
		{
			name: "passwordに11文字以下の文字列は指定できない",
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&simoompb.SignInRequest{
					Email:    "dummy@example.com",
					Password: "short-pw123",
				}),
			},
		},
		{
			name: "passwordに21文字以上の文字列は指定できない",
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&simoompb.SignInRequest{
					Email:    "dummy@example.com",
					Password: "long-long-password123",
				}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := th.SignIn(tt.args.ctx, tt.args.req)
			assert.Error(t, err)
		})
	}
}

func TestHandler_RefreshAccessToken(t *testing.T) {
	type args struct {
		ctx context.Context
		req *connect.Request[simoompb.RefreshTokenRequest]
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "refresh_tokenに空文字列は指定できない",
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&simoompb.RefreshTokenRequest{
					RefreshToken: "",
				}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := th.RefreshToken(tt.args.ctx, tt.args.req)
			assert.Error(t, err)
		})
	}
}
