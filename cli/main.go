package main

import (
	"context"
	"errors"
	"os"

	"github.com/minguu42/simoom/cli/cmd"
	"github.com/minguu42/simoom/cli/cmdutil"
	"github.com/spf13/cobra"
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

type globalFlags struct {
	profile string
}

func mainRun() exitCode {
	var flags globalFlags
	var f *cmdutil.Factory
	rootCmd := &cobra.Command{
		Use:   "simoom <command> <subcommand> [flags]",
		Short: "Simoom CLI",
		Long:  `Work seamlessly with Simoom from the command line.`,
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			if cmdutil.IsAuthCheckEnabled(cmd) && !f.Client.CheckCredentials() {
				return errors.New("authentication failed")
			}
			return nil
		},
	}
	rootCmd.PersistentFlags().StringVar(&flags.profile, "profile", "default", "user profile")

	var err error
	if f, err = cmdutil.NewFactory(flags.profile); err != nil {
		return exitError
	}

	rootCmd.AddCommand(cmd.NewCmdAuth(f))
	rootCmd.AddCommand(cmd.NewCmdProject(f))
	rootCmd.AddCommand(cmd.NewCmdStep(f))
	rootCmd.AddCommand(cmd.NewCmdTag(f))
	rootCmd.AddCommand(cmd.NewCmdTask(f))

	if err := rootCmd.ExecuteContext(context.Background()); err != nil {
		return exitError
	}
	return exitOK
}
