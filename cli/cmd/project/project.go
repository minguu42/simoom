// Package project パッケージはprojectサブコマンドを定義する
package project

import (
	"github.com/minguu42/simoom/cli/cmdutil"
	"github.com/spf13/cobra"
)

func NewCmdProject(core cmdutil.Core) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "project <command>",
		Short: "Manage projects",
	}
	cmd.AddCommand(newCmdProjectCreate(core))
	cmd.AddCommand(newCmdProjectDelete(core))
	cmd.AddCommand(newCmdProjectEdit(core))
	cmd.AddCommand(newCmdProjectList(core))
	cmd.AddCommand(newCmdProjectView(core))
	return cmd
}
