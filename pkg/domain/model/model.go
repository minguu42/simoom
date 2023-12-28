// Package model はドメインモデルを定義する
package model

// IDGenerator はドメインモデルの識別子を生成するジェネレータ
type IDGenerator interface {
	Generate() string
}
