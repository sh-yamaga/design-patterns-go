package creator

import (
	"fmt"

	"github.com/sh-yamaga/design-patterns-go/internal/patterns/factory/category"
	"github.com/sh-yamaga/design-patterns-go/internal/patterns/factory/work"
)

type MovieCreator struct {
	BaseCreator
}

func (mc MovieCreator) createWork(title, creator string) *work.Work {
	return &work.Work{
		Title:    title,
		Creater:  creator,
		Category: category.Movie,
	}
}

func (mc MovieCreator) registerWork(w *work.Work) work.IWork {
	fmt.Printf("registered movie: %s\n", w.Title)
	var movie work.Movie = work.Movie{
		Work: *w,
	}

	return movie
}
