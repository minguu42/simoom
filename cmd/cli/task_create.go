package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newCmdTaskCreate() *cobra.Command {
	return &cobra.Command{
		Use:   "create",
		Short: "Create a task",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runTaskCreate()
		},
	}
}

func runTaskCreate() error {
	fmt.Println("runTaskCreate executed")
	return nil
}
