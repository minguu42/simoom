package cmd

import (
	"context"
	"errors"
	"fmt"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/cli/api"
	"github.com/minguu42/simoom/cli/cmdutil"
	"github.com/minguu42/simoom/lib/go/simoompb/v1"
	"github.com/spf13/cobra"
)

type tagViewOpts struct {
	client api.Client

	id     string
	limit  uint64
	offset uint64
}

func newCmdTagView(f cmdutil.Factory) *cobra.Command {
	opts := tagViewOpts{
		client: f.Client,
	}
	cmd := &cobra.Command{
		Use:   "view <id> [flags]",
		Short: "List the tasks with the tag",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.id = args[0]
			if len(opts.id) != 26 {
				return errors.New("id is a 26-character string")
			}
			return runTagView(cmd.Context(), opts)
		},
	}

	cmd.Flags().Uint64Var(&opts.limit, "limit", 10, "limit")
	cmd.Flags().Uint64Var(&opts.offset, "offset", 0, "offset")

	return cmd
}

func runTagView(ctx context.Context, opts tagViewOpts) error {
	resp, err := opts.client.ListTasksByTagID(ctx, connect.NewRequest(&simoompb.ListTasksByTagIDRequest{
		TagId:  opts.id,
		Limit:  opts.limit,
		Offset: opts.offset,
	}))
	if err != nil {
		return fmt.Errorf("failed to call ListTasksByTagID method: %w", err)
	}

	if err := cmdutil.PrintJSON(resp.Msg); err != nil {
		return fmt.Errorf("failed to print json output: %w", err)
	}
	return nil
}