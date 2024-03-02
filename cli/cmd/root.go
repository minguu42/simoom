// Package cmd パッケージはコマンドを定義する
package cmd

import (
	"github.com/minguu42/simoom/cli/cmdutil"
	"github.com/spf13/cobra"
)

func NewCmdRoot(core cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "simoom <command> <subcommand> [flags]",
		Short: "Simoom CLI",
		Long:  `Work seamlessly with Simoom from the command line.`,
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			if cmdutil.IsAuthCheckEnabled(cmd) {
				// TODO: 認証情報をチェックし、認証情報が正しくない場合はエラーを返す
				return nil
			}
			return nil
		},
	}
	cmd.AddCommand(newCmdAuth(core))
	cmd.AddCommand(newCmdProject(core))
	cmd.AddCommand(newCmdStep(core))
	cmd.AddCommand(newCmdTag(core))
	cmd.AddCommand(newCmdTask(core))
	return cmd
}
