package editor

// TextEditor
type Editor struct {
	text           string
	cursorPosition int
}

func NewEditor() *Editor {
	return &Editor{
		text:           "",
		cursorPosition: 0,
	}
}

// Insert inserts string at given cursorPosition
func (e *Editor) Insert(cursorPosition int, text string) {
	if cursorPosition < 0 {
		cursorPosition = 0
	} else if cursorPosition > len(e.text) {
		cursorPosition = len(e.text)
	}

	e.text = e.text[:cursorPosition] + text + e.text[cursorPosition:]
}

// Delete deletes given length string from cursorPosition
func (e *Editor) Delete(cursorPosition int, length int) {
	// Unexpected arguments
	if length < 0 {
		return
	}
	if cursorPosition < 0 {
		return
	}
	if cursorPosition >= len(e.text) {
		return
	}

	endPosition := min(cursorPosition+length, len(e.text))

	e.text = e.text[:cursorPosition] + e.text[endPosition:]
}

func (e *Editor) Text() string {
	return e.text
}

func (e *Editor) SetText(initialText string) {
	e.text = initialText
}
