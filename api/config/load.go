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
