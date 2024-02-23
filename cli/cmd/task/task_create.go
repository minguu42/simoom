package task

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/cli/cmdutil"
	"github.com/minguu42/simoom/lib/go/simoompb/v1"
	"github.com/spf13/cobra"
)

type taskCreateOpts struct {
	projectID string
	name      string
	priority  uint32
}

func newCmdTaskCreate(core cmdutil.Core) *cobra.Command {
	var opts taskCreateOpts
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a task",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, _ []string) error {
			if opts.projectID == "" {
				return fmt.Errorf("project-id is required")
			}
			if opts.name == "" {
				return fmt.Errorf("name is required")
			}
			if opts.priority == 99 {
				return fmt.Errorf("priority is required")
			}
			return runTaskCreate(cmd.Context(), core, opts)
		},
	}

	cmd.Flags().StringVar(&opts.projectID, "project-id", "", "project id")
	cmd.Flags().StringVar(&opts.name, "name", "", "task name")
	cmd.Flags().Uint32Var(&opts.priority, "priority", 99, "task priority")

	return cmd
}

func runTaskCreate(ctx context.Context, core cmdutil.Core, opts taskCreateOpts) error {
	req := connect.NewRequest(&simoompb.CreateTaskRequest{
		ProjectId: opts.projectID,
		Name:      opts.name,
		Priority:  opts.priority,
	})
	req.Header().Set("Authorization", fmt.Sprintf("Bearer %s", core.Credentials.AccessToken))

	resp, err := core.Client.CreateTask(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to call CreateTask method: %w", err)
	}

	if err := cmdutil.PrintJSON(resp.Msg); err != nil {
		return fmt.Errorf("failed to print json output: %w", err)
	}
	return nil
}
