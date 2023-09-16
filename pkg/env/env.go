// Package env は環境変数に関するパッケージ
package env

import (
	"github.com/cockroachdb/errors"
	"github.com/kelseyhightower/envconfig"
)

// Load は環境変数を読み込み、取得する
func Load() (Env, error) {
	var appEnv Env
	if err := envconfig.Process("", &appEnv); err != nil {
		return Env{}, errors.WithStack(err)
	}
	return appEnv, nil
}
