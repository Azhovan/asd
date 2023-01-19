package cmd

import (
	"io"

	"github.com/spf13/cobra"
)

func addCmd(output io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "add title, synosis, headlines, dialogs",
	}

	cmd.AddCommand(titleCmd(output))

	return cmd
}
