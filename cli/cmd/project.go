package cmd

import "github.com/spf13/cobra"

func NewCmdProject() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "project <command>",
		Short: "Manage projects",
	}
	cmd.AddCommand(
		NewCmdProjectCreate(),
		NewCmdProjectDelete(),
		NewCmdProjectEdit(),
		NewCmdProjectList(),
	)
	return cmd
}
