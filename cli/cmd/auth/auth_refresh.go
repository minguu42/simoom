package auth

import (
	"context"
	"errors"
	"fmt"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/cli/cmdutil"
	"github.com/minguu42/simoom/lib/go/simoompb/v1"
	"github.com/spf13/cobra"
)

type authRefreshOpts struct {
	refreshToken string
}

func newCmdAuthRefresh(core cmdutil.Factory) *cobra.Command {
	opts := authRefreshOpts{}
	cmd := &cobra.Command{
		Use:   "refresh",
		Short: "Refresh the access token",
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.refreshToken == "" {
				return errors.New("refresh-token is required")
			}
			return runAuthRefresh(cmd.Context(), core, opts)
		},
	}

	cmd.Flags().StringVar(&opts.refreshToken, "refresh-token", "", "refresh token")

	return cmd
}

func runAuthRefresh(ctx context.Context, core cmdutil.Factory, opts authRefreshOpts) error {
	resp, err := core.Client.RefreshToken(ctx, connect.NewRequest(&simoompb.RefreshTokenRequest{
		RefreshToken: opts.refreshToken,
	}))
	if err != nil {
		return fmt.Errorf("failed to call RefreshToken method: %w", err)
	}
	fmt.Println("Successfully authenticated.")

	if err := cmdutil.WriteCredentials(resp.Msg.AccessToken, resp.Msg.RefreshToken); err != nil {
		return fmt.Errorf("failed to write credentials: %w", err)
	}
	return nil
}
