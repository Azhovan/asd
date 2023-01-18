package cmd

import (
	"fmt"
	"github.com/azhovan/asd/pkg/note"
	"io"
	"os"

	"github.com/spf13/cobra"
)

func contextCmd(output io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "context",
		Short: "prints current context in output",
		RunE: func(cmd *cobra.Command, args []string) error {
			asdDIR, err := note.GetDefaultASDDirectory()
			if err != nil {
				return err
			}

			ctx, err := note.GetContextFile(asdDIR)
			if err != nil && !os.IsNotExist(err) {
				return err
			}

			if ctx.CurrentContext == "" {
				fmt.Fprintln(output, "There is no defined context")
				return nil
			}

			fmt.Fprintf(output, "Current context: %s\n", ctx.CurrentContext)
			return nil
		},
	}

	return cmd
}
