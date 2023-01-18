package note

import (
	"fmt"
	"os"
)

// NewNote creates a new context, adds it to context list
// and set the current context to it.
func NewNote(name string) (string, error) {
	userhome, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	noteDIR := fmt.Sprintf("%s/.asd/%s", userhome, name)
	err = os.MkdirAll(noteDIR, os.ModePerm)
	if err != nil {
		return "", err
	}

	err = SetNewContext(name, noteDIR)
	return noteDIR, err
}
