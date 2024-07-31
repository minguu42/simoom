package factory

import (
	"fmt"

	"github.com/minguu42/simoom/api/adapter/jwtauth"
	"github.com/minguu42/simoom/api/adapter/mysql"
	"github.com/minguu42/simoom/api/adapter/ulidgen"
	"github.com/minguu42/simoom/api/config"
)

type Factory struct {
	Authn jwtauth.Authenticator
	IDGen ulidgen.Generator
	Repo  *mysql.Client
}

func New(conf config.Config) (*Factory, error) {
	c, err := mysql.NewClient(conf.DB)
	if err != nil {
		return nil, fmt.Errorf("failed to create mysql client: %w", err)
	}
	return &Factory{
		Authn: jwtauth.Authenticator{
			AccessTokenExpiryHour:  conf.Auth.AccessTokenExpiryHour,
			RefreshTokenExpiryHour: conf.Auth.RefreshTokenExpiryHour,
			AccessTokenSecret:      conf.Auth.AccessTokenSecret,
			RefreshTokenSecret:     conf.Auth.RefreshTokenSecret,
		},
		IDGen: ulidgen.Generator{},
		Repo:  c,
	}, nil
}

func (f *Factory) Close() {
	f.Repo.Close()
}
