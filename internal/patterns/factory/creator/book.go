package creator

import (
	"fmt"

	"github.com/sh-yamaga/design-patterns-go/internal/patterns/factory/category"
	"github.com/sh-yamaga/design-patterns-go/internal/patterns/factory/work"
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
		Work: work.Work{
			Title:    title,
			Creater:  creator,
			Category: category.Book,
		}}
}

func (bc BookCreator) registerBook(b *work.Book) *work.Book {
	fmt.Printf("registered book: %s\n", b.Title)

	return b
}
