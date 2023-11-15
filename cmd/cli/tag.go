package main

import (
	"github.com/minguu42/simoom/cmd/cli/cmdutil"
	"github.com/spf13/cobra"
)

func newCmdTag(core cmdutil.Core) *cobra.Command {
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
