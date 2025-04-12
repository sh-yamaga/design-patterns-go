package factory

import (
	"fmt"

	"github.com/yamaga-shu/design-patterns-go/internal/patterns/factory/work"
)

type BookFactory struct {
}

func (bf BookFactory) New(title, creator string) *work.Book {
	b := bf.createBook(title, creator)
	bf.registerBook(b)

	return b
}

func (bf BookFactory) createBook(title, creator string) *work.Book {
	return &work.Book{
		Title:   title,
		Creator: creator,
	}
}

func (bf BookFactory) registerBook(b *work.Book) *work.Book {
	fmt.Printf("registered book: %s\n", b.Title)

	return b
}
