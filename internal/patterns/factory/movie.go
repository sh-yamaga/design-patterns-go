package factory

import (
	"fmt"

	"github.com/sh-yamaga/design-patterns-go/internal/patterns/factory/work"
)

type MovieFactory struct {
	WorkFactory
}

func (mf MovieFactory) Create(title, creater string) *work.Work {
	var movie *work.Work = mf.CreateWork(title, creater)
	mf.RegisterWork(movie)

	return movie
}

func (mf MovieFactory) CreateWork(title, creater string) *work.Work {
	return &work.Work{
		Title:    title,
		Creater:  creater,
		Category: work.MovieCategory,
	}
}

func (mf MovieFactory) RegisterWork(w *work.Work) {
	fmt.Println("registered a Movie: " + w.Title)
}
