package cmd

import (
	"context"
	"errors"
	"fmt"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/cli/api"
	"github.com/minguu42/simoom/cli/cmdutil"
	"github.com/minguu42/simoom/cli/factory"
	"github.com/minguu42/simoom/lib/go/simoompb/v1"
	"github.com/spf13/cobra"
)

type ProjectCreateOpts struct {
	client api.Client

	name  string
	color string
}

func NewCmdProjectCreate() *cobra.Command {
	var opts ProjectCreateOpts
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a project",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, _ []string) error {
			f := factory.FromContext(cmd.Context())
			opts.client = f.Client

			if opts.name == "" {
				return errors.New("name is required")
			}
			if opts.color == "" {
				return errors.New("color is required")
			}
			return ProjectCreateRun(cmd.Context(), opts)
		},
	}
	cmd.Flags().StringVar(&opts.name, "name", "", "project name")
	cmd.Flags().StringVar(&opts.color, "color", "", "project color")
	return cmd
}

func ProjectCreateRun(ctx context.Context, opts ProjectCreateOpts) error {
	resp, err := opts.client.CreateProject(ctx, connect.NewRequest(&simoompb.CreateProjectRequest{
		Name:  opts.name,
		Color: opts.color,
	}))
	if err != nil {
		return fmt.Errorf("failed to call CreateProject method: %w", err)
	}

	if err := cmdutil.PrintJSON(resp.Msg); err != nil {
		return fmt.Errorf("failed to print json output: %w", err)
	}
	return nil
}
