package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newCmdAuthSignin() *cobra.Command {
	return &cobra.Command{
		Use:   "signin",
		Short: "Sign in to Simoom",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runAuthSignin()
		},
	}
}

func runAuthSignin() error {
	fmt.Println("runAuthSignin executed")
	return nil
}
