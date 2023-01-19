package note

import (
	"fmt"
	"github.com/azhovan/asd/pkg/config"
	"os"
	"testing"
)

func Test_NewNote(t *testing.T) {
	asdDIR, err := config.GetDefaultASDDirectory()
	if err != nil {
		t.Fatalf("error expected to be nil, got: %v", err)
	}
	t.Cleanup(func() {
		_ = os.RemoveAll(asdDIR)
	})

	err = NewNote("note-123")
	if err != nil {
		t.Fatalf("error expected to be nil, got: %v", err)
	}

	n, err := GetNote("note-123")
	if err != nil {
		t.Fatalf("error expected to be nil, got: %v", err)
	}

	if n.ID != "note-123" {
		t.Fatalf("expted note identifier to be note-123, got :%s", n.ID)
	}
}

func Test_MultipleNewNote(t *testing.T) {
	asdDIR, err := config.GetDefaultASDDirectory()
	if err != nil {
		t.Fatalf("error expected to be nil, got: %v", err)
	}
	t.Cleanup(func() {
		_ = os.RemoveAll(asdDIR)
	})

	for i := 0; i < 4; i++ {
		err := NewNote(fmt.Sprintf("note-%d", i))
		if err != nil {
			t.Fatalf("error expected to be nil, got: %v", err)
		}
	}

	ctxFile, err := GetContextFile(asdDIR)
	if err != nil {
		t.Fatalf("error expected to be nil, got: %v", err)
	}
	if len(ctxFile.Contexts) != 4 {
		t.Fatalf("expted 4 notes to be created, got :%d", len(ctxFile.Contexts))
	}
	if ctxFile.CurrentContext == "note-4" {
		t.Fatalf("expted current active notes to be note-4, got :%s", ctxFile.CurrentContext)
	}
}
