package main

import (
	"github.com/minguu42/simoom/cmd/cli/cmdutil"
	"github.com/spf13/cobra"
)

func newCmdTask(core cmdutil.Core) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "task",
		Short: "Work with tasks",
	}
	cmd.AddCommand(newCmdTaskCreate(core))
	cmd.AddCommand(newCmdTaskDelete(core))
	cmd.AddCommand(newCmdTaskEdit(core))
	cmd.AddCommand(newCmdTaskList(core))
	return cmd
}
