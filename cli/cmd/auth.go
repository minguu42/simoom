package cmd

import "github.com/spf13/cobra"

func NewCmdAuth() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "auth <command>",
		Short: "Authenticate simoom",
	}
	cmd.AddCommand(
		NewCmdAuthRefresh(),
		NewCmdAuthSignin(),
		NewCmdAuthSignup(),
	)
	return cmd
}
