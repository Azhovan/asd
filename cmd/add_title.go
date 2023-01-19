package cmd

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"

	notepkg "github.com/azhovan/asd/pkg/note"
)

type TitleOptions struct {
	NoteID string
}

func titleCmd(output io.Writer) *cobra.Command {
	opts := &TitleOptions{}

	cmd := &cobra.Command{
		Use:   "title [NoteID]",
		Short: "Add Note's title",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 1 {
				fmt.Fprintln(output, "too many arguments")
				return nil
			}
			if len(args) < 1 {
				fmt.Fprintln(output, "argument is missing")
				return nil
			}

			n, err := notepkg.GetNote(opts.NoteID)
			if err != nil {
				return err
			}

			n.Title = args[0]
			return err
			//return notepkg.Update(n)
		},
	}

	cmd.Flags().StringVarP(&opts.NoteID, "id", "i", "", "note identifier")
	return cmd
}
