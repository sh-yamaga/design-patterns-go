package creator

import (
	"fmt"

	"github.com/sh-yamaga/design-patterns-go/internal/patterns/factory/category"
	"github.com/sh-yamaga/design-patterns-go/internal/patterns/factory/work"
)

type MovieCreator struct {
	BaseCreator
}

func (mc MovieCreator) createWork(title, creater string) *work.Work {
	return &work.Work{
		Title:    title,
		Creater:  creater,
		Category: category.Movie,
	}
}

func (mc MovieCreator) registerWork(w *work.Work) {
	fmt.Println("registered a Movie: " + w.Title)
}
