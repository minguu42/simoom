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
	profile string
	client  api.Client

	name     string
	email    string
	password string
}

func NewCmdAuthSignup() *cobra.Command {
	var opts AuthSignupOpts
	cmd := &cobra.Command{
		Use:   "signup",
		Short: "Sign up to Simoom",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			f := factory.FromContext(cmd.Context())
			opts.profile = f.Profile
			opts.client = f.Client

			return AuthSignupRun(cmd.Context(), f.Out, opts)
		},
	}
	cmdutil.DisableAuthCheck(cmd)

	cmd.Flags().StringVar(&opts.name, "name", "", "username")
	cmd.Flags().StringVar(&opts.email, "email", "", "email")
	cmd.Flags().StringVar(&opts.password, "password", "", "password")
	_ = cmd.MarkFlagRequired("name")
	_ = cmd.MarkFlagRequired("email")
	_ = cmd.MarkFlagRequired("password")
	return cmd
}

func AuthSignupRun(ctx context.Context, out io.Writer, opts AuthSignupOpts) error {
	resp, err := opts.client.SignUp(ctx, connect.NewRequest(&simoompb.SignUpRequest{
		Name:     opts.name,
		Email:    opts.email,
		Password: opts.password,
	}))
	if err != nil {
		return fmt.Errorf("failed to call SignUp method: %w", err)
	}
	fmt.Fprintln(out, "Successfully authenticated")

	if err := api.SaveCredentials(opts.profile, resp.Msg.AccessToken, resp.Msg.RefreshToken); err != nil {
		return fmt.Errorf("failed to write credentials: %w", err)
	}
	return nil
}
