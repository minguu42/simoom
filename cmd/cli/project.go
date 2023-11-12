package main

import "github.com/spf13/cobra"

func newCmdProject() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "project",
		Short: "Work with projects.",
	}
	cmd.AddCommand(newCmdProjectCreate())
	cmd.AddCommand(newCmdProjectDelete())
	cmd.AddCommand(newCmdProjectEdit())
	cmd.AddCommand(newCmdProjectList())
	return cmd
}
