package main

import (
	"context"
	"net/http"
	"os"

	"github.com/minguu42/simoom/cli/cmd/root"
	"github.com/minguu42/simoom/cli/cmdutil"
	"github.com/minguu42/simoom/lib/go/simoompb/v1/simoompbconnect"
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
	client := simoompbconnect.NewSimoomServiceClient(http.DefaultClient, "http://localhost:8080")
	c := cmdutil.Core{
		Client:      client,
		Credentials: credentials,
	}

	rootCmd := root.NewCmdRoot(c)
	if err := rootCmd.ExecuteContext(context.Background()); err != nil {
		return exitError
	}
	return exitOK
}
