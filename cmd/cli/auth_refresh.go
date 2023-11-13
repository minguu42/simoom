package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newCmdAuthRefresh() *cobra.Command {
	return &cobra.Command{
		Use:   "refresh",
		Short: "Refresh the access token",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runAuthRefresh()
		},
	}
}

func runAuthRefresh() error {
	fmt.Println("runAuthRefresh executed")
	return nil
}
