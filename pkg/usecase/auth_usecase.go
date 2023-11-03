package usecase

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/minguu42/simoom/pkg/domain/auth"
	"github.com/minguu42/simoom/pkg/domain/idgen"
	"github.com/minguu42/simoom/pkg/domain/model"
	"github.com/minguu42/simoom/pkg/domain/repository"
	"github.com/minguu42/simoom/pkg/env"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase struct {
	Repo repository.Repository
	Env  env.Env
}

type SignInInput struct {
	Email    string
	Password string
}

type SignInOutput struct {
	AccessToken  string
	RefreshToken string
}

func (u AuthUsecase) SignIn(ctx context.Context, in SignInInput) (SignInOutput, error) {
	user, err := u.Repo.GetUserByEmail(ctx, in.Email)
	if err != nil {
		return SignInOutput{}, errors.WithStack(err)
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(in.Password)) != nil {
		return SignInOutput{}, errors.New("password is not valid")
	}

	accessToken, err := auth.CreateAccessToken(user, u.Env.API.AccessTokenSecret, u.Env.API.AccessTokenExpiryHour)
	if err != nil {
		return SignInOutput{}, errors.WithStack(err)
	}
	refreshToken, err := auth.CreateRefreshToken(user, u.Env.API.RefreshTokenSecret, u.Env.API.RefreshTokenExpiryHour)
	if err != nil {
		return SignInOutput{}, errors.WithStack(err)
	}
	return SignInOutput{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
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

func (u AuthUsecase) SingUp(ctx context.Context, in SignUpInput) (SignUpOutput, error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return SignUpOutput{}, errors.WithStack(err)
	}

	user := model.User{
		ID:       idgen.Generate(),
		Name:     in.Name,
		Email:    in.Email,
		Password: string(encryptedPassword),
	}
	accessToken, err := auth.CreateAccessToken(user, u.Env.API.AccessTokenSecret, u.Env.API.AccessTokenExpiryHour)
	if err != nil {
		return SignUpOutput{}, errors.WithStack(err)
	}
	refreshToken, err := auth.CreateRefreshToken(user, u.Env.API.RefreshTokenSecret, u.Env.API.RefreshTokenExpiryHour)
	if err != nil {
		return SignUpOutput{}, errors.WithStack(err)
	}
	if err := u.Repo.CreateUser(ctx, user); err != nil {
		return SignUpOutput{}, errors.WithStack(err)
	}
	return SignUpOutput{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
