package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newCmdTaskEdit() *cobra.Command {
	return &cobra.Command{
		Use:   "edit",
		Short: "Edit a task",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runTaskEdit()
		},
	}
}

func runTaskEdit() error {
	fmt.Println("runTaskEdit executed")
	return nil
}
