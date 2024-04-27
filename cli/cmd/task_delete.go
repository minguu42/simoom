package cmd

import (
	"context"
	"fmt"
	"io"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/cli/api"
	"github.com/minguu42/simoom/cli/factory"
	"github.com/minguu42/simoom/lib/go/simoompb/v1"
	"github.com/spf13/cobra"
)

type TaskDeleteOpts struct {
	Client api.Client

	ID string
}

func NewCmdTaskDelete() *cobra.Command {
	var opts TaskDeleteOpts
	return &cobra.Command{
		Use:   "delete",
		Short: "Delete a task",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			f := factory.FromContext(cmd.Context())
			opts.Client = f.Client

			opts.ID = args[0]
			return TaskDeleteRun(cmd.Context(), f.Out, opts)
		},
	}
}

func TaskDeleteRun(ctx context.Context, out io.Writer, opts TaskDeleteOpts) error {
	if _, err := opts.Client.DeleteTask(ctx, connect.NewRequest(&simoompb.DeleteTaskRequest{
		Id: opts.ID,
	})); err != nil {
		return fmt.Errorf("failed to call DeleteTask method: %w", err)
	}

	fmt.Fprintf(out, "Task %s deleted\n", opts.ID)
	return nil
}
