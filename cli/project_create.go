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

type projectCreateOpts struct {
	name  string
	color string
}

func newCmdProjectCreate(core cmdutil.Core) *cobra.Command {
	var opts projectCreateOpts
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a project",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, _ []string) error {
			if opts.name == "" {
				return errors.New("name is required")
			}
			if opts.color == "" {
				return errors.New("color is required")
			}
			return runProjectCreate(cmd.Context(), core, opts)
		},
	}

	cmd.Flags().StringVar(&opts.name, "name", "", "project name")
	cmd.Flags().StringVar(&opts.color, "color", "", "project color")

	return cmd
}

func runProjectCreate(ctx context.Context, core cmdutil.Core, opts projectCreateOpts) error {
	req := connect.NewRequest(&simoompb.simoompb{
		Name:  opts.name,
		Color: opts.color,
	})
	req.Header().Set("Authorization", fmt.Sprintf("Bearer %s", core.Credentials.AccessToken))

	resp, err := core.Client.CreateProject(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to call CreateProject method: %w", err)
	}

	if err := cmdutil.PrintJSON(resp.Msg); err != nil {
		return fmt.Errorf("failed to print json output: %w", err)
	}
	return nil
}
