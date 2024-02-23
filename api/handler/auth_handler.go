package handler

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/api/usecase"
	"github.com/minguu42/simoom/lib/go/simoompb/v1"
)

func (h handler) SignUp(ctx context.Context, req *connect.Request[simoompb.SignUpRequest]) (*connect.Response[simoompb.SignUpResponse], error) {
	if err := h.validator.Validate(req.Msg); err != nil {
		return nil, ErrInvalidRequest
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
	if err := h.validator.Validate(req.Msg); err != nil {
		return nil, ErrInvalidRequest
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
	if err := h.validator.Validate(req.Msg); err != nil {
		return nil, ErrInvalidRequest
	}

	out, err := h.auth.RefreshToken(ctx, usecase.RefreshTokenInput{RefreshToken: req.Msg.RefreshToken})
	if err != nil {
		return nil, fmt.Errorf("failed to execute RefreshToken usecase: %w", err)
	}
	return connect.NewResponse(&simoompb.RefreshTokenResponse{
		AccessToken:  out.AccessToken,
		RefreshToken: out.RefreshToken,
	}), nil
}
