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

type taskDeleteOpts struct {
	client api.Client

	id string
}

func newCmdTaskDelete() *cobra.Command {
	var opts taskDeleteOpts
	return &cobra.Command{
		Use:   "delete",
		Short: "Delete a task",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			f := factory.Value(cmd.Context())
			opts.client = f.Client

			opts.id = args[0]
			return runTaskDelete(cmd.Context(), opts)
		},
	}
}

func runTaskDelete(ctx context.Context, opts taskDeleteOpts) error {
	if _, err := opts.client.DeleteTask(ctx, connect.NewRequest(&simoompb.DeleteTaskRequest{
		Id: opts.id,
	})); err != nil {
		return fmt.Errorf("failed to call DeleteTask method: %w", err)
	}

	fmt.Println("task deleted")
	return nil
}
