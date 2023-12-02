package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/minguu42/simoom/pkg/clock"
	"github.com/minguu42/simoom/pkg/config"
	"github.com/minguu42/simoom/pkg/domain/auth"
	"github.com/minguu42/simoom/pkg/domain/idgen"
	"github.com/minguu42/simoom/pkg/domain/model"
	"github.com/minguu42/simoom/pkg/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	authenticator auth.Authenticator
	repo          repository.Repository
	conf          config.Auth
}

func NewAuth(authenticator auth.Authenticator, repo repository.Repository, conf config.Auth) Auth {
	return Auth{
		authenticator: authenticator,
		repo:          repo,
		conf:          conf,
	}
}

type SignUpInput struct {
	Name     string
	Email    string
	Password string
}

type SignUpOutput struct {
	AccessToken  string
	RefreshToken string
}

func (uc Auth) SingUp(ctx context.Context, in SignUpInput) (SignUpOutput, error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return SignUpOutput{}, fmt.Errorf("failed to generate encypted password: %w", err)
	}

	now := clock.Now(ctx)
	user := model.User{
		ID:        idgen.Generate(),
		Name:      in.Name,
		Email:     in.Email,
		Password:  string(encryptedPassword),
		CreatedAt: now,
		UpdatedAt: now,
	}
	accessToken, err := uc.authenticator.CreateAccessToken(user, uc.conf.AccessTokenSecret, uc.conf.AccessTokenExpiryHour)
	if err != nil {
		return SignUpOutput{}, fmt.Errorf("failed to create access token: %w", err)
	}
	refreshToken, err := uc.authenticator.CreateRefreshToken(user, uc.conf.RefreshTokenSecret, uc.conf.RefreshTokenExpiryHour)
	if err != nil {
		return SignUpOutput{}, fmt.Errorf("failed to create refresh token: %w", err)
	}
	if err := uc.repo.CreateUser(ctx, user); err != nil {
		return SignUpOutput{}, fmt.Errorf("failed to create user: %w", err)
	}
	return SignUpOutput{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

type SignInInput struct {
	Email    string
	Password string
}

type SignInOutput struct {
	AccessToken  string
	RefreshToken string
}

func (uc Auth) SignIn(ctx context.Context, in SignInInput) (SignInOutput, error) {
	user, err := uc.repo.GetUserByEmail(ctx, in.Email)
	if err != nil {
		return SignInOutput{}, fmt.Errorf("failed to get user: %w", err)
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(in.Password)) != nil {
		return SignInOutput{}, errors.New("password is not valid")
	}

	accessToken, err := uc.authenticator.CreateAccessToken(user, uc.conf.AccessTokenSecret, uc.conf.AccessTokenExpiryHour)
	if err != nil {
		return SignInOutput{}, fmt.Errorf("failed to create access token: %w", err)
	}
	refreshToken, err := uc.authenticator.CreateRefreshToken(user, uc.conf.RefreshTokenSecret, uc.conf.RefreshTokenExpiryHour)
	if err != nil {
		return SignInOutput{}, fmt.Errorf("failed to create refresh token: %w", err)
	}
	return SignInOutput{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

type RefreshAccessTokenInput struct {
	RefreshToken string
}

type RefreshAccessTokenOutput struct {
	AccessToken  string
	RefreshToken string
}

func (uc Auth) RefreshToken(ctx context.Context, in RefreshAccessTokenInput) (RefreshAccessTokenOutput, error) {
	id, err := uc.authenticator.ExtractIDFromToken(in.RefreshToken, uc.conf.RefreshTokenSecret)
	if err != nil {
		return RefreshAccessTokenOutput{}, fmt.Errorf("failed to extract id from token: %w", err)
	}
	user, err := uc.repo.GetUserByID(ctx, id)
	if err != nil {
		return RefreshAccessTokenOutput{}, fmt.Errorf("failed to get user: %w", err)
	}

	accessToken, err := uc.authenticator.CreateAccessToken(user, uc.conf.AccessTokenSecret, uc.conf.AccessTokenExpiryHour)
	if err != nil {
		return RefreshAccessTokenOutput{}, fmt.Errorf("failed to create access token: %w", err)
	}
	refreshToken, err := uc.authenticator.CreateRefreshToken(user, uc.conf.RefreshTokenSecret, uc.conf.RefreshTokenExpiryHour)
	if err != nil {
		return RefreshAccessTokenOutput{}, fmt.Errorf("failed to create refresh token: %w", err)
	}
	if err := uc.repo.CreateUser(ctx, user); err != nil {
		return RefreshAccessTokenOutput{}, fmt.Errorf("failed to create user: %w", err)
	}
	return RefreshAccessTokenOutput{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
