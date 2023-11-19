package main

import (
	"github.com/minguu42/simoom/cli/cmdutil"
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
	cmd.AddCommand(newCmdStep(core))
	cmd.AddCommand(newCmdTag(core))
	cmd.AddCommand(newCmdTask(core))
	return cmd
}
