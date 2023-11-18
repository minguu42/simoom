// Package config はアプリケーションの設定値を扱うパッケージ
package config

import (
	"github.com/cockroachdb/errors"
	"github.com/kelseyhightower/envconfig"
)

// Load は環境変数から設定値を読み込んだ Config を返す
func Load() (Config, error) {
	var conf Config
	if err := envconfig.Process("", &conf); err != nil {
		return Config{}, errors.WithStack(err)
	}
	return conf, nil
}
