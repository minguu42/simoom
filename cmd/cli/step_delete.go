package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newCmdStepDelete() *cobra.Command {
	return &cobra.Command{
		Use:   "delete",
		Short: "Delete a step",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runStepDelete()
		},
	}
}

func runStepDelete() error {
	fmt.Println("runStepDelete executed")
	return nil
}
