package factory

import (
	"errors"
	"fmt"

	"github.com/sh-yamaga/design-patterns-go/internal/patterns/factory/work"
)

type IWorkFactory interface {
	Create(title, creater string) *work.Work
	createWork(title, creater string) *work.Work
	registerWork(w *work.Work)
}

type RootWorkFactory struct {
	Category work.WorkCategory
}

func (rwf RootWorkFactory) Create(wc work.WorkCategory) IWorkFactory {
	factory, err := rwf.registerFactory(wc)
	if err != nil {
		fmt.Println(err)
	}

	return factory
}

func (rfc RootWorkFactory) registerFactory(wc work.WorkCategory) (IWorkFactory, error) {
	switch wc {
	case work.BookCategory:
		return BookFactory{}, nil
	case work.MovieCategory:
		return MovieFactory{}, nil
	default:
		return nil, errors.New("unexpected WorkCategory was given")
	}
}
