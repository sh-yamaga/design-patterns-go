package factory

import (
	"testing"

	"github.com/sh-yamaga/design-patterns-go/internal/patterns/factory/book"
	"github.com/sh-yamaga/design-patterns-go/internal/patterns/factory/movie"
)

func TestWorkFactory(t *testing.T) {
	var wf WorkFactory = WorkFactory{}
	var bf book.BookFactory = wf.BookFactory()
	var mf movie.MovieFactory = wf.MovieFactory()

	// Check if bf is of the expected type
	if _, ok := interface{}(bf).(book.BookFactory); !ok {
		t.Errorf("Expected bf to be of type book.BookFactory, got %T", bf)
	}

	// Check if mf is of the expected type
	if _, ok := interface{}(mf).(movie.MovieFactory); !ok {
		t.Errorf("Expected mf to be of type movie.MovieFactory, got %T", mf)
	}
}
