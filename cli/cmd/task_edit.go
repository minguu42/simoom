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
	Client api.Client

	ID        string
	Name      string
	Content   string
	Priority  uint32
	Completed bool
}

func NewCmdTaskEdit() *cobra.Command {
	var opts TaskEditOpts
	cmd := &cobra.Command{
		Use:   "edit",
		Short: "Edit a task",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			f := factory.FromContext(cmd.Context())
			opts.Client = f.Client

			if args[0] == "" {
				return fmt.Errorf("id is required")
			}
			opts.ID = args[0]
			return TaskEditRun(cmd.Context(), f.Out, opts)
		},
	}
	cmd.Flags().StringVar(&opts.Name, "name", "", "task name")
	cmd.Flags().StringVar(&opts.Content, "content", "", "task content")
	cmd.Flags().Uint32Var(&opts.Priority, "priority", 99, "task priority")
	cmd.Flags().BoolVar(&opts.Completed, "completed", false, "whether to complete the task")
	return cmd
}

func TaskEditRun(ctx context.Context, out io.Writer, opts TaskEditOpts) error {
	var name *string
	if opts.Name != "" {
		name = &opts.Name
	}
	var content *string
	if opts.Content != "" {
		content = &opts.Content
	}
	var priority *uint32
	if opts.Priority != 99 {
		priority = &opts.Priority
	}
	var completedAt *timestamppb.Timestamp
	if opts.Completed {
		completedAt = timestamppb.Now()
	}
	resp, err := opts.Client.UpdateTask(ctx, connect.NewRequest(&simoompb.UpdateTaskRequest{
		Id:          opts.ID,
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
