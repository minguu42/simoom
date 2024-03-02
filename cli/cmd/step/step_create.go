package step

import (
	"context"
	"errors"
	"fmt"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/cli/cmdutil"
	"github.com/minguu42/simoom/lib/go/simoompb/v1"
	"github.com/minguu42/simoom/lib/go/simoompb/v1/simoompbconnect"
	"github.com/spf13/cobra"
)

type stepCreateOpts struct {
	client      simoompbconnect.SimoomServiceClient
	credentials cmdutil.Credentials

	taskID string
	name   string
}

func newCmdStepCreate(f cmdutil.Factory) *cobra.Command {
	opts := stepCreateOpts{
		client:      f.Client,
		credentials: f.Credentials,
	}
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a step",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, _ []string) error {
			if opts.taskID == "" {
				return errors.New("task-id is required")
			}
			if opts.name == "" {
				return errors.New("name is required")
			}
			return runStepCreate(cmd.Context(), opts)
		},
	}

	cmd.Flags().StringVar(&opts.taskID, "task-id", "", "task id")
	cmd.Flags().StringVar(&opts.name, "name", "", "step name")

	return cmd
}

func runStepCreate(ctx context.Context, opts stepCreateOpts) error {
	req := connect.NewRequest(&simoompb.CreateStepRequest{
		TaskId: opts.taskID,
		Name:   opts.name,
	})
	req.Header().Set("Authorization", fmt.Sprintf("Bearer %s", opts.credentials.AccessToken))
	resp, err := opts.client.CreateStep(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to call CreateStep method: %w", err)
	}

	if err := cmdutil.PrintJSON(resp.Msg); err != nil {
		return fmt.Errorf("failed to print json output: %w", err)
	}
	return nil
}
