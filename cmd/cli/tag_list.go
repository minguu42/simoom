package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newCmdTagList() *cobra.Command {
	return &cobra.Command{
		Use:     "list",
		Short:   "List the tags",
		Aliases: []string{"ls"},
		RunE: func(cmd *cobra.Command, args []string) error {
			return runTagList()
		},
	}
}

func runTagList() error {
	fmt.Println("runTagList executed")
	return nil
}
