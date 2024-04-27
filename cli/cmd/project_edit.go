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
)

type ProjectEditOpts struct {
	Client api.Client

	ID         string
	Name       string
	Color      string
	IsArchived bool
}

func NewCmdProjectEdit() *cobra.Command {
	var opts ProjectEditOpts
	cmd := &cobra.Command{
		Use:   "edit",
		Short: "Edit a project",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			f := factory.FromContext(cmd.Context())
			opts.Client = f.Client

			opts.ID = args[0]
			return ProjectEditRun(cmd.Context(), f.Out, opts)
		},
	}
	cmd.Flags().StringVar(&opts.Name, "name", "", "project name")
	cmd.Flags().StringVar(&opts.Color, "color", "", "project color")
	cmd.Flags().BoolVar(&opts.IsArchived, "archived", false, "whether to archive the project")
	return cmd
}

func ProjectEditRun(ctx context.Context, out io.Writer, opts ProjectEditOpts) error {
	var name *string
	if opts.Name != "" {
		name = &opts.Name
	}
	var color *string
	if opts.Color != "" {
		color = &opts.Color
	}
	var isArchived *bool
	if opts.IsArchived {
		isArchived = &opts.IsArchived
	}
	resp, err := opts.Client.UpdateProject(ctx, connect.NewRequest(&simoompb.UpdateProjectRequest{
		Id:         opts.ID,
		Name:       name,
		Color:      color,
		IsArchived: isArchived,
	}))
	if err != nil {
		return fmt.Errorf("failed to call UpdateProject method: %w", err)
	}

	fmt.Fprintf(out, "Project %s (%s) edited\n", resp.Msg.Name, resp.Msg.Id)
	return nil
}
