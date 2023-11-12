package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newCmdTaskDelete() *cobra.Command {
	return &cobra.Command{
		Use:   "delete",
		Short: "Delete a task",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runTaskDelete()
		},
	}
}

func runTaskDelete() error {
	fmt.Println("runTaskDelete executed")
	return nil
}
