// Package step パッケージはstepサブコマンドを定義する
package step

import (
	"github.com/minguu42/simoom/cli/cmdutil"
	"github.com/spf13/cobra"
)

func NewCmdStep(core cmdutil.Core) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "step <command>",
		Short: "Manage steps",
	}
	cmd.AddCommand(newCmdStepCreate(core))
	cmd.AddCommand(newCmdStepDelete(core))
	cmd.AddCommand(newCmdStepEdit(core))
	return cmd
}
