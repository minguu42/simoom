package cmd

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/cli/api"
	"github.com/minguu42/simoom/cli/cmdutil"
	"github.com/minguu42/simoom/cli/factory"
	"github.com/minguu42/simoom/lib/go/simoompb/v1"
	"github.com/spf13/cobra"
)

type projectEditOpts struct {
	client api.Client

	id         string
	name       string
	color      string
	isArchived bool
}

func newCmdProjectEdit() *cobra.Command {
	var opts projectEditOpts
	cmd := &cobra.Command{
		Use:   "edit",
		Short: "Edit a project",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			f := factory.FromContext(cmd.Context())
			opts.client = f.Client

			opts.id = args[0]
			return runProjectEdit(cmd.Context(), opts)
		},
	}

	cmd.Flags().StringVar(&opts.name, "name", "", "project name")
	cmd.Flags().StringVar(&opts.color, "color", "", "project color")
	cmd.Flags().BoolVar(&opts.isArchived, "archived", false, "whether to archive the project")

	return cmd
}

func runProjectEdit(ctx context.Context, opts projectEditOpts) error {
	var name *string
	if opts.name != "" {
		name = &opts.name
	}
	var color *string
	if opts.color != "" {
		color = &opts.color
	}
	var isArchived *bool
	if opts.isArchived {
		isArchived = &opts.isArchived
	}
	resp, err := opts.client.UpdateProject(ctx, connect.NewRequest(&simoompb.UpdateProjectRequest{
		Id:         opts.id,
		Name:       name,
		Color:      color,
		IsArchived: isArchived,
	}))
	if err != nil {
		return fmt.Errorf("failed to call UpdateProject method: %w", err)
	}

	if err := cmdutil.PrintJSON(resp.Msg); err != nil {
		return fmt.Errorf("failed to print json output: %w", err)
	}
	return nil
}
