package main

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/cli/cmdutil"
	"github.com/minguu42/simoom/lib/simoompb/v1"
	"github.com/spf13/cobra"
)

type projectDeleteOpts struct {
	id string
}

func newCmdProjectDelete(core cmdutil.Core) *cobra.Command {
	return &cobra.Command{
		Use:   "delete",
		Short: "Delete a project",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts := projectDeleteOpts{id: args[0]}
			return runProjectDelete(cmd.Context(), core, opts)
		},
	}
}

func runProjectDelete(ctx context.Context, core cmdutil.Core, opts projectDeleteOpts) error {
	req := connect.NewRequest(&simoompb.simoompb{Id: opts.id})
	req.Header().Set("Authorization", fmt.Sprintf("Bearer %s", core.Credentials.AccessToken))

	if _, err := core.Client.DeleteProject(ctx, req); err != nil {
		return fmt.Errorf("failed to call DeleteProject method: %w", err)
	}

	fmt.Println("project deleted")
	return nil
}
