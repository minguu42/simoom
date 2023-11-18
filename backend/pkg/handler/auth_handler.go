package handler

import (
	"context"

	"connectrpc.com/connect"
	"github.com/cockroachdb/errors"
	"github.com/minguu42/simoom/backend/pkg/usecase"
	"github.com/minguu42/simoom/lib/simoompb/v1"
)

func (h handler) SignUp(ctx context.Context, req *connect.Request[simoompb.SignUpRequest]) (*connect.Response[simoompb.SignUpResponse], error) {
	if req.Msg.Name == "" {
		return nil, newErrInvalidArgument("name cannot be an empty string")
	}
	if req.Msg.Email == "" {
		return nil, newErrInvalidArgument("email cannot be an empty string")
	}
	if len(req.Msg.Password) < 12 || 20 < len(req.Msg.Password) {
		return nil, newErrInvalidArgument("password must be at least 12 and no more than 20 characters long")
	}

	out, err := h.auth.SingUp(ctx, usecase.SignUpInput{
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

func (h handler) SignIn(ctx context.Context, req *connect.Request[simoompb.SignInRequest]) (*connect.Response[simoompb.SignInResponse], error) {
	if req.Msg.Email == "" {
		return nil, newErrInvalidArgument("email cannot be an empty string")
	}
	if len(req.Msg.Password) < 12 || 20 < len(req.Msg.Password) {
		return nil, newErrInvalidArgument("password must be at least 12 and no more than 20 characters long")
	}

	out, err := h.auth.SignIn(ctx, usecase.SignInInput{
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

func (h handler) RefreshAccessToken(ctx context.Context, req *connect.Request[simoompb.RefreshAccessTokenRequest]) (*connect.Response[simoompb.RefreshAccessTokenResponse], error) {
	if req.Msg.RefreshToken == "" {
		return nil, newErrInvalidArgument("refresh_token cannot be an empty string")
	}

	out, err := h.auth.RefreshAccessToken(ctx, usecase.RefreshAccessTokenInput{RefreshToken: req.Msg.RefreshToken})
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return connect.NewResponse(&simoompb.RefreshAccessTokenResponse{
		AccessToken:  out.AccessToken,
		RefreshToken: out.RefreshToken,
	}), nil
}
