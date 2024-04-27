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

			return ProjectCreateRun(cmd.Context(), f.Out, opts)
		},
	}
	cmd.Flags().StringVar(&opts.name, "name", "", "project name")
	cmd.Flags().StringVar(&opts.color, "color", "", "project color")
	_ = cmd.MarkFlagRequired("name")
	_ = cmd.MarkFlagRequired("color")
	return cmd
}

func ProjectCreateRun(ctx context.Context, out io.Writer, opts ProjectCreateOpts) error {
	resp, err := opts.client.CreateProject(ctx, connect.NewRequest(&simoompb.CreateProjectRequest{
		Name:  opts.name,
		Color: opts.color,
	}))
	if err != nil {
		return fmt.Errorf("failed to call CreateProject method: %w", err)
	}

	fmt.Fprintf(out, "Project %s (%s) created\n", resp.Msg.Name, resp.Msg.Id)
	return nil
}
