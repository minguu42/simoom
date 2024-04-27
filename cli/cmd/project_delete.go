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

type ProjectDeleteOpts struct {
	client api.Client

	id string
}

func NewCmdProjectDelete() *cobra.Command {
	var opts ProjectDeleteOpts
	return &cobra.Command{
		Use:   "delete",
		Short: "Delete a project",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			f := factory.FromContext(cmd.Context())
			opts.client = f.Client

			opts.id = args[0]
			return ProjectDeleteRun(cmd.Context(), f.Out, opts)
		},
	}
}

func ProjectDeleteRun(ctx context.Context, out io.Writer, opts ProjectDeleteOpts) error {
	if _, err := opts.client.DeleteProject(ctx, connect.NewRequest(&simoompb.DeleteProjectRequest{
		Id: opts.id,
	})); err != nil {
		return fmt.Errorf("failed to call DeleteProject method: %w", err)
	}

	fmt.Fprintf(out, "Project %s deleted\n", opts.id)
	return nil
}
