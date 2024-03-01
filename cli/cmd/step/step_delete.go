package step

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/cli/cmdutil"
	"github.com/minguu42/simoom/lib/go/simoompb/v1"
	"github.com/spf13/cobra"
)

type stepDeleteOpts struct {
	id string
}

func newCmdStepDelete(core cmdutil.Factory) *cobra.Command {
	return &cobra.Command{
		Use:   "delete",
		Short: "Delete a step",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts := stepDeleteOpts{id: args[0]}
			return runStepDelete(cmd.Context(), core, opts)
		},
	}
}

func runStepDelete(ctx context.Context, core cmdutil.Factory, opts stepDeleteOpts) error {
	req := connect.NewRequest(&simoompb.DeleteStepRequest{Id: opts.id})
	req.Header().Set("Authorization", fmt.Sprintf("Bearer %s", core.Credentials.AccessToken))

	if _, err := core.Client.DeleteStep(ctx, req); err != nil {
		return fmt.Errorf("failed to call DeleteStep method: %w", err)
	}

	fmt.Println("step deleted")
	return nil
}
