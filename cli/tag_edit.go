package main

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/cli/cmdutil"
	"github.com/minguu42/simoom/library/simoompb/v1"
	"github.com/spf13/cobra"
)

type tagEditOpts struct {
	id   string
	name string
}

func newCmdTagEdit(core cmdutil.Core) *cobra.Command {
	var opts tagEditOpts
	cmd := &cobra.Command{
		Use:   "edit",
		Short: "Edit a tag",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.id = args[0]
			return runTagEdit(cmd.Context(), core, opts)
		},
	}

	cmd.Flags().StringVar(&opts.name, "name", "", "tag name")

	return cmd
}

func runTagEdit(ctx context.Context, core cmdutil.Core, opts tagEditOpts) error {
	var name *string
	if opts.name != "" {
		name = &opts.name
	}
	req := connect.NewRequest(&simoompb.simoompb{
		Id:   opts.id,
		Name: name,
	})
	req.Header().Set("Authorization", fmt.Sprintf("Bearer %s", core.Credentials.AccessToken))

	resp, err := core.Client.UpdateTag(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to call UpdateTag method: %w", err)
	}

	if err := cmdutil.PrintJSON(resp.Msg); err != nil {
		return fmt.Errorf("failed to print json output: %w", err)
	}
	return nil
}
