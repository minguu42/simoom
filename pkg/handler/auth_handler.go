package handler

import (
	"context"

	"connectrpc.com/connect"
	"github.com/cockroachdb/errors"
	"github.com/minguu42/simoom/gen/simoompb/v1"
	"github.com/minguu42/simoom/pkg/usecase"
)

func (s simoom) SignIn(_ context.Context, _ *connect.Request[simoompb.SignInRequest]) (*connect.Response[simoompb.SignInResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("unimplemented"))
}

func (s simoom) SignUp(ctx context.Context, req *connect.Request[simoompb.SignUpRequest]) (*connect.Response[simoompb.SignUpResponse], error) {
	out, err := s.auth.SingUp(ctx, usecase.SignUpInput{
		Name:     req.Msg.Name,
		Email:    req.Msg.Email,
		Password: req.Msg.Password,
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return connect.NewResponse(&simoompb.SignUpResponse{
		AccessToken:  out.AccessToken,
		RefreshToken: out.RefreshToken,
	}), nil
}

func (s simoom) RefreshAccessToken(_ context.Context, _ *connect.Request[simoompb.RefreshAccessTokenRequest]) (*connect.Response[simoompb.RefreshAccessTokenResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("unimplemented"))
}
