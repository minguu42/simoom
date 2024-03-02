package tag

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/cli/cmdutil"
	"github.com/minguu42/simoom/lib/go/simoompb/v1"
	"github.com/minguu42/simoom/lib/go/simoompb/v1/simoompbconnect"
	"github.com/spf13/cobra"
)

type tagCreateOpts struct {
	client      simoompbconnect.SimoomServiceClient
	credentials cmdutil.Credentials

	name string
}

func newCmdTagCreate(f cmdutil.Factory) *cobra.Command {
	opts := tagCreateOpts{
		client:      f.Client,
		credentials: f.Credentials,
	}
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a tag",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, _ []string) error {
			if opts.name == "" {
				return fmt.Errorf("name is required")
			}
			return runTagCreate(cmd.Context(), opts)
		},
	}

	cmd.Flags().StringVar(&opts.name, "name", "", "tag name")

	return cmd
}

func runTagCreate(ctx context.Context, opts tagCreateOpts) error {
	req := connect.NewRequest(&simoompb.CreateTagRequest{
		Name: opts.name,
	})
	req.Header().Set("Authorization", fmt.Sprintf("Bearer %s", opts.credentials.AccessToken))
	resp, err := opts.client.CreateTag(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to call CreateTag method: %w", err)
	}

	if err := cmdutil.PrintJSON(resp.Msg); err != nil {
		return fmt.Errorf("failed to print json output: %w", err)
	}
	return nil
}
