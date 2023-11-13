package main

import "github.com/spf13/cobra"

func newCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "simoom",
		Short: "Simoom CLI",
		Long:  `Work seamlessly with Simoom from the command line.`,
	}
	cmd.AddCommand(newCmdAuth())
	cmd.AddCommand(newCmdProject())
	cmd.AddCommand(newCmdStep())
	cmd.AddCommand(newCmdTag())
	cmd.AddCommand(newCmdTask())
	return cmd
}
