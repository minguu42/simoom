package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newCmdProjectEdit() *cobra.Command {
	return &cobra.Command{
		Use:   "edit",
		Short: "Edit a project",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runProjectEdit()
		},
	}
}

func runProjectEdit() error {
	fmt.Println("runProjectEdit executed")
	return nil
}
