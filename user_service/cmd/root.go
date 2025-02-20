package cmd

import "github.com/spf13/cobra"

func Root() *cobra.Command {
	cmdRoot := &cobra.Command{}
	commands := []func() *cobra.Command{server}
	for _, item := range commands {
		cmdRoot.AddCommand(item())
	}
	return cmdRoot
}
