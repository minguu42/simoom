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

type StepEditOpts struct {
	client api.Client

	id        string
	name      string
	completed bool
}

func NewCmdStepEdit() *cobra.Command {
	var opts StepEditOpts
	cmd := &cobra.Command{
		Use:   "edit",
		Short: "Edit a step",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			f := factory.FromContext(cmd.Context())
			opts.client = f.Client

			opts.id = args[0]
			return StepEditRun(cmd.Context(), f.Out, opts)
		},
	}
	cmd.Flags().StringVar(&opts.name, "name", "", "step name")
	cmd.Flags().BoolVar(&opts.completed, "completed", false, "completed")
	return cmd
}

func StepEditRun(ctx context.Context, out io.Writer, opts StepEditOpts) error {
	var name *string
	if opts.name != "" {
		name = &opts.name
	}
	var completedAt *timestamppb.Timestamp
	if opts.completed {
		completedAt = timestamppb.Now()
	}
	resp, err := opts.client.UpdateStep(ctx, connect.NewRequest(&simoompb.UpdateStepRequest{
		Id:          opts.id,
		Name:        name,
		CompletedAt: completedAt,
	}))
	if err != nil {
		return fmt.Errorf("failed to call UpdateStep method: %w", err)
	}

	fmt.Fprintf(out, "Step %s (%s) edited\n", resp.Msg.Name, resp.Msg.Id)
	return nil
}
