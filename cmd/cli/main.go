package main

import (
	"context"
	"fmt"
	"os"
)

type exitCode int

const (
	exitOK    exitCode = 0
	exitError exitCode = 1
)

func main() {
	code := mainRun()
	os.Exit(int(code))
}

func mainRun() exitCode {
	rootCmd := newCmdRoot()
	if err := rootCmd.ExecuteContext(context.Background()); err != nil {
		fmt.Printf("failed to execute root command: %s\n", err)
		return exitError
	}
	return exitOK
}
