// Package config はアプリケーションの設定値を扱うパッケージ
package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

// Load は環境変数から設定値を読み込んだ Config を返す
func Load() (Config, error) {
	var conf Config
	if err := envconfig.Process("", &conf); err != nil {
		return Config{}, fmt.Errorf("failed to populate the specified struct based on environment variables: %w", err)
	}
	return conf, nil
}

// Config はアプリケーションで使用する設定値
type Config struct {
	API  API
	Auth Auth
	DB   DB
}

// API はAPIに関する設定
type API struct {
	Host string `envconfig:"API_HOST" default:"0.0.0.0"`
	Port int    `envconfig:"API_PORT" default:"8080"`
}

// Auth は認証に関する設定
type Auth struct {
	AccessTokenExpiryHour  int    `envconfig:"ACCESS_TOKEN_EXPIRY_HOUR" default:"2"`
	RefreshTokenExpiryHour int    `envconfig:"REFRESH_TOKEN_EXPIRY_HOUR" default:"168"`
	AccessTokenSecret      string `envconfig:"ACCESS_TOKEN_SECRET" required:"true"`
	RefreshTokenSecret     string `envconfig:"REFRESH_TOKEN_SECRET" required:"true"`
}

// DB はデータベースに関する設定
type DB struct {
	Host               string `envconfig:"DB_HOST" required:"true"`
	Port               int    `envconfig:"DB_PORT" required:"true"`
	Database           string `envconfig:"DB_DATABASE" required:"true"`
	User               string `envconfig:"DB_USER" required:"true"`
	Password           string `envconfig:"DB_PASSWORD" required:"true"`
	ConnMaxLifetimeMin int    `envconfig:"DB_CONN_MAX_LIFETIME_MIN" default:"5"`
	MaxOpenConns       int    `envconfig:"DB_MAX_OPEN_CONNS" default:"25"`
	MaxIdleConns       int    `envconfig:"DB_MAX_IDLE_CONNS" default:"25"`
}
