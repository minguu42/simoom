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
	Client api.Client

	ID        string
	Name      string
	Completed bool
}

func NewCmdStepEdit() *cobra.Command {
	var opts StepEditOpts
	cmd := &cobra.Command{
		Use:   "edit",
		Short: "Edit a step",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			f := factory.FromContext(cmd.Context())
			opts.Client = f.Client

			opts.ID = args[0]
			return StepEditRun(cmd.Context(), f.Out, opts)
		},
	}
	cmd.Flags().StringVar(&opts.Name, "name", "", "step name")
	cmd.Flags().BoolVar(&opts.Completed, "completed", false, "completed")
	return cmd
}

func StepEditRun(ctx context.Context, out io.Writer, opts StepEditOpts) error {
	var name *string
	if opts.Name != "" {
		name = &opts.Name
	}
	var completedAt *timestamppb.Timestamp
	if opts.Completed {
		completedAt = timestamppb.Now()
	}
	resp, err := opts.Client.UpdateStep(ctx, connect.NewRequest(&simoompb.UpdateStepRequest{
		Id:          opts.ID,
		Name:        name,
		CompletedAt: completedAt,
	}))
	if err != nil {
		return fmt.Errorf("failed to call UpdateStep method: %w", err)
	}

	fmt.Fprintf(out, "Step %s (%s) edited\n", resp.Msg.Name, resp.Msg.Id)
	return nil
}
