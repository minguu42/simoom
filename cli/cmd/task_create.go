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

type TaskCreateOpts struct {
	Client api.Client

	ProjectID string
	Name      string
	Priority  uint32
}

func NewCmdTaskCreate() *cobra.Command {
	var opts TaskCreateOpts
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a task",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, _ []string) error {
			f := factory.FromContext(cmd.Context())
			opts.Client = f.Client

			return TaskCreateRun(cmd.Context(), f.Out, opts)
		},
	}
	cmd.Flags().StringVar(&opts.ProjectID, "project-id", "", "project id")
	cmd.Flags().StringVar(&opts.Name, "name", "", "task name")
	cmd.Flags().Uint32Var(&opts.Priority, "priority", 0, "task priority")
	_ = cmd.MarkFlagRequired("project-id")
	_ = cmd.MarkFlagRequired("name")
	_ = cmd.MarkFlagRequired("priority")
	return cmd
}

func TaskCreateRun(ctx context.Context, out io.Writer, opts TaskCreateOpts) error {
	resp, err := opts.Client.CreateTask(ctx, connect.NewRequest(&simoompb.CreateTaskRequest{
		ProjectId: opts.ProjectID,
		Name:      opts.Name,
		Priority:  opts.Priority,
	}))
	if err != nil {
		return fmt.Errorf("failed to call CreateTask method: %w", err)
	}

	fmt.Fprintf(out, "Task %s (%s) created\n", resp.Msg.Name, resp.Msg.Id)
	return nil
}
