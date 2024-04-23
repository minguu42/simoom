package cmd

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/cli/api"
	"github.com/minguu42/simoom/cli/cmdutil"
	"github.com/minguu42/simoom/cli/factory"
	"github.com/minguu42/simoom/lib/go/simoompb/v1"
	"github.com/spf13/cobra"
)

type taskListOpts struct {
	client api.Client

	limit     uint64
	offset    uint64
	projectID string
	tagID     string
}

func newCmdTaskList() *cobra.Command {
	var opts taskListOpts
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List the tasks",
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			f := factory.Value(cmd.Context())
			opts.client = f.Client
			return runTaskList(cmd.Context(), opts)
		},
	}
	cmd.Flags().Uint64Var(&opts.limit, "limit", 10, "Maximum number of tasks to fetch")
	cmd.Flags().Uint64Var(&opts.offset, "offset", 0, "offset")
	cmd.Flags().StringVar(&opts.projectID, "project-id", "", "Filter by project")
	cmd.Flags().StringVar(&opts.tagID, "tag-id", "", "Filter by tag")
	return cmd
}

func runTaskList(ctx context.Context, opts taskListOpts) error {
	var projectID *string
	if opts.projectID != "" {
		projectID = &opts.projectID
	}
	var tagID *string
	if opts.tagID != "" {
		tagID = &opts.tagID
	}
	resp, err := opts.client.ListTasks(ctx, connect.NewRequest(&simoompb.ListTasksRequest{
		Limit:     opts.limit,
		Offset:    opts.offset,
		ProjectId: projectID,
		TagId:     tagID,
	}))
	if err != nil {
		return fmt.Errorf("failed to call ListTasks method: %w", err)
	}

	if err := cmdutil.PrintJSON(resp.Msg); err != nil {
		return fmt.Errorf("failed to print json output: %w", err)
	}
	return nil
}
