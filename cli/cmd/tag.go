package cmd

import (
	"github.com/minguu42/simoom/cli/cmdutil"
	"github.com/spf13/cobra"
)

func newCmdTag(f cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tag <command>",
		Short: "Manage tags",
	}
	cmd.AddCommand(newCmdTagCreate(f))
	cmd.AddCommand(newCmdTagDelete(f))
	cmd.AddCommand(newCmdTagEdit(f))
	cmd.AddCommand(newCmdTagList(f))
	return cmd
}
