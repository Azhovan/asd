package note

import (
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"testing"
)

func TestSetNewContext_WithNoExistingContexts(t *testing.T) {
	newContextDir, err := os.MkdirTemp("", "new-context-dir-*")
	if err != nil {
		t.Fatalf("can't create temporary contexts-file, %v", err)
	}

	asdhome := filepath.Join(newContextDir, "..")
	err = os.MkdirAll(asdhome, os.ModePerm)
	if err != nil {
		t.Fatalf("can't create asd directry")
	}
	defer t.Cleanup(func() {
		_ = os.RemoveAll(asdhome)
	})

	err = SetNewContext("new-contex", newContextDir)
	if err != nil {
		t.Fatalf("can't set new context: %v", err)
	}
}

func TestNewContext_WithExistingContext(t *testing.T) {
	asdhome, err := os.MkdirTemp("", "asd-*")
	if err != nil {
		t.Fatalf("can't create asd temp home directory: %v", err)
	}
	// create a context file
	ctxs, err := os.Create(filepath.Join(asdhome, "contexts.yaml"))
	if err != nil {
		t.Fatalf("can't create contexts.yaml file: %v", err)
	}

	t.Cleanup(func() {
		_ = os.RemoveAll(asdhome)
	})

	existingContexts := NewEmptyContexts()
	existingContexts.Contexts = map[string]string{
		"ctx1": "/path/to/ctx1",
		"ctx2": "/path/to/ctx2",
		"ctx3": "/path/to/ctx3",
	}
	existingContexts.CurrentContext = "ctx3"
	err = yaml.NewEncoder(ctxs).Encode(existingContexts)
	if err != nil {
		t.Fatalf("can't update contexts.yaml file content")
	}

	// new context
	newContextPath := filepath.Join(asdhome, "new-context")
	err = os.Mkdir(newContextPath, os.ModePerm)
	if err != nil {
		t.Fatalf("can't create new context directory: %v", err)
	}

	err = SetNewContext("new-context", newContextPath)
	if err != nil {
		t.Fatalf("Can't set context: %v", err)
	}
}

func TestSetNewContext_AddMultipleContext(t *testing.T) {
	asdhome, err := os.MkdirTemp("", "asd-*")
	if err != nil {
		t.Fatalf("can't create asd temp home directory: %v", err)
	}
	// create a context file
	_, err = os.Create(filepath.Join(asdhome, "contexts.yaml"))
	if err != nil {
		t.Fatalf("can't create contexts.yaml file: %v", err)
	}

	t.Cleanup(func() {
		_ = os.RemoveAll(asdhome)
	})

	// full context path
	// context name
	testdata := []struct {
		fullContextPath string
		contextName     string
	}{
		{
			fullContextPath: filepath.Join(asdhome, "ctx1"),
			contextName:     "ctx1",
		},
		{
			fullContextPath: filepath.Join(asdhome, "ctx2"),
			contextName:     "ctx2",
		},
		{
			fullContextPath: filepath.Join(asdhome, "ctx3"),
			contextName:     "ctx3",
		},
		{
			fullContextPath: filepath.Join(asdhome, "ctx4"),
			contextName:     "ctx4",
		},
	}
	// create all context
	for _, v := range testdata {
		err := os.Mkdir(v.fullContextPath, os.ModePerm)
		if err != nil {
			t.Fatalf("can't create DIR for new context: %v", err)
		}

		err = SetNewContext(v.contextName, v.fullContextPath)
		if err != nil {
			t.Fatalf("can't set new context: %v, err: %v", v, err)
		}
	}

	ctxFileStruct, err := GetContextFile(asdhome)
	if err != nil {
		t.Fatalf("can't find context file in the given path: %v", err)
	}
	if curr := ctxFileStruct.CurrentContext; curr != "ctx4" {
		t.Fatalf("expected current context to be ctx4, got %s", curr)
	}
	if l := len(ctxFileStruct.Contexts); l != 4 {
		t.Fatalf("expected context-file to has length of 4, got: %d", l)
	}

	for _, v := range testdata {
		_, ok := ctxFileStruct.Contexts[v.contextName]
		if !ok {
			t.Fatalf("expected value of %s be present in the contexts file, got nil", v.contextName)
		}
	}

}
