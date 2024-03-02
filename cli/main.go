package main

import (
	"context"
	"fmt"
	"os"

	"github.com/minguu42/simoom/cli/api"
	"github.com/minguu42/simoom/cli/cmd/root"
	"github.com/minguu42/simoom/cli/cmdutil"
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
	credentials, err := cmdutil.ReadCredentials()
	if err != nil {
		return exitError
	}
	c, err := api.NewClient()
	if err != nil {
		fmt.Printf("failed to create api client: %s\n", err)
		return exitError
	}
	f := cmdutil.Factory{
		Client:      c,
		Credentials: credentials,
	}

	rootCmd := root.NewCmdRoot(f)
	if err := rootCmd.ExecuteContext(context.Background()); err != nil {
		return exitError
	}
	return exitOK
}
