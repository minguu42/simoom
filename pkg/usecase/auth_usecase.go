package usecase

import (
	"context"
	"time"

	"github.com/cockroachdb/errors"
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
		return SignUpOutput{}, errors.WithStack(err)
	}

	now := time.Now()
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
		return SignUpOutput{}, errors.WithStack(err)
	}
	refreshToken, err := uc.authenticator.CreateRefreshToken(user, uc.conf.RefreshTokenSecret, uc.conf.RefreshTokenExpiryHour)
	if err != nil {
		return SignUpOutput{}, errors.WithStack(err)
	}
	if err := uc.repo.CreateUser(ctx, user); err != nil {
		return SignUpOutput{}, errors.WithStack(err)
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
		return SignInOutput{}, errors.WithStack(err)
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(in.Password)) != nil {
		return SignInOutput{}, errors.New("password is not valid")
	}

	accessToken, err := uc.authenticator.CreateAccessToken(user, uc.conf.AccessTokenSecret, uc.conf.AccessTokenExpiryHour)
	if err != nil {
		return SignInOutput{}, errors.WithStack(err)
	}
	refreshToken, err := uc.authenticator.CreateRefreshToken(user, uc.conf.RefreshTokenSecret, uc.conf.RefreshTokenExpiryHour)
	if err != nil {
		return SignInOutput{}, errors.WithStack(err)
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

func (uc Auth) RefreshAccessToken(ctx context.Context, in RefreshAccessTokenInput) (RefreshAccessTokenOutput, error) {
	id, err := uc.authenticator.ExtractIDFromToken(in.RefreshToken, uc.conf.RefreshTokenSecret)
	if err != nil {
		return RefreshAccessTokenOutput{}, errors.WithStack(err)
	}
	user, err := uc.repo.GetUserByID(ctx, id)
	if err != nil {
		return RefreshAccessTokenOutput{}, errors.WithStack(err)
	}

	accessToken, err := uc.authenticator.CreateAccessToken(user, uc.conf.AccessTokenSecret, uc.conf.AccessTokenExpiryHour)
	if err != nil {
		return RefreshAccessTokenOutput{}, errors.WithStack(err)
	}
	refreshToken, err := uc.authenticator.CreateRefreshToken(user, uc.conf.RefreshTokenSecret, uc.conf.RefreshTokenExpiryHour)
	if err != nil {
		return RefreshAccessTokenOutput{}, errors.WithStack(err)
	}
	if err := uc.repo.CreateUser(ctx, user); err != nil {
		return RefreshAccessTokenOutput{}, errors.WithStack(err)
	}
	return RefreshAccessTokenOutput{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
