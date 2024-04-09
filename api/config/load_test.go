package config_test

import (
	"testing"

	"github.com/minguu42/simoom/api/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoad(t *testing.T) {
	envs := map[string]string{
		"ACCESS_TOKEN_SECRET":  "test-access-token-secret",
		"REFRESH_TOKEN_SECRET": "test-refresh-token-secret",
		"DB_HOST":              "localhost",
		"DB_PORT":              "3000",
		"DB_DATABASE":          "testdb",
		"DB_USER":              "test-user",
		"DB_PASSWORD":          "test-password",
	}
	for k, v := range envs {
		t.Setenv(k, v)
	}

	want := config.Config{
		API: config.API{
			Host: "0.0.0.0",
			Port: 8080,
		},
		Auth: config.Auth{
			AccessTokenExpiryHour:  2,
			RefreshTokenExpiryHour: 168,
			AccessTokenSecret:      "test-access-token-secret",
			RefreshTokenSecret:     "test-refresh-token-secret",
		},
		DB: config.DB{
			Host:               "localhost",
			Port:               3000,
			Database:           "testdb",
			User:               "test-user",
			Password:           "test-password",
			ConnMaxLifetimeMin: 5,
			MaxOpenConns:       25,
			MaxIdleConns:       25,
		},
	}
	got, err := config.Load()
	require.NoError(t, err)
	assert.Equal(t, want, got)
}
