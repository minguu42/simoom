package cmd

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/cli/api"
	"github.com/minguu42/simoom/cli/cmdutil"
	"github.com/minguu42/simoom/lib/go/simoompb/v1"
	"github.com/spf13/cobra"
)

type tagEditOpts struct {
	client api.Client

	id   string
	name string
}

func newCmdTagEdit(f cmdutil.Factory) *cobra.Command {
	opts := tagEditOpts{
		client: f.Client,
	}
	cmd := &cobra.Command{
		Use:   "edit",
		Short: "Edit a tag",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.id = args[0]
			return runTagEdit(cmd.Context(), opts)
		},
	}

	cmd.Flags().StringVar(&opts.name, "name", "", "tag name")

	return cmd
}

func runTagEdit(ctx context.Context, opts tagEditOpts) error {
	var name *string
	if opts.name != "" {
		name = &opts.name
	}
	resp, err := opts.client.UpdateTag(ctx, connect.NewRequest(&simoompb.UpdateTagRequest{
		Id:   opts.id,
		Name: name,
	}))
	if err != nil {
		return fmt.Errorf("failed to call UpdateTag method: %w", err)
	}

	if err := cmdutil.PrintJSON(resp.Msg); err != nil {
		return fmt.Errorf("failed to print json output: %w", err)
	}
	return nil
}
