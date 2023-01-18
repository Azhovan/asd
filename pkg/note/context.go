package note

import (
	"errors"
	"io"
	"os"

	"gopkg.in/yaml.v3"
	"path/filepath"
)

type ContextFile struct {
	// A list of all existing contexts.
	Contexts map[string]string `yaml:"contexts"`

	// Current active context.
	CurrentContext string `yaml:"current-context"`
}

var (
	errContextNameIsMissing = errors.New("field name is missing")
	errContextIsMissing     = errors.New("context DIR is missing")
)

// SetNewContext sets the current context to the new context and also add
// the new context path to the list of the existing context's paths.
func SetNewContext(name, path string) error {
	ctxFile, err := ensureContextFile(path)
	if err != nil {
		return err
	}

	err = ensureNewContextDir(path)
	if err != nil {
		return err
	}

	if name == "" {
		return errContextNameIsMissing
	}

	return setContext(ctxFile, name, path)
}

// GetContextFile GteContextFile returns context file as a struct.
func GetContextFile(path string) (ContextFile, error) {
	path = filepath.Join(path, "contexts.yaml")

	ctxFile, err := os.Open(path)
	if err != nil {
		return ContextFile{}, err
	}

	ctx := NewEmptyContexts()
	err = yaml.NewDecoder(ctxFile).Decode(&ctx)

	return ctx, err
}

func setContext(ctxFile *os.File, name, path string) error {
	contexts := NewEmptyContexts()

	err := yaml.NewDecoder(ctxFile).Decode(&contexts)
	if err != nil && err != io.EOF {
		return err
	}
	contexts.CurrentContext = name
	contexts.Contexts[name] = path

	ctxFile, err = os.Create(ctxFile.Name())
	if err != nil {
		return err
	}
	return yaml.NewEncoder(ctxFile).Encode(&contexts)
}

func ensureNewContextDir(path string) error {
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return errContextIsMissing
	}

	return nil
}

func ensureContextFile(path string) (*os.File, error) {
	// i.e /tmp/asd/contexts.yaml
	ctxPath := filepath.Join(path, "..", "contexts.yaml")
	return os.OpenFile(ctxPath, os.O_RDWR|os.O_CREATE, os.ModePerm)
}

// NewEmptyContexts returns an empty context struct.
func NewEmptyContexts() ContextFile {
	return ContextFile{Contexts: make(map[string]string)}
}
