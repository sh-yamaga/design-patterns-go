package memento

// EditorMemento holds a snapshot of the Editor's state
type EditorMemento struct {
	text           string
	cursorPosition int
}

// Editor holds the current text and cursorPosition
// and is responsible for creating and restoring EditorMemento
type Editor struct {
	Text           string
	CursorPosition int
}

func NewEditor() *Editor {
	return &Editor{}
}

// TypeWords add given words
func (e *Editor) TypeWords(words string) {
	e.Text += words
	e.CursorPosition = len(e.Text)
}

// Save returns EditorMemento which holds current Editor's state
func (e *Editor) Save() *EditorMemento {
	return &EditorMemento{
		text:           e.Text,
		cursorPosition: e.CursorPosition,
	}
}

// Restore restores Editor's state from Memento
func (e *Editor) Restore(em *EditorMemento) {
	e.Text = em.text
	e.CursorPosition = em.cursorPosition
}
