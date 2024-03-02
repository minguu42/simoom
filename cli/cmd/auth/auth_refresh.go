package auth

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/cli/api"
	"github.com/minguu42/simoom/cli/cmdutil"
	"github.com/minguu42/simoom/lib/go/simoompb/v1"
	"github.com/minguu42/simoom/lib/go/simoompb/v1/simoompbconnect"
	"github.com/spf13/cobra"
)

type authRefreshOpts struct {
	client      simoompbconnect.SimoomServiceClient
	credentials cmdutil.Credentials

	refreshToken string
}

func newCmdAuthRefresh(f cmdutil.Factory) *cobra.Command {
	opts := authRefreshOpts{
		client:      f.Client,
		credentials: f.Credentials,
	}
	cmd := &cobra.Command{
		Use:   "refresh",
		Short: "Refresh the access token",
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.refreshToken != "" {
				opts.credentials.RefreshToken = opts.refreshToken
			}
			return runAuthRefresh(cmd.Context(), opts)
		},
	}

	cmd.Flags().StringVar(&opts.refreshToken, "refresh-token", "", "refresh token")

	return cmd
}

func runAuthRefresh(ctx context.Context, opts authRefreshOpts) error {
	resp, err := opts.client.RefreshToken(ctx, connect.NewRequest(&simoompb.RefreshTokenRequest{
		RefreshToken: opts.credentials.RefreshToken,
	}))
	if err != nil {
		return fmt.Errorf("failed to call RefreshToken method: %w", err)
	}
	fmt.Println("Successfully authenticated.")

	if err := api.WriteCredentials(resp.Msg.AccessToken, resp.Msg.RefreshToken); err != nil {
		return fmt.Errorf("failed to write credentials: %w", err)
	}
	return nil
}
