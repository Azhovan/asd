package cmd

import (
	"io"
	"os"

	"github.com/spf13/cobra"

	"github.com/azhovan/asd/pkg/gui/terminal"
	notepkg "github.com/azhovan/asd/pkg/note"
)

type NoteOptions struct {
	NoteID string
}

func getCmd(output io.Writer) *cobra.Command {
	opts := &NoteOptions{}

	cmd := &cobra.Command{
		Use:   "note",
		Short: "prints current note in the output",
		RunE: func(cmd *cobra.Command, args []string) error {
			note, err := notepkg.GetNote(opts.NoteID)
			if err != nil {
				return err
			}

			// export file in current directory
			f, err := os.Create("./note.md")
			if err != nil {
				return err
			}

			return terminal.Render(*note, f)
		},
	}

	cmd.Flags().StringVarP(&opts.NoteID, "id", "i", "", "note identifier")

	return cmd
}
