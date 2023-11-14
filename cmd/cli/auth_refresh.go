package main

import (
	"context"
	"encoding/json"
	"fmt"

	"connectrpc.com/connect"
	"github.com/cockroachdb/errors"
	"github.com/minguu42/simoom/cmd/cli/cmdutil"
	"github.com/minguu42/simoom/gen/simoompb/v1"
	"github.com/spf13/cobra"
)

type authRefreshOpts struct {
	refreshToken string
}

func newCmdAuthRefresh(core cmdutil.Core) *cobra.Command {
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

func runAuthRefresh(ctx context.Context, core cmdutil.Core, opts authRefreshOpts) error {
	resp, err := core.Client.RefreshAccessToken(ctx, connect.NewRequest(&simoompb.RefreshAccessTokenRequest{
		RefreshToken: opts.refreshToken,
	}))
	if err != nil {
		return fmt.Errorf("failed to call RefreshAccessToken method. %w", err)
	}

	data, err := json.MarshalIndent(resp.Msg, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal: %w", err)
	}
	fmt.Println(string(data))
	return nil
}
