package task

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/cli/cmdutil"
	"github.com/minguu42/simoom/pkg/simoompb/v1"
	"github.com/spf13/cobra"
)

type taskListOpts struct {
	projectID string
	limit     uint64
	offset    uint64
}

func newCmdTaskList(core cmdutil.Core) *cobra.Command {
	var opts taskListOpts
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List the tasks",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, _ []string) error {
			if opts.projectID == "" {
				return fmt.Errorf("project-id is required")
			}
			return runTaskList(cmd.Context(), core, opts)
		},
	}

	cmd.Flags().StringVar(&opts.projectID, "project-id", "", "project id")
	cmd.Flags().Uint64Var(&opts.limit, "limit", 10, "limit")
	cmd.Flags().Uint64Var(&opts.offset, "offset", 0, "offset")

	return cmd
}

func runTaskList(ctx context.Context, core cmdutil.Core, opts taskListOpts) error {
	req := connect.NewRequest(&simoompb.ListTasksByProjectIDRequest{
		ProjectId: opts.projectID,
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
