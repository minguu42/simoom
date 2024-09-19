package cmd

import (
	"context"
	"fmt"
	"io"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/cli/api"
	"github.com/minguu42/simoom/cli/cmdutil"
	"github.com/minguu42/simoom/cli/factory"
	"github.com/minguu42/simoom/lib/go/simoompb/v1"
	"github.com/spf13/cobra"
)

type AuthRefreshOpts struct {
	Profile string
	Client  api.Client

	RefreshToken string
}

func NewCmdAuthRefresh() *cobra.Command {
	var opts AuthRefreshOpts
	cmd := &cobra.Command{
		Use:   "refresh",
		Short: "Refresh the access token",
		RunE: func(cmd *cobra.Command, args []string) error {
			f := factory.FromContext(cmd.Context())
			opts.Profile = f.Profile
			opts.Client = f.Client

			if opts.RefreshToken == "" {
				if opts.Client.GetRefreshToken() == "" {
					opts.RefreshToken = f.Prompter.Input("Refresh token")
				} else {
					opts.RefreshToken = opts.Client.GetRefreshToken()
				}
			}
			return AuthRefreshRun(cmd.Context(), f.Out, opts)
		},
	}
	cmdutil.DisableAuthCheck(cmd)

	cmd.Flags().StringVar(&opts.RefreshToken, "refresh-token", "", "refresh token")
	return cmd
}

func AuthRefreshRun(ctx context.Context, out io.Writer, opts AuthRefreshOpts) error {
	resp, err := opts.Client.RefreshToken(ctx, connect.NewRequest(&simoompb.RefreshTokenRequest{
		RefreshToken: opts.RefreshToken,
	}))
	if err != nil {
		return fmt.Errorf("failed to call RefreshToken method: %w", err)
	}
	fmt.Fprintln(out, "Successfully authenticated")

	if err := api.SaveCredentials(opts.Profile, resp.Msg.AccessToken, resp.Msg.RefreshToken); err != nil {
		return fmt.Errorf("failed to write credentials: %w", err)
	}
	return nil
}
