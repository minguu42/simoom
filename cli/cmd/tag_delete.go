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

type tagDeleteOpts struct {
	client api.Client

	id string
}

func newCmdTagDelete() *cobra.Command {
	var opts tagDeleteOpts
	return &cobra.Command{
		Use:   "delete",
		Short: "Delete a tag",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			f := factory.Value(cmd.Context())
			opts.client = f.Client

			opts.id = args[0]
			return runTagDelete(cmd.Context(), opts)
		},
	}
}

func runTagDelete(ctx context.Context, opts tagDeleteOpts) error {
	if _, err := opts.client.DeleteTag(ctx, connect.NewRequest(&simoompb.DeleteTagRequest{
		Id: opts.id,
	})); err != nil {
		return fmt.Errorf("failed to call DeleteTag method: %w", err)
	}

	fmt.Println("tag deleted")
	return nil
}
