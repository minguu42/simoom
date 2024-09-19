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

type AuthSignupOpts struct {
	Profile string
	Client  api.Client

	Name     string
	Email    string
	Password string
}

func NewCmdAuthSignup() *cobra.Command {
	var opts AuthSignupOpts
	cmd := &cobra.Command{
		Use:   "signup",
		Short: "Sign up to Simoom",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			f := factory.FromContext(cmd.Context())
			opts.Profile = f.Profile
			opts.Client = f.Client

			if opts.Name == "" {
				opts.Name = f.Prompter.Input("Username")
			}
			if opts.Email == "" {
				opts.Email = f.Prompter.Input("Email")
			}
			if opts.Password == "" {
				opts.Password = f.Prompter.Input("Password")
			}
			return AuthSignupRun(cmd.Context(), f.Out, opts)
		},
	}
	cmdutil.DisableAuthCheck(cmd)

	cmd.Flags().StringVar(&opts.Name, "name", "", "username")
	cmd.Flags().StringVar(&opts.Email, "email", "", "email")
	cmd.Flags().StringVar(&opts.Password, "password", "", "password")
	return cmd
}

func AuthSignupRun(ctx context.Context, out io.Writer, opts AuthSignupOpts) error {
	resp, err := opts.Client.SignUp(ctx, connect.NewRequest(&simoompb.SignUpRequest{
		Name:     opts.Name,
		Email:    opts.Email,
		Password: opts.Password,
	}))
	if err != nil {
		return fmt.Errorf("failed to call SignUp method: %w", err)
	}
	fmt.Fprintln(out, "Successfully authenticated")

	if err := api.SaveCredentials(opts.Profile, resp.Msg.AccessToken, resp.Msg.RefreshToken); err != nil {
		return fmt.Errorf("failed to write credentials: %w", err)
	}
	return nil
}
