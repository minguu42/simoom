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

type authSigninOpts struct {
	profile string
	client  api.Client

	email    string
	password string
}

func newCmdAuthSignin() *cobra.Command {
	var opts authSigninOpts
	cmd := &cobra.Command{
		Use:   "signin",
		Short: "Sign in to Simoom",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			f := factory.FromContext(cmd.Context())
			opts.profile = f.Profile
			opts.client = f.Client

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

	if err := api.SaveCredentials(opts.profile, resp.Msg.AccessToken, resp.Msg.RefreshToken); err != nil {
		return fmt.Errorf("failed to write credentials: %w", err)
	}
	return nil
}
