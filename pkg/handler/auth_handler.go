package handler

import (
	"context"

	"connectrpc.com/connect"
	"github.com/cockroachdb/errors"
	"github.com/minguu42/simoom/gen/simoompb/v1"
)

func (s simoom) SignIn(_ context.Context, _ *connect.Request[simoompb.SignInRequest]) (*connect.Response[simoompb.SignInResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("unimplemented"))
}

func (s simoom) SignUp(_ context.Context, _ *connect.Request[simoompb.SignUpRequest]) (*connect.Response[simoompb.SignUpResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("unimplemented"))
}

func (s simoom) RefreshAccessToken(_ context.Context, _ *connect.Request[simoompb.RefreshAccessTokenRequest]) (*connect.Response[simoompb.RefreshAccessTokenResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("unimplemented"))
}
