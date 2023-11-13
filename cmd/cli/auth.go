package main

import "github.com/spf13/cobra"

func newCmdAuth() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "auth",
		Short: "Authenticate Simoom",
	}
	cmd.AddCommand(newCmdAuthRefresh())
	cmd.AddCommand(newCmdAuthSignin())
	cmd.AddCommand(newCmdAuthSignup())
	return cmd
}
