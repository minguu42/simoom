package handler

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/pkg/simoompb/v1"
	"github.com/minguu42/simoom/pkg/usecase"
)

func (h handler) SignUp(ctx context.Context, req *connect.Request[simoompb.SignUpRequest]) (*connect.Response[simoompb.SignUpResponse], error) {
	if len(req.Msg.Name) < 1 || 15 < len(req.Msg.Name) {
		return nil, newErrInvalidArgument("name must be at least 1 and no more than 15 characters")
	}
	if len(req.Msg.Email) < 1 || 254 < len(req.Msg.Email) {
		return nil, newErrInvalidArgument("email must be at least 1 and no more than 254 characters")
	}
	if len(req.Msg.Password) < 12 || 20 < len(req.Msg.Password) {
		return nil, newErrInvalidArgument("password must be at least 12 and no more than 20 characters")
	}

	out, err := h.auth.SingUp(ctx, usecase.SignUpInput{
		Name:     req.Msg.Name,
		Email:    req.Msg.Email,
		Password: req.Msg.Password,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to execute SignUp usecase: %w", err)
	}
	return connect.NewResponse(&simoompb.SignUpResponse{
		AccessToken:  out.AccessToken,
		RefreshToken: out.RefreshToken,
	}), nil
}

func (h handler) SignIn(ctx context.Context, req *connect.Request[simoompb.SignInRequest]) (*connect.Response[simoompb.SignInResponse], error) {
	if len(req.Msg.Email) < 1 || 254 < len(req.Msg.Email) {
		return nil, newErrInvalidArgument("email must be at least 1 and no more than 254 characters")
	}
	if len(req.Msg.Password) < 12 || 20 < len(req.Msg.Password) {
		return nil, newErrInvalidArgument("password must be at least 12 and no more than 20 characters long")
	}

	out, err := h.auth.SignIn(ctx, usecase.SignInInput{
		Email:    req.Msg.Email,
		Password: req.Msg.Password,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to execute SignIn usecase: %w", err)
	}
	return connect.NewResponse(&simoompb.SignInResponse{
		AccessToken:  out.AccessToken,
		RefreshToken: out.RefreshToken,
	}), nil
}

func (h handler) RefreshToken(ctx context.Context, req *connect.Request[simoompb.RefreshTokenRequest]) (*connect.Response[simoompb.RefreshTokenResponse], error) {
	if req.Msg.RefreshToken == "" {
		return nil, newErrInvalidArgument("refresh_token cannot be an empty string")
	}

	out, err := h.auth.RefreshToken(ctx, usecase.RefreshAccessTokenInput{RefreshToken: req.Msg.RefreshToken})
	if err != nil {
		return nil, fmt.Errorf("failed to execute RefreshToken usecase: %w", err)
	}
	return connect.NewResponse(&simoompb.RefreshTokenResponse{
		AccessToken:  out.AccessToken,
		RefreshToken: out.RefreshToken,
	}), nil
}
