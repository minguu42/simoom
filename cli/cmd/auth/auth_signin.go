package auth

import (
	"context"
	"errors"
	"fmt"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/cli/cmdutil"
	"github.com/minguu42/simoom/lib/go/simoompb/v1"
	"github.com/minguu42/simoom/lib/go/simoompb/v1/simoompbconnect"
	"github.com/spf13/cobra"
)

type authSigninOpts struct {
	client simoompbconnect.SimoomServiceClient

	email    string
	password string
}

func newCmdAuthSignin(f cmdutil.Factory) *cobra.Command {
	opts := authSigninOpts{
		client: f.Client,
	}
	cmd := &cobra.Command{
		Use:   "signin",
		Short: "Sign in to Simoom",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.email == "" {
				return errors.New("email is required")
			}
			if opts.password == "" {
				return errors.New("password is required")
			}
			return runAuthSignin(cmd.Context(), opts)
		},
	}
	cmdutil.DisableAuthCheck(cmd)

	cmd.Flags().StringVar(&opts.email, "email", "", "email")
	cmd.Flags().StringVar(&opts.password, "password", "", "password")

	return cmd
}

func runAuthSignin(ctx context.Context, opts authSigninOpts) error {
	resp, err := opts.client.SignIn(ctx, connect.NewRequest(&simoompb.SignInRequest{
		Email:    opts.email,
		Password: opts.password,
	}))
	if err != nil {
		return fmt.Errorf("failed to call SignIn method. %w", err)
	}
	fmt.Println("Successfully authenticated.")

	if err := cmdutil.WriteCredentials(resp.Msg.AccessToken, resp.Msg.RefreshToken); err != nil {
		return fmt.Errorf("failed to write credentials: %w", err)
	}
	return nil
}
