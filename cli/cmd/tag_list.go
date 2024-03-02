package cmd

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/cli/api"
	"github.com/minguu42/simoom/cli/cmdutil"
	"github.com/minguu42/simoom/lib/go/simoompb/v1"
	"github.com/spf13/cobra"
)

type tagListOpts struct {
	client *api.Client

	limit  uint64
	offset uint64
}

func newCmdTagList(f cmdutil.Factory) *cobra.Command {
	opts := tagListOpts{
		client: f.Client,
	}
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List the tags",
		Args:    cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runTagList(cmd.Context(), opts)
		},
	}

	cmd.Flags().Uint64Var(&opts.limit, "limit", 10, "limit")
	cmd.Flags().Uint64Var(&opts.offset, "offset", 0, "offset")

	return cmd
}

func runTagList(ctx context.Context, opts tagListOpts) error {
	resp, err := opts.client.ListTags(ctx, connect.NewRequest(&simoompb.ListTagsRequest{
		Limit:  opts.limit,
		Offset: opts.offset,
	}))
	if err != nil {
		return fmt.Errorf("failed to call ListTags method: %w", err)
	}

	if err := cmdutil.PrintJSON(resp.Msg); err != nil {
		return fmt.Errorf("failed to print json output: %w", err)
	}
	return nil
}
