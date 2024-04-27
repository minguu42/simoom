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
	Client api.Client

	Limit     uint64
	Offset    uint64
	ProjectID string
	TagID     string
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
			opts.Client = f.Client
			return TaskListRun(cmd.Context(), f.Out, opts)
		},
	}
	cmd.Flags().Uint64Var(&opts.Limit, "limit", 10, "Maximum number of tasks to fetch")
	cmd.Flags().Uint64Var(&opts.Offset, "offset", 0, "offset")
	cmd.Flags().StringVar(&opts.ProjectID, "project-id", "", "Filter by project")
	cmd.Flags().StringVar(&opts.TagID, "tag-id", "", "Filter by tag")
	return cmd
}

func TaskListRun(ctx context.Context, out io.Writer, opts TaskListOpts) error {
	var projectID *string
	if opts.ProjectID != "" {
		projectID = &opts.ProjectID
	}
	var tagID *string
	if opts.TagID != "" {
		tagID = &opts.TagID
	}
	resp, err := opts.Client.ListTasks(ctx, connect.NewRequest(&simoompb.ListTasksRequest{
		Limit:     opts.Limit,
		Offset:    opts.Offset,
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
		fmt.Fprintf(out, "%s %s\n", t.Id, t.Name)
	}
	return nil
}
