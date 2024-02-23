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
	client := simoompbconnect.NewSimoomServiceClient(http.DefaultClient, "http://localhost:8080")
	c := cmdutil.Core{
		Client: client,
		Credentials: cmdutil.Credentials{
			AccessToken:  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDI0Njg5NjIsIm5hbWUiOiJtaW5ndXU0MiIsImlkIjoiMDFIR1pHNEdQTjQxRDFGTU5CMlNWR0RDQkUifQ.yshdci246JdXrKkJ_gtvLD48WGuE8aHjLXcFRh5FYf0",
			RefreshToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDI0Njg5NjIsImlkIjoiMDFIR1pHNEdQTjQxRDFGTU5CMlNWR0RDQkUifQ.7A6g3F57Tji-aycdrmWli9_cbXkNF3Q-fAMgAu7g0Yw",
		},
	}

	rootCmd := root.NewCmdRoot(c)
	if err := rootCmd.ExecuteContext(context.Background()); err != nil {
		return exitError
	}
	return exitOK
}
