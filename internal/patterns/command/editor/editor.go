package editor

// TextEditor
type Editor struct {
	text        string
	cursorIndex int
}

func NewEditor() *Editor {
	return &Editor{
		text:        "",
		cursorIndex: 0,
	}
}

// Insert inserts string at given cursorIndex
func (e *Editor) Insert(cursorIndex int, text string) {
	if cursorIndex < 0 {
		cursorIndex = 0
	} else if cursorIndex > len(e.text) {
		cursorIndex = len(e.text)
	}

	e.text = e.text[:cursorIndex] + text + e.text[cursorIndex:]
}

// Delete deletes given length string from cursorIndex
func (e *Editor) Delete(cursorIndex int, length int) string {
	// Unexpected arguments
	if length < 0 {
		return ""
	}
	if cursorIndex < 0 {
		return ""
	}
	if cursorIndex > len(e.text) {
		return ""
	}

	endIndex := min(cursorIndex+length, len(e.text))

	deleted := e.text[cursorIndex:endIndex]

	e.text = e.text[:cursorIndex] + e.text[endIndex:]

	return deleted
}

func (e *Editor) Text() string {
	return e.text
}

func (e *Editor) SetText(text string) {
	e.text = text
}
