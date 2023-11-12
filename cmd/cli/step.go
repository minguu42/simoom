package main

import "github.com/spf13/cobra"

func newCmdStep() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "step",
		Short: "Work with steps",
	}
	cmd.AddCommand(newCmdStepCreate())
	cmd.AddCommand(newCmdStepDelete())
	cmd.AddCommand(newCmdStepEdit())
	return cmd
}
