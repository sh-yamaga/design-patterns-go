package creator

import (
	"fmt"

	"github.com/sh-yamaga/design-patterns-go/internal/patterns/factory/category"
	"github.com/sh-yamaga/design-patterns-go/internal/patterns/factory/work"
)

type BookCreator struct {
	BaseCreator
}

func (bf BookCreator) createWork(title, creater string) *work.Work {
	return &work.Work{
		Title:    title,
		Creater:  creater,
		Category: category.Book,
	}
}

func (bf BookCreator) registerWork(w *work.Work) {
	fmt.Println("registered a Book: " + w.Title)
}
