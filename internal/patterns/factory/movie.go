package factory

import (
	"fmt"

	"github.com/sh-yamaga/design-patterns-go/internal/patterns/factory/work"
)

type MovieFactory struct {
	Factory
}

func (mf MovieFactory) createWork(title, creater string) *work.Work {
	return &work.Work{
		Title:    title,
		Creater:  creater,
		Category: work.MovieCategory,
	}
}

func (mf MovieFactory) registerWork(w *work.Work) {
	fmt.Println("registered a Movie: " + w.Title)
}
