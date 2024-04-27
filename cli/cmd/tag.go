package cmd

import "github.com/spf13/cobra"

func NewCmdTag() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tag <command>",
		Short: "Manage tags",
	}
	cmd.AddCommand(
		NewCmdTagCreate(),
		NewCmdTagDelete(),
		NewCmdTagEdit(),
		NewCmdTagList(),
	)
	return cmd
}
