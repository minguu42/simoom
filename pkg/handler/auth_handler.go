package handler

import (
	"context"

	"connectrpc.com/connect"
	"github.com/cockroachdb/errors"
	"github.com/minguu42/simoom/gen/simoompb/v1"
	"github.com/minguu42/simoom/pkg/usecase"
)

func (s simoom) SignIn(ctx context.Context, req *connect.Request[simoompb.SignInRequest]) (*connect.Response[simoompb.SignInResponse], error) {
	out, err := s.auth.SignIn(ctx, usecase.SignInInput{
		Email:    req.Msg.Email,
		Password: req.Msg.Password,
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return connect.NewResponse(&simoompb.SignInResponse{
		AccessToken:  out.AccessToken,
		RefreshToken: out.RefreshToken,
	}), nil
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

func (s simoom) RefreshAccessToken(ctx context.Context, req *connect.Request[simoompb.RefreshAccessTokenRequest]) (*connect.Response[simoompb.RefreshAccessTokenResponse], error) {
	out, err := s.auth.RefreshAccessToken(ctx, usecase.RefreshAccessTokenInput{RefreshToken: req.Msg.RefreshToken})
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return connect.NewResponse(&simoompb.RefreshAccessTokenResponse{
		AccessToken:  out.AccessToken,
		RefreshToken: out.RefreshToken,
	}), nil
}
