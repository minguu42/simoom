package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newCmdTagCreate() *cobra.Command {
	return &cobra.Command{
		Use:   "create",
		Short: "Create a tag",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runTagCreate()
		},
	}
}

func runTagCreate() error {
	fmt.Println("runTagCreate executed")
	return nil
}
