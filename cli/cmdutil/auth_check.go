package cmdutil

import "github.com/spf13/cobra"

func DisableAuthCheck(cmd *cobra.Command) {
	if cmd.Annotations == nil {
		cmd.Annotations = map[string]string{}
	}
	cmd.Annotations["skipAuthCheck"] = "true"
}

func IsAuthCheckEnabled(cmd *cobra.Command) bool {
	switch cmd.Name() {
	case "help", cobra.ShellCompRequestCmd, cobra.ShellCompNoDescRequestCmd:
		return false
	}
	if cmd.Annotations != nil && cmd.Annotations["skipAuthCheck"] == "true" {
		return false
	}

	return true
}
