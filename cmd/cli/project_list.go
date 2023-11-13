package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newCmdProjectList() *cobra.Command {
	return &cobra.Command{
		Use:     "list",
		Short:   "List the projects",
		Aliases: []string{"ls"},
		RunE: func(cmd *cobra.Command, args []string) error {
			return runProjectList()
		},
	}
}

func runProjectList() error {
	fmt.Println("runProjectList executed")
	return nil
}
