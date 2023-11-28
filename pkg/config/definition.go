package config

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
	Host               string `envconfig:"DB_HOST" default:"db"`
	Port               int    `envconfig:"DB_PORT" default:"3306"`
	Database           string `envconfig:"DB_DATABASE" default:"simoomdb"`
	User               string `envconfig:"DB_USER" default:"root"`
	Password           string `envconfig:"DB_PASSWORD" default:""`
	ConnMaxLifetimeMin int    `envconfig:"DB_CONN_MAX_LIFETIME_MIN" default:"5"`
	MaxOpenConns       int    `envconfig:"DB_MAX_OPEN_CONNS" default:"25"`
	MaxIdleConns       int    `envconfig:"DB_MAX_IDLE_CONNS" default:"25"`
}
