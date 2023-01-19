package cmd

import (
	"fmt"
	"io"
	"time"

	"github.com/spf13/cobra"

	"github.com/azhovan/asd/pkg/note"
)

type options struct {
	NoteID string // A Note identifier.
}

// initCmd creates a new context and sets the current context to it.
// When no name is provided a random name is generated for it.
// This function is idempotent, so creating a context with existing names has no impact
// and no error is returned.
func initCmd(out io.Writer) *cobra.Command {
	opts := &options{}

	cmd := &cobra.Command{
		Use:   "init",
		Short: "create a new context",
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.NoteID == "" {
				opts.NoteID = fmt.Sprintf("init-%d", time.Now().Unix())
			}

			if err := note.NewNote(opts.NoteID); err != nil {
				return err
			}
			fmt.Fprint(out, "An empty note has been created.")

			return nil
		},
	}

	cmd.Flags().StringVarP(&opts.NoteID, "id", "i", "", "note identifier")

	return cmd
}
