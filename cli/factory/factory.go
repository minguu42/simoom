// Package factory はファクトリを定義する
package factory

import (
	"context"
	"fmt"

	"github.com/minguu42/simoom/cli/api"
)

type Factory struct {
	Profile string
	Client  api.Client
}

// New は Factory を生成する
func New(profile string) (*Factory, error) {
	c, err := api.NewClient(profile)
	if err != nil {
		return nil, fmt.Errorf("failed to create api client: %w", err)
	}
	return &Factory{
		Profile: profile,
		Client:  c,
	}, nil
}

type factoryKey struct{}

// FromContext はコンテキストから Factory を取り出す
func FromContext(ctx context.Context) *Factory {
	return ctx.Value(factoryKey{}).(*Factory)
}

// ContextWithFactory は Factory を含めたコンテキストを返す
func ContextWithFactory(ctx context.Context, f *Factory) context.Context {
	return context.WithValue(ctx, factoryKey{}, f)
}
