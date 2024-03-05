package cmdutil

import (
	"fmt"

	"github.com/minguu42/simoom/cli/api"
)

type Factory struct {
	Profile string
	Client  api.Client
}

// NewFactory は Factory を生成し返す
func NewFactory(profile string) (*Factory, error) {
	c, err := api.NewClient(profile)
	if err != nil {
		return nil, fmt.Errorf("failed to create api client: %w", err)
	}
	return &Factory{
		Profile: profile,
		Client:  c,
	}, nil
}
