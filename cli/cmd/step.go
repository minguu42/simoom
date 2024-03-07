package cmd

import "github.com/spf13/cobra"

func NewCmdStep() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "step <command>",
		Short: "Manage steps",
	}
	cmd.AddCommand(
		newCmdStepCreate(),
		newCmdStepDelete(),
		newCmdStepEdit(),
	)
	return cmd
}
