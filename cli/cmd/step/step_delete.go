package step

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/cli/api"
	"github.com/minguu42/simoom/cli/cmdutil"
	"github.com/minguu42/simoom/lib/go/simoompb/v1"
	"github.com/spf13/cobra"
)

type stepDeleteOpts struct {
	client *api.Client

	id string
}

func newCmdStepDelete(f cmdutil.Factory) *cobra.Command {
	opts := stepDeleteOpts{
		client: f.Client,
	}
	return &cobra.Command{
		Use:   "delete",
		Short: "Delete a step",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.id = args[0]
			return runStepDelete(cmd.Context(), opts)
		},
	}
}

func runStepDelete(ctx context.Context, opts stepDeleteOpts) error {
	if _, err := opts.client.DeleteStep(ctx, connect.NewRequest(&simoompb.DeleteStepRequest{
		Id: opts.id,
	})); err != nil {
		return fmt.Errorf("failed to call DeleteStep method: %w", err)
	}

	fmt.Println("step deleted")
	return nil
}
