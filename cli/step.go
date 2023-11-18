package main

import (
	"github.com/minguu42/simoom/cli/cmdutil"
	"github.com/spf13/cobra"
)

func newCmdStep(core cmdutil.Core) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "step",
		Short: "Work with steps",
	}
	cmd.AddCommand(newCmdStepCreate(core))
	cmd.AddCommand(newCmdStepDelete(core))
	cmd.AddCommand(newCmdStepEdit(core))
	return cmd
}
