package book

import (
	"fmt"

	"github.com/sh-yamaga/design-patterns-go/internal/patterns/factory"
)

type BookFactory struct {
	factory.BaseFactory
}

func NewBookFactory() *BookFactory {
	bf := &BookFactory{}
	bf.BaseFactory.IFactory = bf
	return bf
}

func (BookFactory) createWork(creater string, title string, category factory.WorkCategory) *factory.Work {
	bookCategory := BookCategory(category)
	book := &Book{
		Title:    title,
		Author:   creater,
		Category: bookCategory,
	}
	return &factory.Work{
		IWork: book,
	}
}

func (BookFactory) registerWork(book *Book) {
	fmt.Printf("%s was registered.\n", book.Title)
}
