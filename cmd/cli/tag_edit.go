package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newCmdTagEdit() *cobra.Command {
	return &cobra.Command{
		Use:   "edit",
		Short: "Edit a tag",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runTagEdit()
		},
	}
}

func runTagEdit() error {
	fmt.Println("runTagEdit executed")
	return nil
}
