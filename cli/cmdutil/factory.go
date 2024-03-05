package cmdutil

import (
	"fmt"

	"github.com/minguu42/simoom/cli/api"
)

type Factory struct {
	Client api.Client
}

// NewFactory はFactory を生成し返す
func NewFactory() (*Factory, error) {
	c, err := api.NewClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create api client: %w", err)
	}
	return &Factory{
		Client: c,
	}, nil
}
