package support

import cor "github.com/yamaga-shu/design-patterns-go/internal/patterns/CoR"

// Entry represents entry of Chain of Responsibility
type Entry struct {
	HttpResponseSupport
}

func NewEntry() *Entry {
	return &Entry{
		HttpResponseSupport{
			name:        "Entry",
			supportFrom: 0,
			supportTo:   0,
			next:        nil,
			handler:     entryHandler{},
		},
	}
}

type entryHandler struct{}

func (eh entryHandler) handleResponse(hr *cor.HttpResponse) {
	// No response would not passed
}
