// Package auth パッケージはauthサブコマンドを定義する
package auth

import (
	"github.com/minguu42/simoom/cli/cmdutil"
	"github.com/spf13/cobra"
)

func NewCmdAuth(core cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "auth <command>",
		Short: "Authenticate simoom",
	}
	cmd.AddCommand(newCmdAuthRefresh(core))
	cmd.AddCommand(newCmdAuthSignin(core))
	cmd.AddCommand(newCmdAuthSignup(core))
	return cmd
}
