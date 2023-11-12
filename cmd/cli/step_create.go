package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newCmdStepCreate() *cobra.Command {
	return &cobra.Command{
		Use:   "create",
		Short: "Create a step",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runStepCreate()
		},
	}
}

func runStepCreate() error {
	fmt.Println("runStepCreate executed")
	return nil
}
