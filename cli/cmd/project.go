package cmd

import (
	"github.com/minguu42/simoom/cli/cmdutil"
	"github.com/spf13/cobra"
)

func newCmdProject(f cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "project <command>",
		Short: "Manage projects",
	}
	cmd.AddCommand(newCmdProjectCreate(f))
	cmd.AddCommand(newCmdProjectDelete(f))
	cmd.AddCommand(newCmdProjectEdit(f))
	cmd.AddCommand(newCmdProjectList(f))
	cmd.AddCommand(newCmdProjectView(f))
	return cmd
}
