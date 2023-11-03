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
