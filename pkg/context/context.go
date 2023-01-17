package context

type Context struct {
	Title        string
	Descriptions []string
	Headlines    []Headline
}

// Headline represents a single goal tha breaks down into actions items.
type Headline struct {
	Title       string
	Done        bool
	ActionItems []ActionItem
}

// ActionItem represents a single action that is part of a headline.
type ActionItem struct {
	Title string
	Done  bool
}
