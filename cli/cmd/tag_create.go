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

type TagCreateOpts struct {
	client api.Client

	name string
}

func NewCmdTagCreate() *cobra.Command {
	var opts TagCreateOpts
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a tag",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, _ []string) error {
			f := factory.FromContext(cmd.Context())
			opts.client = f.Client

			if opts.name == "" {
				return fmt.Errorf("name is required")
			}
			return TagCreateRun(cmd.Context(), opts)
		},
	}
	cmd.Flags().StringVar(&opts.name, "name", "", "tag name")
	return cmd
}

func TagCreateRun(ctx context.Context, opts TagCreateOpts) error {
	resp, err := opts.client.CreateTag(ctx, connect.NewRequest(&simoompb.CreateTagRequest{
		Name: opts.name,
	}))
	if err != nil {
		return fmt.Errorf("failed to call CreateTag method: %w", err)
	}

	if err := cmdutil.PrintJSON(resp.Msg); err != nil {
		return fmt.Errorf("failed to print json output: %w", err)
	}
	return nil
}
