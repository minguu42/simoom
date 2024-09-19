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

type AuthSigninOpts struct {
	Profile string
	Client  api.Client

	Email    string
	Password string
}

func NewCmdAuthSignin() *cobra.Command {
	var opts AuthSigninOpts
	cmd := &cobra.Command{
		Use:   "signin",
		Short: "Sign in to Simoom",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			f := factory.FromContext(cmd.Context())
			opts.Profile = f.Profile
			opts.Client = f.Client

			if opts.Email == "" {
				opts.Email = f.Prompter.Input("Email")
			}
			if opts.Password == "" {
				opts.Password = f.Prompter.Input("Password")
			}
			return AuthSigninRun(cmd.Context(), f.Out, opts)
		},
	}
	cmdutil.DisableAuthCheck(cmd)

	cmd.Flags().StringVar(&opts.Email, "email", "", "email")
	cmd.Flags().StringVar(&opts.Password, "password", "", "password")
	return cmd
}

func AuthSigninRun(ctx context.Context, out io.Writer, opts AuthSigninOpts) error {
	resp, err := opts.Client.SignIn(ctx, connect.NewRequest(&simoompb.SignInRequest{
		Email:    opts.Email,
		Password: opts.Password,
	}))
	if err != nil {
		return fmt.Errorf("failed to call SignIn method. %w", err)
	}
	fmt.Fprintln(out, "Successfully authenticated")

	if err := api.SaveCredentials(opts.Profile, resp.Msg.AccessToken, resp.Msg.RefreshToken); err != nil {
		return fmt.Errorf("failed to write credentials: %w", err)
	}
	return nil
}
