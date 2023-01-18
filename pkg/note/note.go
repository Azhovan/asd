package note

import (
	"fmt"
	"os"
	"path/filepath"
)

// NewNote creates a new note(context), adds it to context list
// and set it to the current context.
func NewNote(name string) (string, error) {
	asdDIR, err := GetDefaultASDDirectory()
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

// GetDefaultASDDirectory returns the default directory where all notes exist.
func GetDefaultASDDirectory() (string, error) {
	userhome, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/.asd", userhome), nil
}
