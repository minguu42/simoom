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

type TagListOpts struct {
	client api.Client

	limit  uint64
	offset uint64
}

func NewCmdTagList() *cobra.Command {
	var opts TagListOpts
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List the tags",
		Args:    cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, _ []string) error {
			f := factory.FromContext(cmd.Context())
			opts.client = f.Client

			return TagListRun(cmd.Context(), f.Out, opts)
		},
	}
	cmd.Flags().Uint64Var(&opts.limit, "limit", 10, "limit")
	cmd.Flags().Uint64Var(&opts.offset, "offset", 0, "offset")
	return cmd
}

func TagListRun(ctx context.Context, out io.Writer, opts TagListOpts) error {
	resp, err := opts.client.ListTags(ctx, connect.NewRequest(&simoompb.ListTagsRequest{
		Limit:  opts.limit,
		Offset: opts.offset,
	}))
	if err != nil {
		return fmt.Errorf("failed to call ListTags method: %w", err)
	}

	for _, t := range resp.Msg.Tags {
		if t == nil {
			continue
		}
		fmt.Fprintf(out, "%s %s\n", t.Id, t.Name)
	}
	return nil
}
