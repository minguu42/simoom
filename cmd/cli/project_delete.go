package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newCmdProjectDelete() *cobra.Command {
	return &cobra.Command{
		Use:   "delete",
		Short: "Delete a project",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runProjectDelete()
		},
	}
}

func runProjectDelete() error {
	fmt.Println("runProjectDelete executed")
	return nil
}
