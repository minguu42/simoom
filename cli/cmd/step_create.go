package cmd

import (
	"context"
	"errors"
	"fmt"
	"io"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/cli/api"
	"github.com/minguu42/simoom/cli/factory"
	"github.com/minguu42/simoom/lib/go/simoompb/v1"
	"github.com/spf13/cobra"
)

type StepCreateOpts struct {
	Client api.Client

	TaskID string
	Name   string
}

func NewCmdStepCreate() *cobra.Command {
	var opts StepCreateOpts
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a step",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, _ []string) error {
			f := factory.FromContext(cmd.Context())
			opts.Client = f.Client

			if opts.TaskID == "" {
				return errors.New("task-id is required")
			}
			if opts.Name == "" {
				return errors.New("name is required")
			}
			return StepCreateRun(cmd.Context(), f.Out, opts)
		},
	}
	cmd.Flags().StringVar(&opts.TaskID, "task-id", "", "task id")
	cmd.Flags().StringVar(&opts.Name, "name", "", "step name")
	_ = cmd.MarkFlagRequired("task-id")
	_ = cmd.MarkFlagRequired("name")
	return cmd
}

func StepCreateRun(ctx context.Context, out io.Writer, opts StepCreateOpts) error {
	resp, err := opts.Client.CreateStep(ctx, connect.NewRequest(&simoompb.CreateStepRequest{
		TaskId: opts.TaskID,
		Name:   opts.Name,
	}))
	if err != nil {
		return fmt.Errorf("failed to call CreateStep method: %w", err)
	}

	fmt.Fprintf(out, "Step %s (%s) created\n", resp.Msg.Name, resp.Msg.Id)
	return nil
}
