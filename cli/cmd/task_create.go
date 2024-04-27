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

type TaskCreateOpts struct {
	client api.Client

	projectID string
	name      string
	priority  uint32
}

func NewCmdTaskCreate() *cobra.Command {
	var opts TaskCreateOpts
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a task",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, _ []string) error {
			f := factory.FromContext(cmd.Context())
			opts.client = f.Client

			if opts.projectID == "" {
				return fmt.Errorf("project-id is required")
			}
			if opts.name == "" {
				return fmt.Errorf("name is required")
			}
			if opts.priority == 99 {
				return fmt.Errorf("priority is required")
			}
			return TaskCreateRun(cmd.Context(), opts)
		},
	}
	cmd.Flags().StringVar(&opts.projectID, "project-id", "", "project id")
	cmd.Flags().StringVar(&opts.name, "name", "", "task name")
	cmd.Flags().Uint32Var(&opts.priority, "priority", 99, "task priority")
	return cmd
}

func TaskCreateRun(ctx context.Context, opts TaskCreateOpts) error {
	resp, err := opts.client.CreateTask(ctx, connect.NewRequest(&simoompb.CreateTaskRequest{
		ProjectId: opts.projectID,
		Name:      opts.name,
		Priority:  opts.priority,
	}))
	if err != nil {
		return fmt.Errorf("failed to call CreateTask method: %w", err)
	}

	if err := cmdutil.PrintJSON(resp.Msg); err != nil {
		return fmt.Errorf("failed to print json output: %w", err)
	}
	return nil
}
