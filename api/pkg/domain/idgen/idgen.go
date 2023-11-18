// Package idgen は ID 生成に関するパッケージ
package idgen

import "github.com/oklog/ulid/v2"

// Generate は ID を生成する
func Generate() string {
	return ulid.Make().String()
}
