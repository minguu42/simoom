package cmd

import (
	"github.com/minguu42/simoom/cli/cmdutil"
	"github.com/spf13/cobra"
)

func NewCmdTask(f cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "task <command>",
		Short: "Manage tasks",
	}
	cmd.AddCommand(newCmdTaskCreate(f))
	cmd.AddCommand(newCmdTaskDelete(f))
	cmd.AddCommand(newCmdTaskEdit(f))
	return cmd
}
