package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/minguu42/simoom/api/config"
	"github.com/minguu42/simoom/api/domain/auth"
	"github.com/minguu42/simoom/api/domain/model"
	"github.com/minguu42/simoom/api/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	authenticator auth.Authenticator
	repo          repository.Repository
	conf          config.Auth
	idgen         model.IDGenerator
}

func NewAuth(authenticator auth.Authenticator, repo repository.Repository, conf config.Auth, idgen model.IDGenerator) Auth {
	return Auth{
		authenticator: authenticator,
		repo:          repo,
		conf:          conf,
		idgen:         idgen,
	}
}

type SignUpInput struct {
	Name     string
	Email    string
	Password string
}

func (in SignUpInput) Create(g model.IDGenerator) (model.User, error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return model.User{}, fmt.Errorf("failed to generate encypted password: %w", err)
	}
	return model.User{
		ID:       g.Generate(),
		Name:     in.Name,
		Email:    in.Email,
		Password: string(encryptedPassword),
	}, nil
}

type SignUpOutput struct {
	AccessToken  string
	RefreshToken string
}

func (uc Auth) SingUp(ctx context.Context, in SignUpInput) (SignUpOutput, error) {
	user, err := in.Create(uc.idgen)
	if err != nil {
		return SignUpOutput{}, fmt.Errorf("failed to create user: %w", err)
	}
	accessToken, err := uc.authenticator.CreateAccessToken(ctx, user, uc.conf.AccessTokenSecret, uc.conf.AccessTokenExpiryHour)
	if err != nil {
		return SignUpOutput{}, fmt.Errorf("failed to create access token: %w", err)
	}
	refreshToken, err := uc.authenticator.CreateRefreshToken(ctx, user, uc.conf.RefreshTokenSecret, uc.conf.RefreshTokenExpiryHour)
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

	accessToken, err := uc.authenticator.CreateAccessToken(ctx, user, uc.conf.AccessTokenSecret, uc.conf.AccessTokenExpiryHour)
	if err != nil {
		return SignInOutput{}, fmt.Errorf("failed to create access token: %w", err)
	}
	refreshToken, err := uc.authenticator.CreateRefreshToken(ctx, user, uc.conf.RefreshTokenSecret, uc.conf.RefreshTokenExpiryHour)
	if err != nil {
		return SignInOutput{}, fmt.Errorf("failed to create refresh token: %w", err)
	}
	return SignInOutput{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

type RefreshTokenInput struct {
	RefreshToken string
}

type RefreshTokenOutput struct {
	AccessToken  string
	RefreshToken string
}

func (uc Auth) RefreshToken(ctx context.Context, in RefreshTokenInput) (RefreshTokenOutput, error) {
	id, err := uc.authenticator.ExtractIDFromToken(in.RefreshToken, uc.conf.RefreshTokenSecret)
	if err != nil {
		return RefreshTokenOutput{}, fmt.Errorf("failed to extract id from token: %w", err)
	}
	user, err := uc.repo.GetUserByID(ctx, id)
	if err != nil {
		return RefreshTokenOutput{}, fmt.Errorf("failed to get user: %w", err)
	}

	accessToken, err := uc.authenticator.CreateAccessToken(ctx, user, uc.conf.AccessTokenSecret, uc.conf.AccessTokenExpiryHour)
	if err != nil {
		return RefreshTokenOutput{}, fmt.Errorf("failed to create access token: %w", err)
	}
	refreshToken, err := uc.authenticator.CreateRefreshToken(ctx, user, uc.conf.RefreshTokenSecret, uc.conf.RefreshTokenExpiryHour)
	if err != nil {
		return RefreshTokenOutput{}, fmt.Errorf("failed to create refresh token: %w", err)
	}
	return RefreshTokenOutput{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
