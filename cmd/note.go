package cmd

import (
	"io"
	"os"

	"github.com/spf13/cobra"

	"github.com/azhovan/asd/pkg/gui/terminal"
	notepkg "github.com/azhovan/asd/pkg/note"
)

func noteCmd(output io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "note",
		Short: "prints current note in the output",
		RunE: func(cmd *cobra.Command, args []string) error {
			note, err := notepkg.GetCurrentNote()
			if err != nil {
				return err
			}

			// by default, we use terminal
			// TODO: this should be configurable
			f, err := os.Create("./note.md")
			if err != nil {
				return err
			}
			return terminal.Render(*note, f)
		},
	}

	return cmd
}
