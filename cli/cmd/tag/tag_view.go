package tag

import (
	"context"
	"errors"
	"fmt"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/cli/cmdutil"
	"github.com/minguu42/simoom/pkg/simoompb/v1"
	"github.com/spf13/cobra"
)

type tagViewOpts struct {
	id     string
	limit  uint64
	offset uint64
}

func newCmdTagView(core cmdutil.Core) *cobra.Command {
	var opts tagViewOpts
	cmd := &cobra.Command{
		Use:   "view",
		Short: "List the tasks with the tag",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, _ []string) error {
			if opts.id == "" {
				return errors.New("id is required")
			}
			return runTagView(cmd.Context(), core, opts)
		},
	}

	cmd.Flags().StringVar(&opts.id, "id", "", "tag id")
	cmd.Flags().Uint64Var(&opts.limit, "limit", 10, "limit")
	cmd.Flags().Uint64Var(&opts.offset, "offset", 0, "offset")

	return cmd
}

func runTagView(ctx context.Context, core cmdutil.Core, opts tagViewOpts) error {
	req := connect.NewRequest(&simoompb.ListTasksByTagIDRequest{
		TagId:  opts.id,
		Limit:  opts.limit,
		Offset: opts.offset,
	})
	req.Header().Set("Authorization", fmt.Sprintf("Bearer %s", core.Credentials.AccessToken))

	resp, err := core.Client.ListTasksByTagID(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to call ListTasksByTagID method: %w", err)
	}

	if err := cmdutil.PrintJSON(resp.Msg); err != nil {
		return fmt.Errorf("failed to print json output: %w", err)
	}
	return nil
}
