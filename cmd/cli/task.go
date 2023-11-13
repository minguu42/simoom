package main

import "github.com/spf13/cobra"

func newCmdTask() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "task",
		Short: "Work with tasks",
	}
	cmd.AddCommand(newCmdTaskCreate())
	cmd.AddCommand(newCmdTaskDelete())
	cmd.AddCommand(newCmdTaskEdit())
	cmd.AddCommand(newCmdTaskList())
	return cmd
}
