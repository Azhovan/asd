package note

import (
	"os"

	"gopkg.in/yaml.v3"
	"path/filepath"
)

type ContextFile struct {
	// A key/value of all existing contexts.
	// key represent Note identifier, and value is
	// the path to where Note information is saved.
	Contexts map[string]string `yaml:"contexts"`

	// Current active context.
	CurrentContext string `yaml:"current-context"`
}

// GetContextFile GteContextFile returns context file as a struct within the
// given directory. Usually this directory is asd home directory.
func GetContextFile(path string) (ContextFile, error) {
	path = filepath.Join(path, DefaultContextsFilename)

	ctxFile, err := os.Open(path)
	if err != nil {
		return ContextFile{}, err
	}

	ctx := NewEmptyContexts()
	err = yaml.NewDecoder(ctxFile).Decode(&ctx)

	return ctx, err
}

// ensureContextFile creates a contexts file if it doesn't exist already.
func ensureContextFile(noteDIR string) (*os.File, error) {
	// i.e /tmp/asd/contexts.yaml
	ctxPath := filepath.Join(noteDIR, "..", "contexts.yaml")
	return os.OpenFile(ctxPath, os.O_RDWR|os.O_CREATE, os.ModePerm)
}

// NewEmptyContexts returns an empty context struct.
func NewEmptyContexts() ContextFile {
	return ContextFile{Contexts: make(map[string]string)}
}
