package terminal

import (
	"os"
	"testing"
	"time"

	notepkg "github.com/azhovan/asd/pkg/note"
)

func TestRender(t *testing.T) {
	note := notepkg.Note{
		Title:    "Note Title",
		Synopsis: []string{"this is great", "imagine how powerful this is!"},
		Headlines: []notepkg.Headline{
			notepkg.Headline{
				Title: "Drafting a KEP",
				Done:  false,
				ActionItems: []notepkg.ActionItem{
					{
						Title: "github permissions",
						Done:  false,
					},
					{
						Title: "talk to product managers",
						Done:  true,
					},
				},
			},
			notepkg.Headline{
				Title: "Merging the KEP",
				Done:  false,
				ActionItems: []notepkg.ActionItem{
					{
						Title: "get buy-in",
						Done:  false,
					},
					{
						Title: "prepare a good presentation",
						Done:  true,
					},
				},
			},
		},
		Dialogs: []notepkg.Dialog{
			{
				Person:    "Alex",
				Message:   "this is great, lets do it",
				Timestamp: time.Now().Format("2006-01-02 15:04"),
			}, {
				Person:    "Ali",
				Message:   "sign me up!",
				Timestamp: time.Now().Format("2006-01-02 15:04"),
			},
		},
	}

	f, _ := os.Create("./note.md")
	err := Render(note, f)
	if err != nil {
		t.Fatalf("err expected to be nil, got: %v", err)
	}

}
