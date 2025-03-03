package factory

import (
	"testing"
)

func TestWorkFactory(t *testing.T) {
	var rwf RootWorkFactory = RootWorkFactory{}
	var bookFactory *WorkFactory = rwf.Create(Book)
	var movieFactory *WorkFactory = rwf.Create(Movie)

	// Check if bookFactory is of the expected type
	if _, ok := interface{}(bookFactory).(WorkFactory); !ok {
		t.Errorf("Expected bookFactory to be of type WorkFactory, got %T", bookFactory)
	}

	// Check if movieFactory is of the expected type
	if _, ok := interface{}(movieFactory).(WorkFactory); !ok {
		t.Errorf("Expected movieFactory to be of type WorkFactory, got %T", movieFactory)
	}
}
