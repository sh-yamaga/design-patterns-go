package command

import "github.com/sh-yamaga/design-patterns-go/internal/patterns/command/editor"

// InsertCommand is a command which insert string to editor
type InsertCommand struct {
	Editor         *editor.Editor
	CursorPosition int
	TextToInsert   string
	prevState      string
}

// Execute excutes insert
func (ic *InsertCommand) Execute() {
	// Save previous state
	ic.prevState = ic.Editor.Text()

	ic.Editor.Insert(ic.CursorPosition, ic.TextToInsert)
}

// Undo reset to previous state
func (ic *InsertCommand) Undo() {
	ic.Editor.SetText(ic.prevState)
}
