package step

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/cli/cmdutil"
	"github.com/minguu42/simoom/lib/go/simoompb/v1"
	"github.com/minguu42/simoom/lib/go/simoompb/v1/simoompbconnect"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type stepEditOpts struct {
	client      simoompbconnect.SimoomServiceClient
	credentials cmdutil.Credentials

	id        string
	name      string
	completed bool
}

func newCmdStepEdit(f cmdutil.Factory) *cobra.Command {
	opts := stepEditOpts{
		client:      f.Client,
		credentials: f.Credentials,
	}
	cmd := &cobra.Command{
		Use:   "edit",
		Short: "Edit a step",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.id = args[0]
			return runStepEdit(cmd.Context(), opts)
		},
	}

	cmd.Flags().StringVar(&opts.name, "name", "", "step name")
	cmd.Flags().BoolVar(&opts.completed, "completed", false, "completed")

	return cmd
}

func runStepEdit(ctx context.Context, opts stepEditOpts) error {
	var name *string
	if opts.name != "" {
		name = &opts.name
	}
	var completedAt *timestamppb.Timestamp
	if opts.completed {
		completedAt = timestamppb.Now()
	}
	req := connect.NewRequest(&simoompb.UpdateStepRequest{
		Id:          opts.id,
		Name:        name,
		CompletedAt: completedAt,
	})
	req.Header().Set("Authorization", fmt.Sprintf("Bearer %s", opts.credentials.AccessToken))
	resp, err := opts.client.UpdateStep(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to call UpdateStep method: %w", err)
	}

	if err := cmdutil.PrintJSON(resp.Msg); err != nil {
		return fmt.Errorf("failed to print json output: %w", err)
	}
	return nil
}
