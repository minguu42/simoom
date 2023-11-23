package task

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/cli/cmdutil"
	"github.com/minguu42/simoom/pkg/simoompb/v1"
	"github.com/spf13/cobra"
)

type taskDeleteOpts struct {
	id string
}

func newCmdTaskDelete(core cmdutil.Core) *cobra.Command {
	return &cobra.Command{
		Use:   "delete",
		Short: "Delete a task",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts := taskDeleteOpts{id: args[0]}
			return runTaskDelete(cmd.Context(), core, opts)
		},
	}
}

func runTaskDelete(ctx context.Context, core cmdutil.Core, opts taskDeleteOpts) error {
	req := connect.NewRequest(&simoompb.DeleteTaskRequest{Id: opts.id})
	req.Header().Set("Authorization", fmt.Sprintf("Bearer %s", core.Credentials.AccessToken))

	if _, err := core.Client.DeleteTask(ctx, req); err != nil {
		return fmt.Errorf("failed to call DeleteTask method: %w", err)
	}

	fmt.Println("task deleted")
	return nil
}
