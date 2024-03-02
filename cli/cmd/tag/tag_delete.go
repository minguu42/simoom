package tag

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/cli/api"
	"github.com/minguu42/simoom/cli/cmdutil"
	"github.com/minguu42/simoom/lib/go/simoompb/v1"
	"github.com/spf13/cobra"
)

type tagDeleteOpts struct {
	client *api.Client

	id string
}

func newCmdTagDelete(core cmdutil.Factory) *cobra.Command {
	return &cobra.Command{
		Use:   "delete",
		Short: "Delete a tag",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts := tagDeleteOpts{id: args[0]}
			return runTagDelete(cmd.Context(), core, opts)
		},
	}
}

func runTagDelete(ctx context.Context, core cmdutil.Factory, opts tagDeleteOpts) error {
	if _, err := core.Client.DeleteTag(ctx, connect.NewRequest(&simoompb.DeleteTagRequest{
		Id: opts.id,
	})); err != nil {
		return fmt.Errorf("failed to call DeleteTag method: %w", err)
	}

	fmt.Println("tag deleted")
	return nil
}
