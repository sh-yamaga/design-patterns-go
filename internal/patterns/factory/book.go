package factory

import (
	"fmt"

	"github.com/sh-yamaga/design-patterns-go/internal/patterns/factory/work"
)

type BookFactory struct {
	WorkFactory
}

func (bf BookFactory) Create(title, creater string) *work.Work {
	var book *work.Work = bf.createWork(title, creater)
	bf.registerWork(book)

	return book
}

func (bf BookFactory) createWork(title, creater string) *work.Work {
	return &work.Work{
		Title:    title,
		Creater:  creater,
		Category: work.BookCategory,
	}
}

func (bf BookFactory) registerWork(w *work.Work) {
	fmt.Println("registered a Book: " + w.Title)
}
