package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func newCommand() *cobra.Command {
	cmd := &cobra.Command{
		Short:        "A developer driven TODO app",
		Use:          "asd init --name [NAME]",
		SilenceUsage: true,
	}
	cmd.AddCommand(initCmd(os.Stdout))
	cmd.AddCommand(getCmd(os.Stdout))
	cmd.AddCommand(addCmd(os.Stdout))

	return cmd
}

func Execute() {
	rootCmd := newCommand()
	if err := rootCmd.Execute(); err != nil {
		os.Exit(0)
	}
}
