package task

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/cli/cmdutil"
	"github.com/minguu42/simoom/pkg/simoompb/v1"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type taskEditOpts struct {
	id        string
	title     string
	content   string
	priority  uint32
	completed bool
}

func newCmdTaskEdit(core cmdutil.Core) *cobra.Command {
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

	cmd.Flags().StringVar(&opts.title, "title", "", "task title")
	cmd.Flags().StringVar(&opts.content, "content", "", "task content")
	cmd.Flags().Uint32Var(&opts.priority, "priority", 99, "task priority")
	cmd.Flags().BoolVar(&opts.completed, "completed", false, "whether to complete the task")

	return cmd
}

func runTaskEdit(ctx context.Context, core cmdutil.Core, opts taskEditOpts) error {
	var title *string
	if opts.title != "" {
		title = &opts.title
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
		Name:        title,
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
