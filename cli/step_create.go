package main

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"github.com/cockroachdb/errors"
	"github.com/minguu42/simoom/cli/cmdutil"
	"github.com/minguu42/simoom/library/simoompb/v1"
	"github.com/spf13/cobra"
)

type stepCreateOpts struct {
	taskID string
	title  string
}

func newCmdStepCreate(core cmdutil.Core) *cobra.Command {
	var opts stepCreateOpts
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a step",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, _ []string) error {
			if opts.taskID == "" {
				return errors.New("task-id is required")
			}
			if opts.title == "" {
				return errors.New("title is required")
			}
			return runStepCreate(cmd.Context(), core, opts)
		},
	}

	cmd.Flags().StringVar(&opts.taskID, "task-id", "", "task id")
	cmd.Flags().StringVar(&opts.title, "title", "", "step title")

	return cmd
}

func runStepCreate(ctx context.Context, core cmdutil.Core, opts stepCreateOpts) error {
	req := connect.NewRequest(&simoompb.simoompb{
		TaskId: opts.taskID,
		Title:  opts.title,
	})
	req.Header().Set("Authorization", fmt.Sprintf("Bearer %s", core.Credentials.AccessToken))

	resp, err := core.Client.CreateStep(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to call CreateStep method: %w", err)
	}

	if err := cmdutil.PrintJSON(resp.Msg); err != nil {
		return fmt.Errorf("failed to print json output: %w", err)
	}
	return nil
}
