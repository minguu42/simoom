package main

import (
	"github.com/minguu42/simoom/cmd/cli/cmdutil"
	"github.com/spf13/cobra"
)

func newCmdAuth(core cmdutil.Core) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "auth",
		Short: "Authenticate Simoom",
	}
	cmd.AddCommand(newCmdAuthRefresh(core))
	cmd.AddCommand(newCmdAuthSignin(core))
	cmd.AddCommand(newCmdAuthSignup(core))
	return cmd
}
