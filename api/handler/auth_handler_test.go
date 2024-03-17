package handler

import (
	"context"
	"testing"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/api/apperr"
	"github.com/minguu42/simoom/api/domain/auth"
	"github.com/minguu42/simoom/api/domain/model"
	"github.com/minguu42/simoom/api/domain/repository"
	"github.com/minguu42/simoom/api/usecase"
	"github.com/minguu42/simoom/lib/go/simoompb/v1"
	"github.com/stretchr/testify/assert"
)

func TestHandler_SignUp(t *testing.T) {
	th := handler{
		validator: testValidator,
		auth: usecase.NewAuth(
			&auth.AuthenticatorMock{
				CreateAccessTokenFunc: func(_ context.Context, _ model.User) (string, error) {
					return "some-access-token", nil
				},
				CreateRefreshTokenFunc: func(_ context.Context, _ model.User) (string, error) {
					return "some-refresh-token", nil
				},
			},
			&repository.RepositoryMock{
				CreateUserFunc: func(_ context.Context, _ model.User) error {
					return nil
				},
				GetUserByNameFunc: func(_ context.Context, _ string) (model.User, error) {
					return model.User{}, repository.ErrModelNotFound
				},
				GetUserByEmailFunc: func(_ context.Context, _ string) (model.User, error) {
					return model.User{}, repository.ErrModelNotFound
				},
			},
			&model.IDGeneratorMock{
				GenerateFunc: func() string {
					return "user_01"
				},
			},
		),
	}
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
			name: "ユーザを登録し、認証情報を返す",
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&simoompb.SignUpRequest{
					Name:     "ユーザ1",
					Email:    "test1@example.com",
					Password: "valid-password",
				}),
			},
			want: connect.NewResponse(&simoompb.SignUpResponse{
				AccessToken:  "some-access-token",
				RefreshToken: "some-refresh-token",
			}),
			wantErr: apperr.Error{},
		},
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
			resp, err := th.SignUp(tt.args.ctx, tt.args.req)
			assert.Equal(t, tt.want, resp)

			var appErr apperr.Error
			if tt.wantErr.Error() != "" && assert.ErrorAs(t, err, &appErr) {
				assert.Equal(t, tt.wantErr.ConnectError().Code(), appErr.ConnectError().Code())
			}
		})
	}
}
