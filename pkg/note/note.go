package note

import (
	"io"
	"os"

	"gopkg.in/yaml.v3"
	"path/filepath"

	"github.com/azhovan/asd/pkg/config"
)

type Note struct {
	// ID is the Note identifier
	ID string

	// Title of the note
	Title string

	// Synopsis represents the summary of a note.
	Synopsis []string

	// Headlines are a list of goals or items that are the focus area of
	// the note.
	Headlines []Headline

	// Dialogue is a conversation involving any number of people.
	Dialogs []Dialog
}

// Headline represents a single goal tha breaks down into actions items.
type Headline struct {
	Title string

	// when true, headline marked as completed.
	Done bool

	// list of actions items
	ActionItems []ActionItem
}

// ActionItem represents a single action that is part of a headline.
type ActionItem struct {
	Title string

	// when true, action item marked as completed.
	Done bool
}

// A Dialog is a single conversation owned by one person in a chat.
type Dialog struct {
	Person    string
	Message   string
	Timestamp string
}

// NewNote creates a new note(context), adds it to context list
// and set it to the current context.
func NewNote(noteID string) error {
	if noteID == "" {
		return ErrNoteIDIsMissing
	}

	asdDIR, err := config.GetDefaultASDDirectory()
	if err != nil {
		return err
	}

	noteDIR := filepath.Join(asdDIR, noteID)
	err = createNoteStructure(noteID, noteDIR)

	return err
}

// GetNote loads Note object for the given Note name.
// when name is empty, it returns the last active Note.
// A Note is active if it has just been edited or created.
// Reading the contexts of a Note won't make it recent active Note.
func GetNote(noteID string) (*Note, error) {
	asdDIR, err := config.GetDefaultASDDirectory()
	if err != nil {
		return nil, err
	}

	ctxFile, err := GetContextFile(asdDIR)
	if err != nil && os.IsNotExist(err) {
		return nil, ErrContextIsMissing
	}
	if err != nil {
		return nil, err
	}

	// when no name is provided, recent active Note is loaded
	var noteDIR string
	if noteID == "" {
		noteDIR = filepath.Join(asdDIR, ctxFile.CurrentContext, DefaultNoteFilename)
	} else {
		noteDIR = filepath.Join(ctxFile.Contexts[noteID], DefaultNoteFilename)
	}

	noteFile, err := os.Open(noteDIR)
	if err != nil && os.IsNotExist(err) {
		return nil, ErrNoteIsMissing
	}
	if err != nil {
		return nil, err
	}

	note := newEmptyNote(noteID)
	err = yaml.NewDecoder(noteFile).Decode(note)
	// Note can be empty
	if err != io.EOF {
		return nil, err
	}

	return note, nil
}

// createNoteStructure sets the current context to the new context and also add
// the new context path to the list of the existing context's paths.
func createNoteStructure(NoteID, noteDIR string) error {
	if err := ensureNoteDIR(noteDIR); err != nil {
		return err
	}

	return createNote(NoteID, noteDIR)
}

// createNote create an empty Note file and set the current context
// to the Note identifier.
func createNote(NoteID, noteDIR string) error {
	ctxFile, err := ensureContextFile(noteDIR)
	if err != nil {
		return err
	}

	// we populate an empty context struct with
	// existing context file, then set add the new context
	// to this struct and finally write it over existing one.
	newContexts := NewEmptyContexts()

	err = yaml.NewDecoder(ctxFile).Decode(&newContexts)
	if err != nil && err != io.EOF {
		return err
	}

	newContexts.CurrentContext = NoteID
	newContexts.Contexts[NoteID] = noteDIR

	ctxFile, err = os.Create(ctxFile.Name())
	if err != nil {
		return err
	}

	// overwrite existing context file
	err = yaml.NewEncoder(ctxFile).Encode(&newContexts)
	if err != nil {
		return err
	}

	// create an empty note.yaml file
	noteFile := filepath.Join(noteDIR, DefaultNoteFilename)
	_, err = os.Create(noteFile)

	return err
}

// ensureNoteDir creates an empty directory if doesn't exist already.
func ensureNoteDIR(noteDIR string) error {
	err := os.MkdirAll(noteDIR, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// newEmptyNote returns an empty Note struct.
func newEmptyNote(noteID string) *Note {
	return &Note{
		ID:        noteID,
		Title:     "",
		Synopsis:  make([]string, 0),
		Headlines: make([]Headline, 0),
		Dialogs:   make([]Dialog, 0),
	}
}
