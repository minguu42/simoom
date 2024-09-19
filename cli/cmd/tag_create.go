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

type TagCreateOpts struct {
	Client api.Client

	Name string
}

func NewCmdTagCreate() *cobra.Command {
	var opts TagCreateOpts
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a tag",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, _ []string) error {
			f := factory.FromContext(cmd.Context())
			opts.Client = f.Client

			if opts.Name == "" {
				opts.Name = f.Prompter.Input("Tag name")
			}
			return TagCreateRun(cmd.Context(), f.Out, opts)
		},
	}
	cmd.Flags().StringVar(&opts.Name, "name", "", "tag name")
	return cmd
}

func TagCreateRun(ctx context.Context, out io.Writer, opts TagCreateOpts) error {
	resp, err := opts.Client.CreateTag(ctx, connect.NewRequest(&simoompb.CreateTagRequest{
		Name: opts.Name,
	}))
	if err != nil {
		return fmt.Errorf("failed to call CreateTag method: %w", err)
	}

	fmt.Fprintf(out, "Tag %s (%s) created\n", resp.Msg.Name, resp.Msg.Id)
	return nil
}
