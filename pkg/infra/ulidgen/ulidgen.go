// Package ulidgen はULIDを生成する機能を提供するパッケージ
package ulidgen

import "github.com/oklog/ulid/v2"

type Generator struct{}

// Generate はULIDを生成する
func (g Generator) Generate() string {
	return ulid.Make().String()
}
