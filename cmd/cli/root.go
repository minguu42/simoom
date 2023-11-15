package main

import (
	"github.com/minguu42/simoom/cmd/cli/cmdutil"
	"github.com/spf13/cobra"
)

func newCmdRoot(core cmdutil.Core) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "simoom",
		Short: "Simoom CLI",
		Long:  `Work seamlessly with Simoom from the command line.`,
	}
	cmd.AddCommand(newCmdAuth(core))
	cmd.AddCommand(newCmdProject(core))
	cmd.AddCommand(newCmdStep())
	cmd.AddCommand(newCmdTag(core))
	cmd.AddCommand(newCmdTask())
	return cmd
}
