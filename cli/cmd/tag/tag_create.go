package tag

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/cli/cmdutil"
	"github.com/minguu42/simoom/pkg/simoompb/v1"
	"github.com/spf13/cobra"
)

type tagCreateOpts struct {
	name string
}

func newCmdTagCreate(core cmdutil.Core) *cobra.Command {
	var opts tagCreateOpts
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a tag",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, _ []string) error {
			if opts.name == "" {
				return fmt.Errorf("name is required")
			}
			return runTagCreate(cmd.Context(), core, opts)
		},
	}

	cmd.Flags().StringVar(&opts.name, "name", "", "tag name")

	return cmd
}

func runTagCreate(ctx context.Context, core cmdutil.Core, opts tagCreateOpts) error {
	req := connect.NewRequest(&simoompb.CreateTagRequest{
		Name: opts.name,
	})
	req.Header().Set("Authorization", fmt.Sprintf("Bearer %s", core.Credentials.AccessToken))
	resp, err := core.Client.CreateTag(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to call CreateTag method: %w", err)
	}

	if err := cmdutil.PrintJSON(resp.Msg); err != nil {
		return fmt.Errorf("failed to print json output: %w", err)
	}
	return nil
}
