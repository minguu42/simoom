package cmd

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/cli/api"
	"github.com/minguu42/simoom/cli/factory"
	"github.com/minguu42/simoom/lib/go/simoompb/v1"
	"github.com/spf13/cobra"
)

type StepDeleteOpts struct {
	client api.Client

	id string
}

func NewCmdStepDelete() *cobra.Command {
	var opts StepDeleteOpts
	return &cobra.Command{
		Use:   "delete",
		Short: "Delete a step",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			f := factory.FromContext(cmd.Context())
			opts.client = f.Client

			opts.id = args[0]
			return StepDeleteRun(cmd.Context(), opts)
		},
	}
}

func StepDeleteRun(ctx context.Context, opts StepDeleteOpts) error {
	if _, err := opts.client.DeleteStep(ctx, connect.NewRequest(&simoompb.DeleteStepRequest{
		Id: opts.id,
	})); err != nil {
		return fmt.Errorf("failed to call DeleteStep method: %w", err)
	}

	fmt.Println("step deleted")
	return nil
}
