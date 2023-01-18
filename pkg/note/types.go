package note

type Note struct {
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
