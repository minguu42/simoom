package cmd

import (
	"github.com/minguu42/simoom/cli/cmdutil"
	"github.com/spf13/cobra"
)

func NewCmdAuth(f cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "auth <command>",
		Short: "Authenticate simoom",
	}
	cmd.AddCommand(newCmdAuthRefresh(f))
	cmd.AddCommand(newCmdAuthSignin(f))
	cmd.AddCommand(newCmdAuthSignup(f))
	return cmd
}
