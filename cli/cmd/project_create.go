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

type ProjectCreateOpts struct {
	Client api.Client

	Name  string
	Color string
}

func NewCmdProjectCreate() *cobra.Command {
	var opts ProjectCreateOpts
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a project",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, _ []string) error {
			f := factory.FromContext(cmd.Context())
			opts.Client = f.Client

			if opts.Name == "" {
				opts.Name = f.Prompter.Input("Project name")
			}
			return ProjectCreateRun(cmd.Context(), f.Out, opts)
		},
	}
	cmd.Flags().StringVar(&opts.Name, "name", "", "project name")
	cmd.Flags().StringVar(&opts.Color, "color", "#000000", "project color")
	return cmd
}

func ProjectCreateRun(ctx context.Context, out io.Writer, opts ProjectCreateOpts) error {
	resp, err := opts.Client.CreateProject(ctx, connect.NewRequest(&simoompb.CreateProjectRequest{
		Name:  opts.Name,
		Color: opts.Color,
	}))
	if err != nil {
		return fmt.Errorf("failed to call CreateProject method: %w", err)
	}

	fmt.Fprintf(out, "Project %s (%s) created\n", resp.Msg.Name, resp.Msg.Id)
	return nil
}
