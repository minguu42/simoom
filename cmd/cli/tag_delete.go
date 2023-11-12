package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newCmdTagDelete() *cobra.Command {
	return &cobra.Command{
		Use:   "delete",
		Short: "Delete a tag",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runTagDelete()
		},
	}
}

func runTagDelete() error {
	fmt.Println("runTagDelete executed")
	return nil
}
