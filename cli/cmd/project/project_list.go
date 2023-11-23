package project

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/cli/cmdutil"
	"github.com/minguu42/simoom/pkg/simoompb/v1"
	"github.com/spf13/cobra"
)

type projectListOpts struct {
	limit  uint64
	offset uint64
}

func newCmdProjectList(core cmdutil.Core) *cobra.Command {
	var opts projectListOpts
	cmd := &cobra.Command{
		Use:     "list",
		Short:   "List the projects",
		Aliases: []string{"ls"},
		Args:    cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runProjectList(cmd.Context(), core, opts)
		},
	}

	cmd.Flags().Uint64Var(&opts.limit, "limit", 10, "limit")
	cmd.Flags().Uint64Var(&opts.offset, "offset", 0, "offset")

	return cmd
}

func runProjectList(ctx context.Context, core cmdutil.Core, opts projectListOpts) error {
	req := connect.NewRequest(&simoompb.ListProjectsRequest{
		Limit:  opts.limit,
		Offset: opts.offset,
	})
	req.Header().Set("Authorization", fmt.Sprintf("Bearer %s", core.Credentials.AccessToken))

	resp, err := core.Client.ListProjects(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to call ListProjects method: %w", err)
	}

	if err := cmdutil.PrintJSON(resp.Msg); err != nil {
		return fmt.Errorf("failed to print json output: %w", err)
	}
	return nil
}
