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

type ProjectListOpts struct {
	client api.Client

	limit  uint64
	offset uint64
}

func NewCmdProjectList() *cobra.Command {
	var opts ProjectListOpts
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List the projects",
		Args:    cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, _ []string) error {
			f := factory.FromContext(cmd.Context())
			opts.client = f.Client

			return ProjectListRun(cmd.Context(), f.Out, opts)
		},
	}
	cmd.Flags().Uint64Var(&opts.limit, "limit", 10, "limit")
	cmd.Flags().Uint64Var(&opts.offset, "offset", 0, "offset")
	return cmd
}

func ProjectListRun(ctx context.Context, out io.Writer, opts ProjectListOpts) error {
	resp, err := opts.client.ListProjects(ctx, connect.NewRequest(&simoompb.ListProjectsRequest{
		Limit:  opts.limit,
		Offset: opts.offset,
	}))
	if err != nil {
		return fmt.Errorf("failed to call ListProjects method: %w", err)
	}

	for _, p := range resp.Msg.Projects {
		if p == nil {
			continue
		}
		fmt.Fprintf(out, "%s %s\n", p.Id, p.Name)
	}
	return nil
}
