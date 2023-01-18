package note

import (
	"os"

	"gopkg.in/yaml.v3"
	"path/filepath"

	"github.com/azhovan/asd/pkg/config"
)

// NewNote creates a new note(context), adds it to context list
// and set it to the current context.
func NewNote(name string) (string, error) {
	asdDIR, err := config.GetDefaultASDDirectory()
	if err != nil {
		return "", err
	}

	noteDIR := filepath.Join(asdDIR, name)
	err = os.MkdirAll(noteDIR, os.ModePerm)
	if err != nil {
		return "", err
	}

	err = SetNewContext(name, noteDIR)
	return noteDIR, err
}

// GetCurrentNote returns the last active Note.
// A Note is active if it has just been edited or created.
// Reading the contexts of a Note won't make it current Note.
func GetCurrentNote() (*Note, error) {
	asdDIR, err := config.GetDefaultASDDirectory()
	if err != nil {
		return nil, err
	}

	ctxFile, err := GetContextFile(asdDIR)
	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}
	if err != nil && os.IsNotExist(err) {
		return nil, ErrNoteIsNotExist
	}

	noteDIR := filepath.Join(asdDIR, ctxFile.CurrentContext, "note.yaml")

	noteFile, err := os.Open(noteDIR)
	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}
	if err != nil && os.IsNotExist(err) {
		return nil, ErrNoteIsNotExist
	}

	note := newEmptyNote()
	err = yaml.NewDecoder(noteFile).Decode(note)

	return note, err
}

// newEmptyNote returns an empty Note struct.
func newEmptyNote() *Note {
	return &Note{
		Title:     "",
		Synopsis:  make([]string, 0),
		Headlines: make([]Headline, 0),
		Dialogs:   make([]Dialog, 0),
	}
}
