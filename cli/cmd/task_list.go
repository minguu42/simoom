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

type TaskListOpts struct {
	client api.Client

	limit     uint64
	offset    uint64
	projectID string
	tagID     string
}

func NewCmdTaskList() *cobra.Command {
	var opts TaskListOpts
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List the tasks",
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			f := factory.FromContext(cmd.Context())
			opts.client = f.Client
			return TaskListRun(cmd.Context(), f.Out, opts)
		},
	}
	cmd.Flags().Uint64Var(&opts.limit, "limit", 10, "Maximum number of tasks to fetch")
	cmd.Flags().Uint64Var(&opts.offset, "offset", 0, "offset")
	cmd.Flags().StringVar(&opts.projectID, "project-id", "", "Filter by project")
	cmd.Flags().StringVar(&opts.tagID, "tag-id", "", "Filter by tag")
	return cmd
}

func TaskListRun(ctx context.Context, out io.Writer, opts TaskListOpts) error {
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

	for _, t := range resp.Msg.Tasks {
		if t == nil {
			continue
		}
		fmt.Fprintf(out, " %s %s\n", t.Id, t.Name)
	}
	return nil
}
