// Package config はアプリケーションの設定値を扱うパッケージ
package config

import (
	"github.com/cockroachdb/errors"
	"github.com/kelseyhightower/envconfig"
)

// Load は環境変数を読み込み、取得する
func Load() (Env, error) {
	var conf Env
	if err := envconfig.Process("", &conf); err != nil {
		return Env{}, errors.WithStack(err)
	}
	return conf, nil
}
