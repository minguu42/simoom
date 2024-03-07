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

type projectDeleteOpts struct {
	client api.Client

	id string
}

func newCmdProjectDelete() *cobra.Command {
	var opts projectDeleteOpts
	return &cobra.Command{
		Use:   "delete",
		Short: "Delete a project",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			f := factory.Value(cmd.Context())
			opts.client = f.Client

			opts.id = args[0]
			return runProjectDelete(cmd.Context(), opts)
		},
	}
}

func runProjectDelete(ctx context.Context, opts projectDeleteOpts) error {
	if _, err := opts.client.DeleteProject(ctx, connect.NewRequest(&simoompb.DeleteProjectRequest{
		Id: opts.id,
	})); err != nil {
		return fmt.Errorf("failed to call DeleteProject method: %w", err)
	}

	fmt.Println("project deleted")
	return nil
}
