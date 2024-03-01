// Package task パッケージはtaskサブコマンドを定義する
package task

import (
	"github.com/minguu42/simoom/cli/cmdutil"
	"github.com/spf13/cobra"
)

func NewCmdTask(core cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "task <command>",
		Short: "Manage tasks",
	}
	cmd.AddCommand(newCmdTaskCreate(core))
	cmd.AddCommand(newCmdTaskDelete(core))
	cmd.AddCommand(newCmdTaskEdit(core))
	return cmd
}
