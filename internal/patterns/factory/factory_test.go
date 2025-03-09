package factory

import (
	"testing"

	"github.com/sh-yamaga/design-patterns-go/internal/patterns/factory/category"
	"github.com/sh-yamaga/design-patterns-go/internal/patterns/factory/creator"
)

func TestWorkFactory(t *testing.T) {
	var rwf CreatorFactory = CreatorFactory{}
	bookCreator, _ := rwf.NewCreator(category.Book)
	movieCreator, _ := rwf.NewCreator(category.Movie)

	// Check if bookCreator is of the expected type
	if _, ok := interface{}(bookCreator).(creator.IWorkCreator); !ok {
		t.Errorf("Expected bookCreator to be of type WorkFactory, got %T", bookCreator)
	}

	// Check if movieCreator is of the expected type
	if _, ok := interface{}(movieCreator).(creator.IWorkCreator); !ok {
		t.Errorf("Expected movieCreator to be of type WorkFactory, got %T", movieCreator)
	}
}
