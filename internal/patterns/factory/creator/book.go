package creator

import (
	"fmt"

	"github.com/yamaga-shu/design-patterns-go/internal/patterns/factory/work"
)

type BookCreator struct {
}

func (bc BookCreator) New(title, creator string) *work.Book {
	b := bc.createBook(title, creator)
	bc.registerBook(b)

	return b
}

func (bc BookCreator) createBook(title, creator string) *work.Book {
	return &work.Book{
		Title:   title,
		Creator: creator,
	}
}

func (bc BookCreator) registerBook(b *work.Book) *work.Book {
	fmt.Printf("registered book: %s\n", b.Title)

	return b
}
