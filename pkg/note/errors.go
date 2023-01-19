package note

import "errors"

var (
	ErrContextIsMissing = errors.New("context DIR is missing")

	// ErrNoteIsMissing is returned when contexts-file exist but for some reasons there is no Note
	// file. this is very rare scenario though.
	ErrNoteIsMissing = errors.New("note file is missing")

	ErrNoteIDIsMissing = errors.New("note identifier is missing")
)
