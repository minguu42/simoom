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
	"google.golang.org/protobuf/types/known/timestamppb"
)

type TaskEditOpts struct {
	client api.Client

	id        string
	name      string
	content   string
	priority  uint32
	completed bool
}

func NewCmdTaskEdit() *cobra.Command {
	var opts TaskEditOpts
	cmd := &cobra.Command{
		Use:   "edit",
		Short: "Edit a task",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			f := factory.FromContext(cmd.Context())
			opts.client = f.Client

			if args[0] == "" {
				return fmt.Errorf("id is required")
			}
			opts.id = args[0]
			return TaskEditRun(cmd.Context(), f.Out, opts)
		},
	}
	cmd.Flags().StringVar(&opts.name, "name", "", "task name")
	cmd.Flags().StringVar(&opts.content, "content", "", "task content")
	cmd.Flags().Uint32Var(&opts.priority, "priority", 99, "task priority")
	cmd.Flags().BoolVar(&opts.completed, "completed", false, "whether to complete the task")
	return cmd
}

func TaskEditRun(ctx context.Context, out io.Writer, opts TaskEditOpts) error {
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
	resp, err := opts.client.UpdateTask(ctx, connect.NewRequest(&simoompb.UpdateTaskRequest{
		Id:          opts.id,
		Name:        name,
		Content:     content,
		Priority:    priority,
		CompletedAt: completedAt,
	}))
	if err != nil {
		return fmt.Errorf("failed to call UpdateTask method: %w", err)
	}

	fmt.Fprintf(out, "Task %s (%s) edited\n", resp.Msg.Name, resp.Msg.Id)
	return nil
}
