package cmd

import (
	"context"
	"errors"
	"fmt"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/cli/api"
	"github.com/minguu42/simoom/cli/cmdutil"
	"github.com/minguu42/simoom/cli/factory"
	"github.com/minguu42/simoom/lib/go/simoompb/v1"
	"github.com/spf13/cobra"
)

type AuthRefreshOpts struct {
	profile string
	client  api.Client

	refreshToken string
}

func NewCmdAuthRefresh() *cobra.Command {
	var opts AuthRefreshOpts
	cmd := &cobra.Command{
		Use:   "refresh",
		Short: "Refresh the access token",
		RunE: func(cmd *cobra.Command, args []string) error {
			f := factory.FromContext(cmd.Context())
			opts.profile = f.Profile
			opts.client = f.Client

			if opts.refreshToken == "" {
				if opts.client.GetRefreshToken() == "" {
					return errors.New("refresh token is required")
				}
				opts.refreshToken = opts.client.GetRefreshToken()
			}
			return AuthRefreshRun(cmd.Context(), opts)
		},
	}
	cmdutil.DisableAuthCheck(cmd)

	cmd.Flags().StringVar(&opts.refreshToken, "refresh-token", "", "refresh token")
	return cmd
}

func AuthRefreshRun(ctx context.Context, opts AuthRefreshOpts) error {
	resp, err := opts.client.RefreshToken(ctx, connect.NewRequest(&simoompb.RefreshTokenRequest{
		RefreshToken: opts.refreshToken,
	}))
	if err != nil {
		return fmt.Errorf("failed to call RefreshToken method: %w", err)
	}
	fmt.Println("Successfully authenticated.")

	if err := api.SaveCredentials(opts.profile, resp.Msg.AccessToken, resp.Msg.RefreshToken); err != nil {
		return fmt.Errorf("failed to write credentials: %w", err)
	}
	return nil
}
