package cmd

import (
	"github.com/minguu42/simoom/cli/cmdutil"
	"github.com/spf13/cobra"
)

func newCmdStep(f cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "step <command>",
		Short: "Manage steps",
	}
	cmd.AddCommand(newCmdStepCreate(f))
	cmd.AddCommand(newCmdStepDelete(f))
	cmd.AddCommand(newCmdStepEdit(f))
	return cmd
}
