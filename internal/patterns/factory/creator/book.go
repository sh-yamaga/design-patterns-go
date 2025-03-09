package creator

import (
	"fmt"

	"github.com/sh-yamaga/design-patterns-go/internal/patterns/factory/category"
	"github.com/sh-yamaga/design-patterns-go/internal/patterns/factory/work"
)

type BookCreator struct {
	BaseCreator
}

func (bc BookCreator) createWork(title, creator string) *work.Work {
	return &work.Work{
		Title:    title,
		Creater:  creator,
		Category: category.Book,
	}
}

func (bc BookCreator) registerWork(w *work.Work) work.IWork {
	fmt.Printf("registered book: %s\n", w.Title)
	var book work.Book = work.Book{
		Work: *w,
	}

	return book
}
