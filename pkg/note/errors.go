package note

import "errors"

var (
	errContextNameIsMissing = errors.New("field name is missing")
	errContextIsMissing     = errors.New("context DIR is missing")

	// ErrNoteIsNotExist is returned when we want to access current note
	// and there is no note.
	ErrNoteIsNotExist = errors.New("there is no note")
)
