package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newCmdStepEdit() *cobra.Command {
	return &cobra.Command{
		Use:   "edit",
		Short: "Edit a step",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runStepEdit()
		},
	}
}

func runStepEdit() error {
	fmt.Println("runStepEdit executed")
	return nil
}
