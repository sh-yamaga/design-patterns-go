package memento

// History holds a list of EditorMemento
// and is responsible for managing and removing them
type History struct {
	mementos []*EditorMemento
}

// Push appends new EditorMemento
func (h *History) Push(m *EditorMemento) {
	h.mementos = append(h.mementos, m)
}

// Pop removes and returns the last mementos item
func (h *History) Pop() *EditorMemento {
	if len(h.mementos) == 0 {
		return nil
	}

	i := len(h.mementos) - 1
	m := h.mementos[i]
	h.mementos = h.mementos[:i]

	return m
}
