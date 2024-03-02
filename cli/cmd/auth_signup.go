package cmd

import (
	"context"
	"errors"
	"fmt"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/cli/api"
	"github.com/minguu42/simoom/cli/cmdutil"
	"github.com/minguu42/simoom/lib/go/simoompb/v1"
	"github.com/spf13/cobra"
)

type authSignupOpts struct {
	client api.Client

	name     string
	email    string
	password string
}

func newCmdAuthSignup(f cmdutil.Factory) *cobra.Command {
	opts := authSignupOpts{
		client: f.Client,
	}
	cmd := &cobra.Command{
		Use:   "signup",
		Short: "Sign up to Simoom",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.name == "" {
				return errors.New("name is required")
			}
			if opts.email == "" {
				return errors.New("email is required")
			}
			if opts.password == "" {
				return errors.New("password is required")
			}
			return runAuthSignup(cmd.Context(), opts)
		},
	}
	cmdutil.DisableAuthCheck(cmd)

	cmd.Flags().StringVar(&opts.name, "name", "", "username")
	cmd.Flags().StringVar(&opts.email, "email", "", "email")
	cmd.Flags().StringVar(&opts.password, "password", "", "password")

	return cmd
}

func runAuthSignup(ctx context.Context, opts authSignupOpts) error {
	resp, err := opts.client.SignUp(ctx, connect.NewRequest(&simoompb.SignUpRequest{
		Name:     opts.name,
		Email:    opts.email,
		Password: opts.password,
	}))
	if err != nil {
		return fmt.Errorf("failed to call SignUp method: %w", err)
	}
	fmt.Println("Successfully authenticated.")

	if err := api.WriteCredentials(resp.Msg.AccessToken, resp.Msg.RefreshToken); err != nil {
		return fmt.Errorf("failed to write credentials: %w", err)
	}
	return nil
}
