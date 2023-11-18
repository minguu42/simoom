package main

import (
	"context"
	"encoding/json"
	"fmt"

	"connectrpc.com/connect"
	"github.com/cockroachdb/errors"
	"github.com/minguu42/simoom/cli/cmdutil"
	"github.com/minguu42/simoom/library/simoompb/v1"
	"github.com/spf13/cobra"
)

type authSigninOpts struct {
	email    string
	password string
}

func newCmdAuthSignin(core cmdutil.Core) *cobra.Command {
	opts := authSigninOpts{}
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
			return runAuthSignin(cmd.Context(), core, opts)
		},
	}

	cmd.Flags().StringVar(&opts.email, "email", "", "email")
	cmd.Flags().StringVar(&opts.password, "password", "", "password")

	return cmd
}

func runAuthSignin(ctx context.Context, core cmdutil.Core, opts authSigninOpts) error {
	resp, err := core.Client.SignIn(ctx, connect.NewRequest(&simoompb.simoompb{
		Email:    opts.email,
		Password: opts.password,
	}))
	if err != nil {
		return fmt.Errorf("failed to call SignIn method. %w", err)
	}

	data, err := json.MarshalIndent(resp.Msg, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal: %w", err)
	}
	fmt.Println(string(data))
	return nil
}
