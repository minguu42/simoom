package main

import "github.com/spf13/cobra"

func newCmdTag() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "step",
		Short: "Work with steps",
	}
	cmd.AddCommand(newCmdTagCreate())
	cmd.AddCommand(newCmdTagDelete())
	cmd.AddCommand(newCmdTagEdit())
	cmd.AddCommand(newCmdTagList())
	return cmd
}
