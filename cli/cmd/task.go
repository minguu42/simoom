package cmd

import "github.com/spf13/cobra"

func NewCmdTask() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "task <command>",
		Short: "Manage tasks",
	}
	cmd.AddCommand(
		newCmdTaskCreate(),
		newCmdTaskDelete(),
		newCmdTaskEdit(),
		newCmdTaskList(),
	)
	return cmd
}
