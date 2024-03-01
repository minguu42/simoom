package task

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/cli/cmdutil"
	"github.com/minguu42/simoom/lib/go/simoompb/v1"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type taskEditOpts struct {
	id        string
	name      string
	content   string
	priority  uint32
	completed bool
}

func newCmdTaskEdit(core cmdutil.Factory) *cobra.Command {
	var opts taskEditOpts
	cmd := &cobra.Command{
		Use:   "edit",
		Short: "Edit a task",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if args[0] == "" {
				return fmt.Errorf("id is required")
			}
			opts.id = args[0]
			return runTaskEdit(cmd.Context(), core, opts)
		},
	}

	cmd.Flags().StringVar(&opts.name, "name", "", "task name")
	cmd.Flags().StringVar(&opts.content, "content", "", "task content")
	cmd.Flags().Uint32Var(&opts.priority, "priority", 99, "task priority")
	cmd.Flags().BoolVar(&opts.completed, "completed", false, "whether to complete the task")

	return cmd
}

func runTaskEdit(ctx context.Context, core cmdutil.Factory, opts taskEditOpts) error {
	var name *string
	if opts.name != "" {
		name = &opts.name
	}
	var content *string
	if opts.content != "" {
		content = &opts.content
	}
	var priority *uint32
	if opts.priority != 99 {
		priority = &opts.priority
	}
	var completedAt *timestamppb.Timestamp
	if opts.completed {
		completedAt = timestamppb.Now()
	}
	req := connect.NewRequest(&simoompb.UpdateTaskRequest{
		Id:          opts.id,
		Name:        name,
		Content:     content,
		Priority:    priority,
		CompletedAt: completedAt,
	})
	req.Header().Set("Authorization", fmt.Sprintf("Bearer %s", core.Credentials.AccessToken))

	resp, err := core.Client.UpdateTask(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to call UpdateTask method: %w", err)
	}

	if err := cmdutil.PrintJSON(resp.Msg); err != nil {
		return fmt.Errorf("failed to print json output: %w", err)
	}
	return nil
}
