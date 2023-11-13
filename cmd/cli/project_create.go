package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newCmdProjectCreate() *cobra.Command {
	return &cobra.Command{
		Use:   "create",
		Short: "Create a project",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runProjectCreate()
		},
	}
}

func runProjectCreate() error {
	fmt.Println("runProjectCreate executed")
	return nil
}
