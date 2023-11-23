package main

import (
	"github.com/minguu42/simoom/cli/cmd/auth"
	"github.com/minguu42/simoom/cli/cmd/project"
	"github.com/minguu42/simoom/cli/cmd/step"
	"github.com/minguu42/simoom/cli/cmd/tag"
	"github.com/minguu42/simoom/cli/cmdutil"
	"github.com/spf13/cobra"
)

func newCmdRoot(core cmdutil.Core) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "simoom",
		Short: "Simoom CLI",
		Long:  `Work seamlessly with Simoom from the command line.`,
	}
	cmd.AddCommand(auth.NewCmdAuth(core))
	cmd.AddCommand(project.NewCmdProject(core))
	cmd.AddCommand(step.NewCmdStep(core))
	cmd.AddCommand(tag.NewCmdTag(core))
	cmd.AddCommand(newCmdTask(core))
	return cmd
}
