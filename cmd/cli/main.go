package main

import (
	"context"
	"net/http"
	"os"

	"github.com/minguu42/simoom/cmd/cli/cmdutil"
	"github.com/minguu42/simoom/gen/simoompb/v1/simoompbconnect"
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
			AccessToken:  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDA0ODAwMjIsIm5hbWUiOiJtaW5ndXU0MiIsImlkIjoiMDFIRjQ3QVg2UjE3NUUzMzdONDlDVFpENzcifQ.5rnjv-HA3tGLZxn7zD9KIv1f6mcNp1iQ4d-FuivJrdg",
			RefreshToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDA0ODAwMjIsImlkIjoiMDFIRjQ3QVg2UjE3NUUzMzdONDlDVFpENzcifQ.gNk8Vh4Xon1ponEYOO2N2o4Yv1RUTMU5TBKvloHAt3A",
		},
	}

	rootCmd := newCmdRoot(c)
	if err := rootCmd.ExecuteContext(context.Background()); err != nil {
		return exitError
	}
	return exitOK
}
