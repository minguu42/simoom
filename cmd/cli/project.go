package main

import (
	"github.com/minguu42/simoom/cmd/cli/cmdutil"
	"github.com/spf13/cobra"
)

func newCmdProject(core cmdutil.Core) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "project",
		Short: "Work with projects",
	}
	cmd.AddCommand(newCmdProjectCreate(core))
	cmd.AddCommand(newCmdProjectDelete(core))
	cmd.AddCommand(newCmdProjectEdit(core))
	cmd.AddCommand(newCmdProjectList(core))
	return cmd
}
