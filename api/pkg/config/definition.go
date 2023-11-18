package config

// Config はアプリケーションで使用する設定値
type Config struct {
	API   API
	Auth  Auth
	MySQL MySQL
}

// API はAPIに関する設定値
type API struct {
	Host string `envconfig:"API_HOST" default:"0.0.0.0"`
	Port int    `envconfig:"API_PORT" default:"8080"`
}

// Auth は認証に関する設定値
type Auth struct {
	AccessTokenExpiryHour  int    `envconfig:"ACCESS_TOKEN_EXPIRY_HOUR" default:"2"`
	RefreshTokenExpiryHour int    `envconfig:"REFRESH_TOKEN_EXPIRY_HOUR" default:"168"`
	AccessTokenSecret      string `envconfig:"ACCESS_TOKEN_SECRET" required:"true"`
	RefreshTokenSecret     string `envconfig:"REFRESH_TOKEN_SECRET" required:"true"`
}

// MySQL はMySQLに関する設定値
type MySQL struct {
	Host               string `envconfig:"MYSQL_HOST" default:"db"`
	Port               int    `envconfig:"MYSQL_PORT" default:"3306"`
	Database           string `envconfig:"MYSQL_DATABASE" default:"simoomdb"`
	User               string `envconfig:"MYSQL_USER" default:"root"`
	Password           string `envconfig:"MYSQL_PASSWORD" default:""`
	ConnMaxLifetimeMin int    `envconfig:"MYSQL_CONN_MAX_LIFETIME_MIN" default:"5"`
	MaxOpenConns       int    `envconfig:"MYSQL_MAX_OPEN_CONNS" default:"25"`
	MaxIdleConns       int    `envconfig:"MYSQL_MAX_IDLE_CONNS" default:"25"`
}
