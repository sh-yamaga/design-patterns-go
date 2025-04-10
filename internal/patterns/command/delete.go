package command

import "github.com/sh-yamaga/design-patterns-go/internal/patterns/command/editor"

// DeleteCommand is a command which delete string from Editor
type DeleteCommand struct {
	Editor      *editor.Editor
	CursorIndex int
	Length      int
	deletedText string
	prevState   string
}

// Execute executes delete
func (dc *DeleteCommand) Execute() {
	// save previous state
	dc.prevState = dc.Editor.Text()

	// delete text and remind it
	dc.deletedText = dc.Editor.Delete(dc.CursorIndex, dc.Length)
}

// Undo reset to previous state
func (dc *DeleteCommand) Undo() {
	dc.Editor.SetText(dc.prevState)
}
