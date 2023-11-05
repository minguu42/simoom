package config

// Env はアプリケーションで使用する環境変数
type Env struct {
	API   API
	MySQL MySQL
}

// API は API に関する環境変数
type API struct {
	Host                   string `envconfig:"API_HOST" default:"0.0.0.0"`
	Port                   int    `envconfig:"API_PORT" default:"8080"`
	AccessTokenExpiryHour  int    `envconfig:"ACCESS_TOKEN_EXPIRY_HOUR" default:"2"`
	RefreshTokenExpiryHour int    `envconfig:"REFRESH_TOKEN_EXPIRY_HOUR" default:"168"`
	AccessTokenSecret      string `envconfig:"ACCESS_TOKEN_SECRET" required:"true"`
	RefreshTokenSecret     string `envconfig:"REFRESH_TOKEN_SECRET" required:"true"`
}

// MySQL は MySQL に関する環境変数
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
