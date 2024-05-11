package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/minguu42/simoom/cli/cmd"
	"github.com/minguu42/simoom/cli/cmdutil"
	"github.com/minguu42/simoom/cli/factory"
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
	rootCmd := &cobra.Command{
		Use:   "simoom <command> <subcommand> [flags]",
		Short: "Simoom CLI",
		Long:  `Work seamlessly with Simoom from the command line.`,
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			f, err := factory.New(flags.profile)
			if err != nil {
				return fmt.Errorf("failed to create factory: %w", err)
			}
			if cmdutil.IsAuthCheckEnabled(cmd) && !f.Client.CheckCredentials() {
				return errors.New("authentication failed")
			}

			cmd.SetContext(factory.ContextWithFactory(cmd.Context(), f))
			return nil
		},
		SilenceUsage: true,
	}
	rootCmd.PersistentFlags().StringVar(&flags.profile, "profile", "default", "user profile")

	rootCmd.AddCommand(
		cmd.NewCmdAuth(),
		cmd.NewCmdProject(),
		cmd.NewCmdStep(),
		cmd.NewCmdTag(),
		cmd.NewCmdTask(),
	)

	if err := rootCmd.ExecuteContext(context.Background()); err != nil {
		return exitError
	}
	return exitOK
}
