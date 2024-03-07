package cmd

import "github.com/spf13/cobra"

func NewCmdTag() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tag <command>",
		Short: "Manage tags",
	}
	cmd.AddCommand(
		newCmdTagCreate(),
		newCmdTagDelete(),
		newCmdTagEdit(),
		newCmdTagList(),
	)
	return cmd
}
