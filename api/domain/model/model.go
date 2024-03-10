// Package model はドメインモデルを定義する
package model

//go:generate moq -fmt goimports -out ./model_mock.go -rm . IDGenerator

// IDGenerator はドメインモデルの識別子を生成するジェネレータ
type IDGenerator interface {
	Generate() string
}
