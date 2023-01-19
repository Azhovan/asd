package note

const (
	// DefaultContextsFilename represents the file that Note configs are stored.
	DefaultContextsFilename = "contexts.yaml"

	// DefaultNoteFilename represents the file name to store Note data into it.
	// i.e
	// ├── Category1
	// └── note.yaml
	// ├── Category2
	// | └── note.yaml
	// └── contexts.yaml
	DefaultNoteFilename = "note.yaml"
)
