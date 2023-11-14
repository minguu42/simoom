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
	cmd.AddCommand(newCmdProjectDelete())
	cmd.AddCommand(newCmdProjectEdit())
	cmd.AddCommand(newCmdProjectList(core))
	return cmd
}
