package terminal

import (
	"github.com/azhovan/asd/pkg/note"
	"html/template"
	"io"
)

// Render visualize a Note in terminal.
func Render(note note.Note, out io.Writer) error {
	tpl := template.Must(template.ParseFiles("note.md.tpl"))
	err := tpl.Execute(out, note)
	if err != nil {
		return err
	}

	return nil
}
