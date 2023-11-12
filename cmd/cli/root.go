package main

import "github.com/spf13/cobra"

func newCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "simoom",
		Short: "Simoom CLI",
		Long:  `Work seamlessly with Simoom from the command line.`,
	}
	cmd.AddCommand(newCmdProject())
	cmd.AddCommand(newCmdTask())
	return cmd
}
