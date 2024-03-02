package project

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/cli/cmdutil"
	"github.com/minguu42/simoom/lib/go/simoompb/v1"
	"github.com/minguu42/simoom/lib/go/simoompb/v1/simoompbconnect"
	"github.com/spf13/cobra"
)

type projectListOpts struct {
	client      simoompbconnect.SimoomServiceClient
	credentials cmdutil.Credentials

	limit  uint64
	offset uint64
}

func newCmdProjectList(f cmdutil.Factory) *cobra.Command {
	opts := projectListOpts{
		client:      f.Client,
		credentials: f.Credentials,
	}
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List the projects",
		Args:    cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runProjectList(cmd.Context(), opts)
		},
	}

	cmd.Flags().Uint64Var(&opts.limit, "limit", 10, "limit")
	cmd.Flags().Uint64Var(&opts.offset, "offset", 0, "offset")

	return cmd
}

func runProjectList(ctx context.Context, opts projectListOpts) error {
	req := connect.NewRequest(&simoompb.ListProjectsRequest{
		Limit:  opts.limit,
		Offset: opts.offset,
	})
	req.Header().Set("Authorization", fmt.Sprintf("Bearer %s", opts.credentials.AccessToken))
	resp, err := opts.client.ListProjects(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to call ListProjects method: %w", err)
	}

	if err := cmdutil.PrintJSON(resp.Msg); err != nil {
		return fmt.Errorf("failed to print json output: %w", err)
	}
	return nil
}
