package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newCmdAuthSignup() *cobra.Command {
	return &cobra.Command{
		Use:   "signup",
		Short: "Sign up to Simoom",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runAuthSignup()
		},
	}
}

func runAuthSignup() error {
	fmt.Println("runAuthSignup executed")
	return nil
}
