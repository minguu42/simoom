package tag

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/cli/cmdutil"
	"github.com/minguu42/simoom/lib/go/simoompb/v1"
	"github.com/spf13/cobra"
)

type tagDeleteOpts struct {
	id string
}

func newCmdTagDelete(core cmdutil.Core) *cobra.Command {
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

func runTagDelete(ctx context.Context, core cmdutil.Core, opts tagDeleteOpts) error {
	req := connect.NewRequest(&simoompb.DeleteTagRequest{Id: opts.id})
	req.Header().Set("Authorization", fmt.Sprintf("Bearer %s", core.Credentials.AccessToken))

	if _, err := core.Client.DeleteTag(ctx, req); err != nil {
		return fmt.Errorf("failed to call DeleteTag method: %w", err)
	}

	fmt.Println("tag deleted")
	return nil
}
