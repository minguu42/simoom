package main

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/cli/cmdutil"
	"github.com/minguu42/simoom/lib/simoompb/v1"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type stepEditOpts struct {
	id        string
	title     string
	completed bool
}

func newCmdStepEdit(core cmdutil.Core) *cobra.Command {
	var opts stepEditOpts
	cmd := &cobra.Command{
		Use:   "edit",
		Short: "Edit a step",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.id = args[0]
			return runStepEdit(cmd.Context(), core, opts)
		},
	}

	cmd.Flags().StringVar(&opts.title, "title", "", "step title")
	cmd.Flags().BoolVar(&opts.completed, "completed", false, "completed")

	return cmd
}

func runStepEdit(ctx context.Context, core cmdutil.Core, opts stepEditOpts) error {
	var title *string
	if opts.title != "" {
		title = &opts.title
	}
	var completedAt *timestamppb.Timestamp
	if opts.completed {
		completedAt = timestamppb.Now()
	}
	req := connect.NewRequest(&simoompb.simoompb{
		Id:          opts.id,
		Title:       title,
		CompletedAt: completedAt,
	})
	req.Header().Set("Authorization", fmt.Sprintf("Bearer %s", core.Credentials.AccessToken))

	resp, err := core.Client.UpdateStep(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to call UpdateStep method: %w", err)
	}

	if err := cmdutil.PrintJSON(resp.Msg); err != nil {
		return fmt.Errorf("failed to print json output: %w", err)
	}
	return nil
}
