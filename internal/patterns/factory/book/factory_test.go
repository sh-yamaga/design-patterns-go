package book_test

import (
	"fmt"
	"testing"

	"github.com/sh-yamaga/design-patterns-go/internal/patterns/factory"
	"github.com/sh-yamaga/design-patterns-go/internal/patterns/factory/book"
)

func TestNewBookFactory(t *testing.T) {
	var bf book.BookFactory = book.NewBookFactory()

	book := bf.BaseFactory.Create("Kurt Vonnegut", "The Sirens of Titan", factory.WorkCategory(book.SF))

	got := book.String()
	want := fmt.Sprintf("Title: %s, Author: %s, Category: %s", "Kurt Vonnegut", "The Sirens of Titan", "SF")

	if got != want {
		t.Errorf("got: %s, want: %s", got, want)
	}
}
