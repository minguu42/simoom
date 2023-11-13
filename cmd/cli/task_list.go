package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newCmdTaskList() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List the tasks",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runTaskList()
		},
	}
}

func runTaskList() error {
	fmt.Println("runTaskList executed")
	return nil
}
