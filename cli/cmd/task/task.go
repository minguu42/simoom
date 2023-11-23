// Package task パッケージはtaskサブコマンドを定義する
package task

import (
	"github.com/minguu42/simoom/cli/cmdutil"
	"github.com/spf13/cobra"
)

func NewCmdTask(core cmdutil.Core) *cobra.Command {
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
