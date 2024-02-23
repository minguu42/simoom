package usecase

import (
	"context"
	"errors"
	"fmt"
	"unicode/utf8"

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

func (in SignUpInput) Validate() error {
	if len(in.Name) < 1 || 15 < utf8.RuneCountInString(in.Name) {
		return newErrInvalidArgument("name must be at least 1 and no more than 15 characters")
	}
	if len(in.Email) < 1 || 254 < len(in.Email) {
		return newErrInvalidArgument("email must be at least 1 and no more than 254 characters")
	}
	if len(in.Password) < 12 || 20 < len(in.Password) {
		return newErrInvalidArgument("password must be at least 12 and no more than 20 characters")
	}
	return nil
}

type SignUpOutput struct {
	AccessToken  string
	RefreshToken string
}

func (uc Auth) SingUp(ctx context.Context, in SignUpInput) (SignUpOutput, error) {
	if err := in.Validate(); err != nil {
		return SignUpOutput{}, fmt.Errorf("failed to validate input: %w", err)
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return SignUpOutput{}, fmt.Errorf("failed to generate encypted password: %w", err)
	}

	user := model.User{
		ID:       uc.idgen.Generate(),
		Name:     in.Name,
		Email:    in.Email,
		Password: string(encryptedPassword),
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

func (in SignInInput) Validate() error {
	if len(in.Email) < 1 || 254 < len(in.Email) {
		return newErrInvalidArgument("email must be at least 1 and no more than 254 characters")
	}
	if len(in.Password) < 12 || 20 < len(in.Password) {
		return newErrInvalidArgument("password must be at least 12 and no more than 20 characters long")
	}
	return nil
}

type SignInOutput struct {
	AccessToken  string
	RefreshToken string
}

func (uc Auth) SignIn(ctx context.Context, in SignInInput) (SignInOutput, error) {
	if err := in.Validate(); err != nil {
		return SignInOutput{}, fmt.Errorf("failed to validate input: %w", err)
	}

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

func (in RefreshTokenInput) Validate() error {
	if in.RefreshToken == "" {
		return newErrInvalidArgument("refresh_token cannot be an empty string")
	}
	return nil
}

type RefreshTokenOutput struct {
	AccessToken  string
	RefreshToken string
}

func (uc Auth) RefreshToken(ctx context.Context, in RefreshTokenInput) (RefreshTokenOutput, error) {
	if err := in.Validate(); err != nil {
		return RefreshTokenOutput{}, fmt.Errorf("failed to validate input: %w", err)
	}

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
	if err := uc.repo.CreateUser(ctx, user); err != nil {
		return RefreshTokenOutput{}, fmt.Errorf("failed to create user: %w", err)
	}
	return RefreshTokenOutput{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
