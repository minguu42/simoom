package project

import (
	"context"
	"errors"
	"fmt"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/cli/cmdutil"
	"github.com/minguu42/simoom/lib/go/simoompb/v1"
	"github.com/spf13/cobra"
)

type projectViewOpts struct {
	id     string
	limit  uint64
	offset uint64
}

func newCmdProjectView(core cmdutil.Core) *cobra.Command {
	var opts projectViewOpts
	cmd := &cobra.Command{
		Use:   "view <id> [flags]",
		Short: "List the tasks included in that project",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.id = args[0]
			if len(opts.id) != 26 {
				return errors.New("id is a 26-character string")
			}
			return runProjectView(cmd.Context(), core, opts)
		},
	}

	cmd.Flags().Uint64Var(&opts.limit, "limit", 10, "limit")
	cmd.Flags().Uint64Var(&opts.offset, "offset", 0, "offset")

	return cmd
}

func runProjectView(ctx context.Context, core cmdutil.Core, opts projectViewOpts) error {
	req := connect.NewRequest(&simoompb.ListTasksByProjectIDRequest{
		ProjectId: opts.id,
		Limit:     opts.limit,
		Offset:    opts.offset,
	})
	req.Header().Set("Authorization", fmt.Sprintf("Bearer %s", core.Credentials.AccessToken))

	resp, err := core.Client.ListTasksByProjectID(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to call ListTasksByProjectID method: %w", err)
	}

	if err := cmdutil.PrintJSON(resp.Msg); err != nil {
		return fmt.Errorf("failed to print json output: %w", err)
	}
	return nil
}
