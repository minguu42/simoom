package config

import "time"

type Config struct {
	API  API
	Auth Auth
	DB   DB
}

type API struct {
	Host              string        `envconfig:"API_HOST" default:"0.0.0.0"`
	Port              int           `envconfig:"API_PORT" default:"8080"`
	ReadTimeout       time.Duration `envconfig:"API_READ_TIMEOUT" default:"2s"`
	ReadHeaderTimeout time.Duration `envconfig:"API_READ_HEADER_TIMEOUT" default:"2s"`
	Timeout           time.Duration `envconfig:"API_TIMEOUT" default:"5s"`
}

type Auth struct {
	AccessTokenExpiryHour  int    `envconfig:"ACCESS_TOKEN_EXPIRY_HOUR" default:"2"`
	RefreshTokenExpiryHour int    `envconfig:"REFRESH_TOKEN_EXPIRY_HOUR" default:"168"`
	AccessTokenSecret      string `envconfig:"ACCESS_TOKEN_SECRET" required:"true"`
	RefreshTokenSecret     string `envconfig:"REFRESH_TOKEN_SECRET" required:"true"`
}

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
