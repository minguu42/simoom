// Package tag パッケージはtagサブコマンドを定義する
package tag

import (
	"github.com/minguu42/simoom/cli/cmdutil"
	"github.com/spf13/cobra"
)

func NewCmdTag(core cmdutil.Core) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tag",
		Short: "Work with steps",
	}
	cmd.AddCommand(newCmdTagCreate(core))
	cmd.AddCommand(newCmdTagDelete(core))
	cmd.AddCommand(newCmdTagEdit(core))
	cmd.AddCommand(newCmdTagList(core))
	return cmd
}
