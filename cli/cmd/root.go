// Package cmd パッケージはコマンドを定義する
package cmd

import (
	"errors"

	"github.com/minguu42/simoom/cli/cmdutil"
	"github.com/spf13/cobra"
)

func NewCmdRoot(f cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "simoom <command> <subcommand> [flags]",
		Short: "Simoom CLI",
		Long:  `Work seamlessly with Simoom from the command line.`,
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			if cmdutil.IsAuthCheckEnabled(cmd) && !f.Client.CheckCredentials() {
				return errors.New("authentication failed")
			}
			return nil
		},
	}
	cmd.AddCommand(newCmdAuth(f))
	cmd.AddCommand(newCmdProject(f))
	cmd.AddCommand(newCmdStep(f))
	cmd.AddCommand(newCmdTag(f))
	cmd.AddCommand(newCmdTask(f))
	return cmd
}
