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

type TagDeleteOpts struct {
	Client api.Client

	ID string
}

func NewCmdTagDelete() *cobra.Command {
	var opts TagDeleteOpts
	return &cobra.Command{
		Use:   "delete",
		Short: "Delete a tag",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			f := factory.FromContext(cmd.Context())
			opts.Client = f.Client

			opts.ID = args[0]
			return TagDeleteRun(cmd.Context(), f.Out, opts)
		},
	}
}

func TagDeleteRun(ctx context.Context, out io.Writer, opts TagDeleteOpts) error {
	if _, err := opts.Client.DeleteTag(ctx, connect.NewRequest(&simoompb.DeleteTagRequest{
		Id: opts.ID,
	})); err != nil {
		return fmt.Errorf("failed to call DeleteTag method: %w", err)
	}

	fmt.Fprintf(out, "Tag %s deleted\n", opts.ID)
	return nil
}
