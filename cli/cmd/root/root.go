// Package root パッケージはルートコマンドを定義する
package root

import (
	"github.com/minguu42/simoom/cli/cmd/auth"
	"github.com/minguu42/simoom/cli/cmd/project"
	"github.com/minguu42/simoom/cli/cmd/step"
	"github.com/minguu42/simoom/cli/cmd/tag"
	"github.com/minguu42/simoom/cli/cmd/task"
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
	cmd.AddCommand(auth.NewCmdAuth(core))
	cmd.AddCommand(project.NewCmdProject(core))
	cmd.AddCommand(step.NewCmdStep(core))
	cmd.AddCommand(tag.NewCmdTag(core))
	cmd.AddCommand(task.NewCmdTask(core))
	return cmd
}
