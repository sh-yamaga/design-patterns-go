package factory

import (
	"fmt"

	"github.com/sh-yamaga/design-patterns-go/internal/patterns/factory/work"
)

type BookFactory struct {
	Factory
}

func (bf BookFactory) createWork(title, creater string) *work.Work {
	return &work.Work{
		Title:    title,
		Creater:  creater,
		Category: work.Book,
	}
}

func (bf BookFactory) registerWork(w *work.Work) {
	fmt.Println("registered a Book: " + w.Title)
}
